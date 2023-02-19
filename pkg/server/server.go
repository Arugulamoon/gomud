package server

import (
	"log"
	"net"

	"github.com/Arugulamoon/gomud/pkg/session"
)

func handleConnection(conn net.Conn, sessionEvent chan session.SessionEvent) error {
	buf := make([]byte, 4096)

	s := &session.Session{
		Id:         session.GenerateSessionId(),
		Connection: conn,
	}
	log.Println("Server accepted connection and created session:", s.SessionId())

	// Broadcast user joined
	sessionEvent <- session.SessionEvent{
		Session: s,
		Event:   &session.SessionCreatedEvent{},
	}

	for {
		// Broadcast user input
		n, err := conn.Read(buf)
		if err != nil {
			sessionEvent <- session.SessionEvent{
				Session: s,
				Event:   &session.SessionDisconnectEvent{},
			}
			return err
		}
		if n == 0 {
			log.Println("Zero bytes, closing connection for session:", s.SessionId())
			sessionEvent <- session.SessionEvent{
				Session: s,
				Event:   &session.SessionDisconnectEvent{},
			}
			break
		}

		if n > 2 {
			msg := string(buf[0 : n-2])
			log.Printf("Received message on session %s: \"%s\"\r\n", s.SessionId(), msg)

			sessionEvent <- session.SessionEvent{
				Session: s,
				Event: &session.SessionInputEvent{
					Input: msg,
				},
			}
		}
	}

	return nil
}

func Start(eventChannel chan session.SessionEvent) error {
	log.Println("Starting async tcp server to receive messages")

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting connection", err)
			continue
		}
		// Start an async handler to receive messages
		go func(c net.Conn) {
			if err := handleConnection(c, eventChannel); err != nil {
				log.Println("Error handling connection", err)
				return
			}
		}(conn)
	}
}
