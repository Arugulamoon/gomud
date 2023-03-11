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

func NewCharacter() *Character {
	return &Character{
		Id:   generateCharacterId(),
		Name: generateCharacterName(),
	}
}

var nextCharacterId = 1

func generateCharacterId() string {
	var id = nextCharacterId
	nextCharacterId++
	return fmt.Sprintf("%d", id)
}

func generateCharacterName() string {
	return fmt.Sprintf("Character %d", rand.Intn(100)+1)
}
