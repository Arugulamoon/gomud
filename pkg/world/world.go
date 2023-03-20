package world

import (
	"fmt"
	"strings"

	"github.com/Arugulamoon/gomud/pkg/character"
)

const GOTO = "goto"
const SAY = "say"
const SHOUT = "shout"
const TELL = "tell"
const WAVE = "wave"
const WHO = "who"

type command interface {
	Perform(char *character.Character, args string)
}

type World struct {
	Id         string
	Commands   map[string]*command
	Rooms      map[string]*Room
	Characters map[string]*character.Character
}

func New(id string) *World {
	return &World{
		Id:         id,
		Commands:   make(map[string]*command),
		Rooms:      make(map[string]*Room),
		Characters: make(map[string]*character.Character),
	}
}

func (w *World) Load() {
	w.Commands = map[string]*command{
		"goto":  {w},
		"say":   {w},
		"shout": {w},
		"tell":  {w},
		"wave":  {w},
		"who":   {w},
	}

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
		w.Commands[GOTO].GoTo(char, args)
	case SAY:
		w.Commands[SAY].Say(char, args)
	case SHOUT:
		w.Commands[SHOUT].Shout(char, args)
	case TELL:
		w.Commands[TELL].Tell(char, args)
	case WAVE:
		w.Commands[WAVE].Wave(char, args)
	case WHO:
		w.Commands[WHO].Who(char, args)
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
