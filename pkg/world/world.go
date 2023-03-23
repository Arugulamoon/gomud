package world

import (
	"fmt"
	"strings"

	"github.com/Arugulamoon/gomud/pkg/character"
)

// Commands
const GOTO = "goto"
const SAY = "say"
const SHOUT = "shout"
const TELL = "tell"
const WAVE = "wave"
const WHO = "who"

type World struct {
	Id         string
	Rooms      map[string]*Room
	Characters map[string]*character.Character
}

func New(id string) *World {
	return &World{
		Id:         id,
		Rooms:      make(map[string]*Room),
		Characters: make(map[string]*character.Character),
	}
}

func (w *World) Load() {
	bedroom := NewRoom("Bedroom", "You have entered your bedroom. There is a door leading out! (type \"/goto Hallway\" to leave the bedroom)")
	hallway := NewRoom("Hallway", "You have entered a hallway with doors at either end. (type \"/goto LivingRoom\" to enter the living room or \"/goto Bedroom\" to enter the bedroom)")
	livingRoom := NewRoom("LivingRoom", "You have entered the living room. (type \"/goto Hallway\" to enter the hallway)")
	bedroom.Links[hallway.Id] = hallway
	hallway.Links[bedroom.Id] = bedroom
	livingRoom.Links[hallway.Id] = hallway
	hallway.Links[livingRoom.Id] = livingRoom
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

func (w *World) HandleCharacterJoined(c *character.Character) {
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
	c.SendMessage(r.Description)
}

func (w *World) HandleCharacterLeft(c *character.Character) {
	// Update Room
	r := w.Rooms[c.RoomId]
	r.RemoveCharacter(c)

	// Update World
	w.RemoveCharacter(c)
}

func (w *World) HandleCharacterInput(char *character.Character, inp string) {
	cmd, args := splitCommandAndArgs(char, inp)

	switch cmd {
	case GOTO:
		w.goTo(char, args)
	case SAY:
		w.say(char, args)
	case SHOUT:
		w.shout(char, args)
	case TELL:
		w.tell(char, args)
	case WAVE:
		w.wave(char, args)
	case WHO:
		w.who(char, args)
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

func (w *World) goTo(char *character.Character, targRoomId string) {
	currRoom := w.Rooms[char.RoomId]
	targRoom := currRoom.Links[targRoomId]
	if targRoom == nil {
		char.SendMessage("There is no one around with that name...")
	} else {
		currRoom.RemoveCharacter(char)
		targRoom.AddCharacter(char)
		char.RoomId = targRoom.Id
		char.SendMessage(targRoom.Description)
	}
}

func (w *World) say(char *character.Character, msg string) {
	if msg == "" {
		char.SendMessage("Cannot send empty message...")
	} else {
		char.SendMessage(fmt.Sprintf("You say, \"%s\"", msg))
		room := w.Rooms[char.RoomId]
		room.BroadcastMessage(char.Name, fmt.Sprintf("%s said, \"%s\"", char.Name, msg))
	}
}

func (w *World) shout(char *character.Character, msg string) {
	if msg == "" {
		char.SendMessage("Cannot send empty message...")
	} else {
		char.SendMessage(fmt.Sprintf("You shout, \"%s\"", msg))
		w.BroadcastMessage(char.Name, fmt.Sprintf("%s shouted, \"%s\"", char.Name, msg))
	}
}

func (w *World) tell(char *character.Character, args string) {
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
}

func (w *World) wave(char *character.Character, args string) {
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
}

func (w *World) who(char *character.Character, args string) {
	room := w.Rooms[char.RoomId]
	if args != "" {
		if args == "all" {
			char.SendMessage("/who all:")
			for _, name := range w.Characters {
				char.SendMessage(fmt.Sprintf("  %s", name))
			}
		}
	} else {
		char.SendMessage("/who:")
		for _, name := range room.Characters {
			char.SendMessage(fmt.Sprintf("  %s", name))
		}
	}
}
