package server

// CONTEXT: Server

import (
	"log"
	"net"

	"github.com/Arugulamoon/gomud/pkg/world"
)

func Run(ch chan world.SessionEvent) error {
	log.Println("Starting async tcp server to receive messages")

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("Error accepting connection", err)
			continue
		}
		// Start an async handler to receive messages
		go handleConnection(c, ch)
	}
}

func handleConnection(c net.Conn, ch chan world.SessionEvent) {
	s := world.NewSession(c, ch)
	if err := s.Tail(); err != nil {
		log.Println("Error handling connection", err)
		return
	}
}
