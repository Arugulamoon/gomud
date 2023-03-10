package world

// CONTEXT: World

type SessionHandler struct {
	World        *World
	EventChannel <-chan SessionEvent
}

func NewSessionHandler(w *World, ch <-chan SessionEvent) *SessionHandler {
	return &SessionHandler{
		World:        w,
		EventChannel: ch,
	}
}

func (h *SessionHandler) Start() {
	for sessionEvent := range h.EventChannel {
		s := sessionEvent.Session

		switch event := sessionEvent.Event.(type) {

		case *SessionCreateEvent:
			h.World.HandleCharacterJoined(s)

		case *SessionDisconnectEvent:
			h.World.HandleCharacterLeft(s)

		case *SessionInputEvent:
			h.World.HandleCharacterInput(s, event.Input)
		}
	}
}
