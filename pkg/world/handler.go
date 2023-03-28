package world

import (
	"log"

	"github.com/Arugulamoon/gomud/pkg/session"
)

type SessionHandler struct {
	World        *World
	EventChannel <-chan session.SessionEvent
}

func NewSessionHandler(w *World, ch <-chan session.SessionEvent) *SessionHandler {
	return &SessionHandler{
		World:        w,
		EventChannel: ch,
	}
}

func (h *SessionHandler) Start() {
	for sessionEvent := range h.EventChannel {
		c := sessionEvent.Session.Character

		var err error
		switch event := sessionEvent.Event.(type) {

		case *session.SessionCreateEvent:
			err = h.World.HandleCharacterJoined(c)

		case *session.SessionDisconnectEvent:
			err = h.World.HandleCharacterLeft(c)

		case *session.SessionInputEvent:
			err = h.World.HandleCharacterInput(c, event.Input)
		}

		if err != nil {
			log.Fatalln(err)
		}
	}
}
