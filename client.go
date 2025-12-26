package main

import (
	//"time"
	//"fmt"

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
	ID int 
	Available bool
}
var (

	Experts = []*Expert{
    {ID: 1, Available: true},
    {ID: 2, Available: true},
    {ID: 3, Available: true},
}
sessions = make(map[string]*Session)
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
	s:=&Session{}
	for _,e:=range Experts{
		if e.IsAvailable(){
			s.ID=uuid.NewString()
			e.Available=false
			s.expert=e
			s.user=user
			s.active=true
			log.Print("session started")
			sessions[s.ID]=s
			return s,nil
		}
	}
	return nil,errors.New("no vaialable expert at the moment please try again after sonme time")
}
func EndSession(sessionID string) error{
	
	mu.Lock()
	defer mu.Unlock()

	s, ok := sessions[sessionID]
	if !ok {
		return errors.New("the session is not availavble")
	}

	s.expert.Available = true
	s.active = false

	delete(sessions, sessionID)
	log.Printf("Session %s ended", sessionID)
	return nil
	

}

