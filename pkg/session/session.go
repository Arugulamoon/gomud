package session

import (
	"fmt"
	"log"
	"net"
)

type Session struct {
	Id           string
	Connection   net.Conn
	EventChannel chan SessionEvent

	User *User
}

func New(c net.Conn, ch chan SessionEvent) *Session {
	s := &Session{
		Id:           generateId(),
		Connection:   c,
		EventChannel: ch,

		User: &User{
			Character: &Character{
				Name: GenerateName(),
			},
		},
	}
	log.Println("Server accepted connection and created session:", s.SessionId())

	// Broadcast Event: Session Created (User Joined)
	s.EventChannel <- SessionEvent{
		Session: s,
		Event:   &SessionCreatedEvent{},
	}

	return s
}

func (s *Session) SessionId() string {
	return s.Id
}

// TODO: Make non-blocking for scaling
func (s *Session) WriteLine(str string) error {
	_, err := s.Connection.Write([]byte(str + "\r\n"))
	return err
}

func (s *Session) Stream() error {
	buf := make([]byte, 4096)
	for {
		// Broadcast user input
		n, err := s.Connection.Read(buf)
		if err != nil {
			s.EventChannel <- SessionEvent{
				Session: s,
				Event:   &SessionDisconnectEvent{},
			}
			return err
		}
		if n == 0 {
			log.Println("Zero bytes, closing connection for session:", s.SessionId())
			s.EventChannel <- SessionEvent{
				Session: s,
				Event:   &SessionDisconnectEvent{},
			}
			break
		}

		if n > 2 {
			msg := string(buf[0 : n-2])
			log.Printf("Received message on session %s: \"%s\"\r\n", s.SessionId(), msg)

			s.EventChannel <- SessionEvent{
				Session: s,
				Event: &SessionInputEvent{
					Input: msg,
				},
			}
		}
	}

	return nil
}

var nextId = 1

func generateId() string {
	var sId = nextId
	nextId++
	return fmt.Sprintf("%d", sId)
}
