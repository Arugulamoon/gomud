package world

import (
	"fmt"
	"strings"

	"github.com/Arugulamoon/gomud/pkg/character"
)

type World struct {
	Commands   map[string]*Command
	Rooms      map[string]*Room
	Characters map[string]*character.Character
}

func New() *World {
	return &World{
		Commands:   make(map[string]*Command),
		Rooms:      make(map[string]*Room),
		Characters: make(map[string]*character.Character),
	}
}

func (w *World) Load() {
	w.Commands = map[string]*Command{
		GOTO:  {w},
		SAY:   {w},
		SHOUT: {w},
		TELL:  {w},
		WAVE:  {w},
		WHO:   {w},
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
		Tell(char, args)
	case WAVE:
		w.Commands[WAVE].Wave(char, args)
	case WHO:
		Who(char, args)
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
