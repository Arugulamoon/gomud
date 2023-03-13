package session

import (
	"fmt"
	"log"
	"net"
	"strings"
	"unicode"

	"github.com/Arugulamoon/gomud/pkg/character"
)

type Session struct {
	Id           string
	Connection   net.Conn
	EventChannel chan SessionEvent

	Character *character.Character
}

func New(c net.Conn, ch chan SessionEvent) *Session {
	s := &Session{
		Id:           generateId(),
		Connection:   c,
		EventChannel: ch,
	}
	s.Character = character.New(s)
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

		msg := trimEOL(buf)
		log.Printf("Received message on session %s: %s", s.SessionId(), msg)

		s.EventChannel <- SessionEvent{
			Session: s,
			Event: &SessionInputEvent{
				Input: msg,
			},
		}
	}

	return nil
}

// May receive messages ending with \r or \r\n
func trimEOL(input []byte) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, string(input))
}

var nextId = 1

func generateId() string {
	var id = nextId
	nextId++
	return fmt.Sprintf("%d", id)
}
