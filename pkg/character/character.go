package character

import (
	"fmt"
	"math/rand"
)

type session interface {
	WriteLine(msg string) error
}

type Character struct {
	Id, Name        string
	Session         session
	WorldId, RoomId string
}

func New(s session) *Character {
	return &Character{
		Id:      generateId(),
		Name:    generateName(),
		Session: s,
	}
}

func (c *Character) SendMessage(msg string) {
	c.Session.WriteLine(msg)
}

var nextId = 1

func generateId() string {
	var id = nextId
	nextId++
	return fmt.Sprintf("%d", id)
}

func generateName() string {
	return fmt.Sprintf("Character %d", rand.Intn(100)+1)
}
