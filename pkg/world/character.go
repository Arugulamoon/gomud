package world

import (
	"fmt"
	"math/rand"
)

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

func NewCharacter() *Character {
	return &Character{
		Name: generateName(),
	}
}

func generateName() string {
	return fmt.Sprintf("Character %d", rand.Intn(100)+1)
}
