package character

import (
	"fmt"
	"math/rand"
)

type session interface {
	WriteLine(str string) error
}

type room interface {
	GetId() string
	GetDescription() string
}

type Character struct {
	Id, Name string
	Session  session
	Room     room
}

func New(s session) *Character {
	return &Character{
		Id:      generateId(),
		Name:    generateName(),
		Session: s,
	}
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
