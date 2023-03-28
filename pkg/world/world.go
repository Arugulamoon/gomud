package world

import (
	"fmt"
	"strings"

	"github.com/Arugulamoon/gomud/pkg/character"
	"github.com/Arugulamoon/gomud/pkg/item"
	"github.com/Arugulamoon/gomud/pkg/room"
)

// Commands
const GOTO = "goto"
const PICKUP = "pickup"
const SAY = "say"
const SHOUT = "shout"
const TELL = "tell"
const WAVE = "wave"
const WHO = "who"

type World struct {
	Id         string
	Rooms      map[string]*room.Room
	Characters map[string]*character.Character
}

func New(id string) *World {
	return &World{
		Id:         id,
		Rooms:      make(map[string]*room.Room),
		Characters: make(map[string]*character.Character),
	}
}

func (w *World) Load() {
	bedroom := room.NewRoom("Bedroom", "You have entered your bedroom. There is a door leading out! (type \"/goto Hallway\" to leave the bedroom)")
	hallway := room.NewRoom("Hallway", "You have entered a hallway with doors at either end. (type \"/goto LivingRoom\" to enter the living room or \"/goto Bedroom\" to enter the bedroom)")
	livingRoom := room.NewRoom("LivingRoom", "You have entered the living room. (type \"/goto Hallway\" to enter the hallway)")
	bedroom.Paths[hallway.Id] = hallway
	hallway.Paths[bedroom.Id] = bedroom
	livingRoom.Paths[hallway.Id] = hallway
	hallway.Paths[livingRoom.Id] = livingRoom
	book := item.New("Book")
	bedroom.Items[book.Id] = book

	w.Rooms[bedroom.Id] = bedroom
	w.Rooms[hallway.Id] = hallway
	w.Rooms[livingRoom.Id] = livingRoom
}

func (w *World) GetCharacters() map[string]*character.Character {
	return w.Characters
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

func (w *World) AddCharacter(c *character.Character) {
	w.Characters[c.Id] = c
	w.BroadcastMessage(c.Name, fmt.Sprintf("%s entered the world.", c.Name))
}

func (w *World) RemoveCharacter(c *character.Character) {
	delete(w.Characters, c.Id)
	w.BroadcastMessage(c.Name, fmt.Sprintf("%s left the world.", c.Name))
}

func (w *World) HandleCharacterJoined(c *character.Character) error {
	// Update World
	c.WorldId = w.Id
	w.AddCharacter(c)
	c.SendMessage(fmt.Sprintf("Welcome %s!", c.Name))

	// Update Room
	if c.RoomId == "" {
		c.RoomId = "Bedroom" // Make const
	}
	r := w.Rooms[c.RoomId]
	r.AddCharacter(c)
	r.Describe(c)

	return nil
}

func (w *World) HandleCharacterLeft(c *character.Character) error {
	// Update Room
	r := w.Rooms[c.RoomId]
	r.RemoveCharacter(c)

	// Update World
	w.RemoveCharacter(c)

	return nil
}

func (w *World) HandleCharacterInput(
	char *character.Character, inp string) error {

	cmd, args := splitCommandAndArgs(char, inp)

	switch cmd {
	case GOTO: // Move to Room
		return w.goTo(char, args)
	case PICKUP:
		return w.pickUp(char, args)
	case SAY: // Move to Room
		return w.say(char, args)
	case SHOUT:
		return w.shout(char, args)
	case TELL:
		return w.tell(char, args)
	case WAVE: // Move to Room
		return w.wave(char, args)
	case WHO: // Also introduce room.Who()
		return w.who(char, args)
	default:
		return fmt.Errorf("DEV: missing command method")
	}
}

func splitCommandAndArgs(c *character.Character, input string) (string, string) {
	cmd := "/say"
	args := input
	if input[0:1] == "/" { // if first char is slash
		cmd, args, _ = strings.Cut(input, " ")
	}
	return cmd[1:], args
}

func (w *World) goTo(char *character.Character, targRoomId string) error {
	currRoom := w.Rooms[char.RoomId]
	targRoom := currRoom.Paths[targRoomId]
	if targRoom == nil {
		char.SendMessage("There is no one around with that name...")
	} else {
		currRoom.RemoveCharacter(char)
		targRoom.AddCharacter(char)
		char.RoomId = targRoom.Id
		targRoom.Describe(char)
	}

	return nil
}

func (w *World) pickUp(char *character.Character, itemId string) error {
	r := w.Rooms[char.RoomId]
	item := r.Items[itemId]
	if item == nil {
		char.SendMessage("There is no item around with that name...")
	} else {
		return char.PickUp(itemId)
	}

	return nil
}

func (w *World) say(char *character.Character, msg string) error {
	r := w.Rooms[char.RoomId]
	if r == nil {
		return fmt.Errorf("room not found")
	}
	say := room.NewSay(r, char, msg)
	return say.Perform()
}

func (w *World) shout(char *character.Character, msg string) error {
	if msg == "" {
		char.SendMessage("Cannot send empty message...")
	} else {
		char.SendMessage(fmt.Sprintf("You shout, \"%s\"", msg))
		w.BroadcastMessage(char.Name, fmt.Sprintf("%s shouted, \"%s\"", char.Name, msg))
	}

	return nil
}

func (w *World) tell(char *character.Character, args string) error {
	if args != "" {
		targetName, msg, _ := strings.Cut(args, " ")
		if msg != "" {
			if w.ContainsCharacter(targetName) {
				target := w.GetCharacters()[targetName]
				char.SendMessage(fmt.Sprintf("You tell %s, \"%s\"", target.Name, msg))
				target.SendMessage(fmt.Sprintf("%s tells you, \"%s\"", char.Name, msg))
			} else {
				char.SendMessage("There is no one around with that name...")
			}
		} else {
			char.SendMessage("Cannot send empty message...")
		}
	} else {
		char.SendMessage("Cannot send empty target and message...")
	}

	return nil
}

func (w *World) wave(char *character.Character, args string) error {
	room := w.Rooms[char.RoomId]
	if args != "" {
		if room.ContainsCharacter(args) {
			char.SendMessage(fmt.Sprintf("You wave at %s.", args))
			room.BroadcastMessage(char.Name, fmt.Sprintf("%s waved at %s.", char.Name, args))
		} else {
			char.SendMessage("There is no one around with that name...")
		}
	} else {
		char.SendMessage("You wave.")
		room.BroadcastMessage(char.Name, fmt.Sprintf("%s waved.", char.Name))
	}

	return nil
}

func (w *World) who(char *character.Character, args string) error {
	room := w.Rooms[char.RoomId]
	if args != "" {
		if args == "all" {
			char.SendMessage("/who all:")
			for _, c := range w.Characters {
				char.SendMessage(fmt.Sprintf("  %s", c.Name))
			}
		}
	} else {
		char.SendMessage("/who:")
		for _, c := range room.Characters {
			char.SendMessage(fmt.Sprintf("  %s", c.Name))
		}
	}

	return nil
}
