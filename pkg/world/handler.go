package world

import "github.com/Arugulamoon/gomud/pkg/session"

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

		switch event := sessionEvent.Event.(type) {

		case *session.SessionCreateEvent:
			h.World.HandleCharacterJoined(c)

		case *session.SessionDisconnectEvent:
			h.World.HandleCharacterLeft(c)

		case *session.SessionInputEvent:
			h.World.HandleCharacterInput(c, event.Input)
		}
	}
}
