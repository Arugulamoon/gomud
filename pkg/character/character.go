package character

import (
	"fmt"
	"math/rand"
)

type room interface {
	GetId() string
	GetDescription() string
}

type Character struct {
	Id, Name string
	Room     room
}

func New() *Character {
	return &Character{
		Id:   generateId(),
		Name: generateName(),
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
