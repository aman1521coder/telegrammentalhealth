package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	//"errors"
	"fmt"
	//"os"
)

const (
	SendMessageMethod = "sendMessage"
	GetUpdatesMethod  = "getUpdates"
)

var token = "8411335043:AAHLWnflciTp11AaH9jj8zjCuO5luBCBTZA"
type APIResponse struct {
	OK          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result,omitempty"` // Unmarshal manually based on method
	ErrorCode   int                 `json:"error_code,omitempty"`
	Description string              `json:"description,omitempty"`
	Parameters  *ResponseParameters `json:"response_parameters,omitempty"`
}
type Update struct {
	UpdateID int      `json:"update_id"`
	Message  *Message `json:"message,omitempty"`
}
type Message struct {
	MessageID int    `json:"message_id"`
	From      *User  `json:"from,omitempty"`
	Chat      Chat   `json:"chat"`
	Date      int    `json:"date"`
	Text      string `json:"text,omitempty"`
}
type User struct {
	ID        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username,omitempty"`
}
type Chat struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}
type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int   `json:"retry_after,omitempty"`
}

// Params for sendMessage
type SendMessageParams struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}
type GetUpdatesParams struct {
	Offset  int `json:"offset,omitempty"`
	Timeout int `json:"timeout,omitempty"`
}

func CallTelegraMethod(ctx context.Context, method string, params interface{}, token string, httpMethod string) (*APIResponse, error) {
	url := "https://api.telegram.org/bot" + token + "/" + method

	jsondata, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, httpMethod, url, bytes.NewBuffer(jsondata))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("Raw response:", string(data))
	var apiResp APIResponse
	err = json.Unmarshal(data, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil

}

func SendMessage(ctx context.Context, chatID int64, text string, token string) error {
	params := SendMessageParams{ChatID: chatID, Text: text}
	resp, err := CallTelegraMethod(ctx, SendMessageMethod, params, token, "POST")
	if err != nil {
		return err
	}
	if !resp.OK {
		return fmt.Errorf("failed to send message  %d: %s", resp.ErrorCode, resp.Description)

	}
	return nil
}

func GetUpdates(ctx context.Context, token string, update chan<- *Update) error {

	offset := 0
	for {
		params := GetUpdatesParams{
			Offset:  offset,
			Timeout: 10,
		}
		resp, err := CallTelegraMethod(ctx, GetUpdatesMethod, params, token,"GET")
		if err != nil {
			return err
		}
		if !resp.OK {
			return fmt.Errorf("failed to send message  %d: %s", resp.ErrorCode, resp.Description)
		}

		result, err := Handlresult(GetUpdatesMethod, resp.Result)
		if err != nil {
			return err
		}
		for _, u := range result.([]Update) {
			update <- &u
			offset = u.UpdateID + 1
		}

	}

}
func Handlresult(method string, resp json.RawMessage) (interface{}, error) {
	switch method {
	case SendMessageMethod:
		var msg Message
		err := json.Unmarshal(resp, &msg)
		if err != nil {
			return nil, err
		}
		return msg, nil
	case GetUpdatesMethod:
		var updates []Update
		err := json.Unmarshal(resp, &updates)
		if err != nil {
			return nil, err
		}
		return updates, nil
	default:
		return nil, fmt.Errorf("unknown method %s", method)
	}

}
func HandleMessages(updateChan <-chan *Update) error {
	for upd := range updateChan {
		if upd.Message == nil || upd.Message.From.IsBot {
			continue
		}

		msg := upd.Message

		switch msg.Text {
		case "/start":
			s, err := StartSession(msg.From)
			if err != nil {
				SendMessage(context.Background(), msg.Chat.ID, err.Error(), token)
				continue
			}
			SendMessage(context.Background(), msg.Chat.ID,
				"Session started. You are now connected to an expert.", token)
            SendMessage(context.Background(),s.expert.ChatID,"session started",token)
			log.Printf("Session %s started: user=%d expert=%d",
				s.ID, s.user.ID, s.expert.ID)

		case "/end":
			err := EndSession(msg.From.ID)
			if err != nil {
				SendMessage(context.Background(), msg.Chat.ID, err.Error(), token)
				continue
			}
			SendMessage(context.Background(), msg.Chat.ID, "Session ended.", token)

		case "/help":
			SendMessage(context.Background(), msg.Chat.ID,
				"Commands:\n/start\n/end\n/help", token)

		default:
			// normal chat message
			if err := RouteMessage(msg); err != nil {
				log.Println("route error:", err)
			}
		}
	}
	return nil
}

func RouteMessage(msg *Message) error {
	mu.Lock()
	defer mu.Unlock()
	if msg.From.IsBot {
		return fmt.Errorf("not human ")
	}
	s, ok := userSessions[msg.From.ID]
	if !ok || !s.active {
		return fmt.Errorf(" something issue while sending message")
	}
	err := SendMessage(context.Background(), int64(s.expert.ID), msg.Text, token)
	if err != nil {
		return fmt.Errorf("something  is wrong while routing message to expert : %v", err)
	}

	e, ok := expertSessions[msg.Chat.ID]
	if !ok || !e.active {
		return fmt.Errorf("something issue while sending message")
	}
	err = SendMessage(context.Background(), int64(e.user.ID), msg.Text, token)
	if err != nil {
		return fmt.Errorf("something  is wrong while routing message to user : %v", err)
	}
	return nil

}
