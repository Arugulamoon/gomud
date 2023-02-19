package session

type SessionEvent struct {
	Session *Session
	Event   interface{}
}

type SessionCreatedEvent struct{}

type SessionDisconnectEvent struct{}

type SessionInputEvent struct {
	Input string
}
