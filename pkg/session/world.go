package session

import "fmt"

type World struct {
	Characters []*Character
	Rooms      []*Room
}

func NewWorld() *World {
	return &World{}
}

func (w *World) Init() {
	w.Rooms = []*Room{
		{
			Id:   "A",
			Desc: "This is a room with a sign that has the letter A written on it.",
			Links: []*RoomLink{
				{
					Verb:   "east",
					RoomId: "B",
				},
			},
		},
		{
			Id:   "B",
			Desc: "This is a room with a sign that has the letter B written on it.",
			Links: []*RoomLink{
				{
					Verb:   "west",
					RoomId: "A",
				},
			},
		},
	}
}

func (w *World) HandleCharacterJoined(character *Character) {
	w.Rooms[0].AddCharacter(character)

	character.SendMessage(fmt.Sprintf("Welcome %s!", character.Name))
	character.SendMessage("")
	character.SendMessage(character.Room.Desc)
}

func (w *World) HandleCharacterLeft(character *Character) {
	character.Room.RemoveCharacter(character) // Weird; char removing self from room
}

func (w *World) getRoomById(id string) *Room {
	for _, r := range w.Rooms {
		if r.Id == id {
			return r
		}
	}
	return nil
}

func (w *World) HandleCharacterInput(character *Character, input string) {
	room := character.Room
	for _, link := range room.Links {
		if link.Verb == input {
			target := w.getRoomById(link.RoomId)
			if target != nil {
				w.MoveCharacter(character, target)
				return
			}
		}
	}

	character.SendMessage(fmt.Sprintf("You said, \"%s\"", input))

	for _, other := range character.Room.Characters {
		if other != character {
			other.SendMessage(fmt.Sprintf("%s said, \"%s\"", character.Name, input))
		}
	}
}

func (w *World) MoveCharacter(character *Character, targetRoom *Room) {
	// Update Rooms
	currentRoom := character.Room
	currentRoom.RemoveCharacter(character)
	targetRoom.AddCharacter(character)

	// Update Character
	character.SendMessage(character.Room.Desc)
}
