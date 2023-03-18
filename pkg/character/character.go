package character

import (
	"fmt"
	"math/rand"
)

type session interface {
	WriteLine(msg string) error
}

type world interface {
	GetCharacters() map[string]*Character
	GetCharacterNames() []string
	ContainsCharacter(args string) bool
}

type room interface {
	GetId() string
	GetDescription() string
	GetCharacterNames() []string
	ContainsCharacter(args string) bool
	RemoveCharacter(*Character)
}

type Character struct {
	Id, Name string
	Session  session
	World    world
	Room     room
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
