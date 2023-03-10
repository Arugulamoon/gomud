package world

import (
	"fmt"
	"math/rand"
)

// TODO
// Entity
type Entity struct {
	Id string
}

func (e *Entity) EntityId() string {
	return e.Id
}

// User
type User struct {
	Character *Character
}

// TODO: Is it generate character name not User?
func GenerateName() string {
	return fmt.Sprintf("User %d", rand.Intn(100)+1)
}

type room interface {
	Description() string
	RoomLinks() []*RoomLink                 // TODO: Remove
	ConnectedSessions() map[string]*Session // TODO: Remove
	ContainsCharacter(name string) bool     // TODO: Remove?
	RemoveCharacter(s *Session)             // TODO: Remove?
}

// Character
type Character struct {
	Name string
	Room room
}
