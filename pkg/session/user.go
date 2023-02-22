package session

import (
	"fmt"
	"math/rand"
)

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

func GenerateName() string {
	return fmt.Sprintf("User %d", rand.Intn(100)+1)
}

// Character
type Character struct {
	Name string
	Room *Room
}
