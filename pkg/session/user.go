package session

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

// Character
type Character struct {
	Name string
	Room *Room
}
