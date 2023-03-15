package world

import (
	"fmt"
	"strings"

	"github.com/Arugulamoon/gomud/pkg/character"
)

type World struct {
	Characters map[string]*character.Character
	Rooms      map[string]*Room

	// CharactersToRooms map[string]string
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
	verb, hasArgs, args := handleCharacterInput(char, inp)

	switch verb {

	case character.SAY:
		if hasArgs {
			char.Say(args)
		} else {
			char.SendMessage("Cannot send empty message...")
		}

	case character.SHOUT:
		if hasArgs {
			char.Shout(args)
		} else {
			char.SendMessage("Cannot send empty message...")
		}

	case character.TELL:
		if hasArgs {
			target, msg, hasMsg := strings.Cut(args, " ")
			if hasMsg {
				if w.containsCharacter(target) {
					char.Tell(w.Characters[target], msg)
				} else {
					char.SendMessage("There is no one around with that name...")
				}
			} else {
				char.SendMessage("Cannot send empty message...")
			}
		} else {
			char.SendMessage("Cannot send empty target and message...")
		}

	case character.WAVE:
		if hasArgs {
			if room.ContainsCharacter(args) {
				char.WaveAtTarget(args)
			} else {
				char.SendMessage("There is no one around with that name...")
			}
		} else {
			char.Wave()
		}

	case character.WHO:
		if hasArgs {
			if args == "all" {
				char.WhoAll()
			}
		} else {
			char.Who()
		}
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

func handleCharacterInput(c *character.Character, input string) (string, bool, string) {
	verb := "say"
	hasArgs := true
	args := input
	if input[0:1] == "/" {
		var cmd string
		cmd, args, hasArgs = strings.Cut(input, " ")
		verb = cmd[1:]
	}
	return verb, hasArgs, args
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

func (w *World) containsCharacter(name string) bool {
	for _, character := range w.Characters {
		if character.Name == name {
			return true
		}
	}
	return false
}
