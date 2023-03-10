package world

// CONTEXT: Session? Server?

type SessionEvent struct {
	Session *Session
	Event   interface{}
}

type SessionCreateEvent struct{}

type SessionDisconnectEvent struct{}

type SessionInputEvent struct {
	Input string
}
