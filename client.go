package main

import (
	//"time"
	//"fmt"

	"context"
	"errors"

	"log"
	"sync"

	"github.com/google/uuid"
)

type Session struct{
ID    string
expert  *Expert
user *User
active bool 
}
type Expert struct{
	ID int64
	ChatID  int64

	Available bool
}
var (

	Experts = []*Expert{
    {ID: 1, ChatID: 8190427124, Available: true},
    //{ID: 2, ChatID: 1002, Available: true},
    //{ID: 3, ChatID: 1003, Available: true},
}
sessionsByID= make(map[string]*Session)
    userSessions   = make(map[int64]*Session) // Telegram user ID → session
    expertSessions = make(map[int64]*Session) // Expert chat ID → session


    mu          sync.Mutex
)




func (e *Expert) IsAvailable() bool{
	return e.Available
}


func StartSession(user *User )(*Session,error){
	mu.Lock()
	defer mu.Unlock()
	if user.IsBot{
		return nil,errors.New("bots are not allowed to start a session")
	}

	if _, exists := userSessions[user.ID]; exists {
    return nil, errors.New("user already has an active session")
}


	for _,e:=range Experts{
		if e.IsAvailable(){
		e.Available=false
			s:=&Session{
				ID: uuid.NewString(),
				expert: e,
				user: user,
				active: true,
			}
			userSessions[user.ID]=s
			expertSessions[e.ID]=s
		     sessionsByID[s.ID] = s

            log.Println("session started:", s.ID)
            return s, nil
		}
	}
	return nil,errors.New("no vaialable expert at the moment please try again after sonme time")
}
func EndSession(userId int64) error{
	
	mu.Lock()


	s, ok := userSessions[userId]
	if !ok {
		return errors.New("the session is not availavble")
	}
	sessionID := s.ID

	s.expert.Available = true
	s.active = false

    delete(userSessions, s.user.ID)
    delete(expertSessions, s.expert.ID)
    delete(sessionsByID, sessionID)
		 	mu.Unlock()
     _=SendMessage(context.Background(),s.expert.ChatID,"Session ended.",token)


	log.Printf("Session %s ended", sessionID)
	return nil
	

}

