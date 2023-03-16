package world

import (
	"fmt"
	"strings"

	"github.com/Arugulamoon/gomud/pkg/character"
	"github.com/Arugulamoon/gomud/pkg/command"
)

type World struct {
	Characters map[string]*character.Character
	Rooms      map[string]*Room
}

func New() *World {
	return &World{
		Characters: make(map[string]*character.Character),
	}
}

// TODO: Move into yaml
func (w *World) Load() {
	w.Rooms = map[string]*Room{
		"Bedroom": {
			Id:          "Bedroom",
			Description: "You have entered your bedroom. There is a door leading out! (type \"/open door\" to leave the bedroom)",
			Links: map[string]*RoomLink{
				"/open door": {
					Verb:   "/open door",
					RoomId: "Hallway",
				},
			},
			Characters: make(map[string]*character.Character),
		},
		"Hallway": {
			Id:          "Hallway",
			Description: "You have entered a hallway with doors at either end. (type \"/open north door\" to enter the living room or \"/open south door\" to enter the bedroom)",
			Links: map[string]*RoomLink{
				"/open north door": {
					Verb:   "/open north door",
					RoomId: "LivingRoom",
				},
				"/open south door": {
					Verb:   "/open south door",
					RoomId: "Bedroom",
				},
			},
			Characters: make(map[string]*character.Character),
		},
		"LivingRoom": {
			Id:          "LivingRoom",
			Description: "You have entered the living room. (type \"/open door\" to enter the hallway)",
			Links: map[string]*RoomLink{
				"/open door": {
					Verb:   "/open door",
					RoomId: "Hallway",
				},
			},
			Characters: make(map[string]*character.Character),
		},
	}
}

func (w *World) GetCharacterNames() []string {
	// TODO: Make more efficient with map/filter/reduce?
	var names []string
	for _, char := range w.Characters {
		names = append(names, char.Name)
	}
	return names
}

func (w *World) GetCharacters() map[string]*character.Character {
	return w.Characters
}

func (w *World) HandleCharacterJoined(c *character.Character) {
	w.Characters[c.Id] = c
	c.World = w
	w.Rooms["Bedroom"].AddCharacter(c)

	c.SendMessage(fmt.Sprintf("Welcome %s!", c.Name))
	c.SendMessage("")
	c.SendMessage(c.Room.GetDescription())
}

func (w *World) HandleCharacterLeft(c *character.Character) {
	room := w.Rooms[c.Room.GetId()]
	room.RemoveCharacter(c)
	delete(w.Characters, c.Id)
}

func (w *World) getRoomById(id string) *Room {
	for _, r := range w.Rooms {
		if r.Id == id {
			return r
		}
	}
	return nil
}

func (w *World) HandleCharacterInput(char *character.Character, inp string) {
	room := w.Rooms[char.Room.GetId()]

	// If match RoomLink then move
	targetRoom := w.moveToRoom(room, inp)
	if targetRoom != nil {
		w.MoveCharacter(char, targetRoom)
		return
	}

	// otherwise, separate into verb/action and args
	verb, args := handleCharacterInput(char, inp)

	switch verb {
	case command.SAY:
		command.Say(char, args)
	case command.SHOUT:
		command.Shout(char, args)
	case command.TELL:
		command.Tell(char, args)
	case command.WAVE:
		command.Wave(char, args)
	case command.WHO:
		command.Who(char, args)
	}
}

func (w *World) moveToRoom(currentRoom *Room, inp string) *Room {
	for verb, link := range currentRoom.RoomLinks() {
		if verb == inp {
			target := w.getRoomById(link.RoomId)
			if target != nil {
				return target
			}
		}
	}
	return nil
}

func handleCharacterInput(c *character.Character, input string) (string, string) {
	verb := "say"
	args := input
	if input[0:1] == "/" {
		var cmd string
		cmd, args, _ = strings.Cut(input, " ")
		verb = cmd[1:]
	}
	return verb, args
}

func (w *World) MoveCharacter(c *character.Character, targetRoom *Room) {
	// Update Rooms
	currentRoom := w.Rooms[c.Room.GetId()]
	currentRoom.RemoveCharacter(c)
	targetRoom.AddCharacter(c)

	// Update Character
	c.SendMessage(c.Room.GetDescription())
}

func (w *World) BroadcastMessage(speaker, msg string) {
	for _, char := range w.Characters {
		if char.Name != speaker {
			char.SendMessage(msg)
		}
	}
}

func (w *World) ContainsCharacter(name string) bool {
	for _, character := range w.Characters {
		if character.Name == name {
			return true
		}
	}
	return false
}
