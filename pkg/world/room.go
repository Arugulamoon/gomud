package world

import (
	"fmt"

	"github.com/Arugulamoon/gomud/pkg/character"
)

type Room struct {
	Id, Description string
	Links           map[string]*Room
	Characters      map[string]*character.Character
}

func NewRoom(id, desc string) *Room {
	return &Room{
		Id:          id,
		Description: desc,
		Links:       make(map[string]*Room),
		Characters:  make(map[string]*character.Character),
	}
}

func (r *Room) BroadcastMessage(speaker, msg string) {
	for _, char := range r.Characters {
		if char.Name != speaker {
			char.SendMessage(msg)
		}
	}
}

func (r *Room) ContainsCharacter(name string) bool {
	for _, character := range r.Characters {
		if character.Name == name {
			return true
		}
	}
	return false
}

func (r *Room) AddCharacter(c *character.Character) {
	r.Characters[c.Id] = c
	r.BroadcastMessage(c.Name, fmt.Sprintf("%s entered the room.", c.Name))
}

func (r *Room) RemoveCharacter(c *character.Character) {
	delete(r.Characters, c.Id)
	r.BroadcastMessage(c.Name, fmt.Sprintf("%s left the room.", c.Name))
}
