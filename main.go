package main

import (
	"log"

	"github.com/Arugulamoon/gomud/pkg/server"
	"github.com/Arugulamoon/gomud/pkg/session"
)

func main() {
	// Create a channel to receive events
	sessionEventChannel := make(chan session.SessionEvent)

	// Create and initialize world
	world := session.NewWorld()
	world.Init()

	// Start an async handler to react to events
	sessionHandler := *session.NewSessionHandler(world, sessionEventChannel)
	go sessionHandler.Start()

	// Start an async tcp server to receive connections
	// - Announce New Connections by creating user joined events
	// Maintain connections and receive inputs
	// - Announce messages by creating message events
	// Translate inputs into Events
	// Disconnect connections
	if err := server.Start(sessionEventChannel); err != nil {
		log.Fatal(err)
	}
}
