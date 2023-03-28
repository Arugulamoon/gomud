package room

import (
	"fmt"

	"github.com/Arugulamoon/gomud/pkg/character"
	"github.com/Arugulamoon/gomud/pkg/item"
)

type Room struct {
	Id, Description string

	Paths map[string]*Room
	Items map[string]*item.Item

	Characters map[string]*character.Character
}

func NewRoom(id, desc string) *Room {
	return &Room{
		Id:          id,
		Description: desc,
		Paths:       make(map[string]*Room),
		Items:       make(map[string]*item.Item),
		Characters:  make(map[string]*character.Character),
	}
}

func (r *Room) GetCharacters() map[string]*character.Character {
	return r.Characters
}

func (r *Room) ContainsCharacter(name string) bool {
	for _, character := range r.Characters {
		if character.Name == name {
			return true
		}
	}
	return false
}

func (r *Room) BroadcastMessage(speaker, msg string) {
	for _, char := range r.Characters {
		if char.Name != speaker {
			char.SendMessage(msg)
		}
	}
}

func (r *Room) Describe(char *character.Character) {
	char.SendMessage(r.Description)
	if len(r.Items) > 0 {
		char.SendMessage("Items:")
		for _, item := range r.Items {
			char.SendMessage(fmt.Sprintf("  %s", item.Id))
		}
	}
}

func (r *Room) AddCharacter(c *character.Character) {
	r.Characters[c.Id] = c
	r.BroadcastMessage(c.Name, fmt.Sprintf("%s entered the room.", c.Name))
}

func (r *Room) RemoveCharacter(c *character.Character) {
	delete(r.Characters, c.Id)
	r.BroadcastMessage(c.Name, fmt.Sprintf("%s left the room.", c.Name))
}

func (r *Room) Say() {}
