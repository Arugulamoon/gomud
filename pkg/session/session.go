package session

import (
	"fmt"
	"log"
	"net"

	"github.com/Arugulamoon/gomud/pkg/character"
)

type Session struct {
	Id           string
	Connection   net.Conn
	EventChannel chan SessionEvent

	Character *character.Character
}

func NewSession(c net.Conn, ch chan SessionEvent) *Session {
	s := &Session{
		Id:           generateSessionId(),
		Connection:   c,
		EventChannel: ch,

		Character: character.NewCharacter(),
	}
	log.Println("Server accepted connection and created session:", s.SessionId())

	// Broadcast Event: Session Created (User Joined)
	s.EventChannel <- SessionEvent{
		Session: s,
		Event:   &SessionCreateEvent{},
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

func (s *Session) Tail() error {
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

var nextSessionId = 1

func generateSessionId() string {
	var id = nextSessionId
	nextSessionId++
	return fmt.Sprintf("%d", id)
}
