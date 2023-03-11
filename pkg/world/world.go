package world

import (
	"fmt"
	"strings"

	"github.com/Arugulamoon/gomud/pkg/character"
	"github.com/Arugulamoon/gomud/pkg/input"
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
			Links: []*RoomLink{
				{
					Verb:   "/open door",
					RoomId: "Hallway",
				},
			},
			Characters: make(map[string]*character.Character),
		},
		"Hallway": {
			Id:          "Hallway",
			Description: "You have entered a hallway with doors at either end. (type \"/open north door\" to enter the living room or \"/open south door\" to enter the bedroom)",
			Links: []*RoomLink{
				{
					Verb:   "/open north door",
					RoomId: "LivingRoom",
				},
				{
					Verb:   "/open south door",
					RoomId: "Bedroom",
				},
			},
			Characters: make(map[string]*character.Character),
		},
		"LivingRoom": {
			Id:          "LivingRoom",
			Description: "You have entered the living room. (type \"/open door\" to enter the hallway)",
			Links: []*RoomLink{
				{
					Verb:   "/open door",
					RoomId: "Hallway",
				},
			},
			Characters: make(map[string]*character.Character),
		},
	}
}

func (w *World) HandleCharacterJoined(c *character.Character) {
	w.Characters[c.Id] = c
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

func (w *World) HandleCharacterInput(c *character.Character, inp string) {
	subject := c.Name

	roomId := c.Room.GetId()
	room := w.Rooms[roomId]
	for _, link := range room.RoomLinks() {
		if link.Verb == inp {
			target := w.getRoomById(link.RoomId)
			if target != nil {
				w.MoveCharacter(c, target)
				return
			}
		}
	}

	verb := "say"
	args := inp
	hasArgs := true
	if inp[0:1] == "/" {
		var cmd string
		cmd, args, hasArgs = strings.Cut(inp, " ")
		verb = cmd[1:]
	}

	if verb != "say" && hasArgs && !room.ContainsCharacter(args) {
		c.SendMessage("There is no one around with that name...")
	} else {
		c.SendMessage(input.ProcessInput(subject, verb, args, subject, hasArgs))

		for id, other := range room.GetCharacters() {
			if id != c.Id {
				observer := other.Name
				other.SendMessage(input.ProcessInput(subject, verb, args, observer, hasArgs))
			}
		}
	}
}

func (w *World) MoveCharacter(c *character.Character, targetRoom *Room) {
	// Update Rooms
	currentRoom := w.Rooms[c.Room.GetId()]
	currentRoom.RemoveCharacter(c)
	targetRoom.AddCharacter(c)

	// Update Character
	c.SendMessage(c.Room.GetDescription())
}
