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
	Session   *Session
	Character *Character
}

func GenerateName() string {
	return fmt.Sprintf("User %d", rand.Intn(100)+1)
}

// Character
type Character struct {
	Name string
	User *User
	Room *Room
}

func (c *Character) SendMessage(msg string) {
	c.User.Session.WriteLine(msg)
}
