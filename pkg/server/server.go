package server

import (
	"crypto/tls"
	"log"
	"net"

	"github.com/Arugulamoon/gomud/pkg/session"
)

func Run(ch chan session.SessionEvent) error {
	log.Println("Starting async tcp server to receive messages")

	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}

	l, err := tls.Listen("tcp", ":7324", &tls.Config{
		Certificates: []tls.Certificate{cert}},
	)
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

func handleConnection(c net.Conn, ch chan session.SessionEvent) {
	s := session.New(c, ch)
	if err := s.Tail(); err != nil {
		log.Println("Error handling connection", err)
		return
	}
}
