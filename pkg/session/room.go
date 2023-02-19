package session

import "fmt"

type Room struct {
	Id, Desc   string
	Links      []*RoomLink
	Characters []*Character // TODO: Change to map
}

func (r *Room) SendMessage(character *Character, msg string) {
	for _, other := range r.Characters {
		if other != character {
			other.SendMessage(msg)
		}
	}
}

func (r *Room) AddCharacter(character *Character) {
	r.Characters = append(r.Characters, character)
	character.Room = r

	r.SendMessage(character, fmt.Sprintf("%s entered the room.", character.Name))
}

// TODO: Optimize?
func (r *Room) RemoveCharacter(character *Character) {
	character.Room = nil

	var characters []*Character
	for _, c := range r.Characters {
		if c != character {
			characters = append(characters, c)
		}
	}
	r.Characters = characters

	r.SendMessage(character, fmt.Sprintf("%s left the room.", character.Name))
}

type RoomLink struct {
	Verb, RoomId string
}
