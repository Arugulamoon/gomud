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
	Items           []string
}

func New(s session) *Character {
	return &Character{
		Id:      generateId(),
		Name:    generateName(),
		Session: s,
		WorldId: "Prototype",
		RoomId:  "Bedroom",
		Items:   make([]string, 0),
	}
}

func (c *Character) SendMessage(msg string) {
	c.Session.WriteLine(msg)
}

func (c *Character) PickUp(itemId string) error {
	c.Items = append(c.Items, itemId)
	c.SendMessage(fmt.Sprintf("You pick up %s.", itemId))
	return nil
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
