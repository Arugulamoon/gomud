package main

import (
	"log"

	"github.com/Arugulamoon/gomud/pkg/server"
	"github.com/Arugulamoon/gomud/pkg/session"
)

func main() {
	// Create and initialize world
	w := session.NewWorld()
	w.Init()

	// Create a channel to receive session events
	sessionEventChannel := make(chan session.SessionEvent)

	// Start an async handler to react to session events
	h := *session.NewSessionHandler(w, sessionEventChannel)
	go h.Start()

	// Start an async tcp server to receive connections
	// - Announce New Connections by creating user joined events
	// Maintain connections and receive inputs
	// - Announce messages by creating message events
	// Translate inputs into Events
	// Disconnect connections
	if err := server.Run(sessionEventChannel); err != nil {
		log.Fatal(err)
	}
}
