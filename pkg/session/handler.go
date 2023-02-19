package session

type SessionHandler struct {
	World        *World
	EventChannel <-chan SessionEvent
	Users        map[string]*User
}

func NewSessionHandler(w *World, ch <-chan SessionEvent) *SessionHandler {
	return &SessionHandler{
		World:        w,
		EventChannel: ch,
		Users:        map[string]*User{},
	}
}

func (h *SessionHandler) Start() {
	for sessionEvent := range h.EventChannel {
		session := sessionEvent.Session
		sId := session.SessionId()

		switch event := sessionEvent.Event.(type) {

		case *SessionCreatedEvent:
			character := &Character{
				Name: GenerateName(),
			}

			user := &User{
				Session:   session,
				Character: character,
			}
			character.User = user

			h.Users[sId] = user
			h.World.HandleCharacterJoined(character)

		case *SessionDisconnectEvent:
			user := h.Users[sId]
			h.World.HandleCharacterLeft(user.Character)

		case *SessionInputEvent:
			user := h.Users[sId]
			h.World.HandleCharacterInput(user.Character, event.Input)
		}
	}
}
