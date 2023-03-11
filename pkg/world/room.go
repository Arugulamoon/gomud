package world

import (
	"fmt"

	"github.com/Arugulamoon/gomud/pkg/character"
)

type Room struct {
	Id, Description string
	Links           []*RoomLink
	Characters      map[string]*character.Character
}

func (r *Room) GetId() string {
	return r.Id
}

func (r *Room) GetDescription() string {
	return r.Description
}

func (r *Room) RoomLinks() []*RoomLink {
	return r.Links
}

func (r *Room) GetCharacters() map[string]*character.Character {
	return r.Characters
}

func (r *Room) SendMessage(c *character.Character, msg string) {
	for id, other := range r.Characters {
		if id != c.Id {
			other.SendMessage(msg)
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
	c.Room = r
	r.SendMessage(c, fmt.Sprintf("%s entered the room.", c.Name))
}

func (r *Room) RemoveCharacter(c *character.Character) {
	delete(r.Characters, c.Id)
	c.Room = nil
	r.SendMessage(c, fmt.Sprintf("%s left the room.", c.Name))
}

type RoomLink struct {
	Verb, RoomId string
}
