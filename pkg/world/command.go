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

type Command struct {
	World *World
}

func (cmd *Command) GoTo(char *character.Character, targRoomId string) {
	currRoom := cmd.World.Rooms[char.Room.GetId()]
	targRoom := currRoom.Links[targRoomId]
	if targRoom == nil {
		char.SendMessage("There is no one around with that name...")
	} else {
		currRoom.RemoveCharacter(char)
		targRoom.AddCharacter(char)
		char.SendMessage(char.Room.GetDescription())
	}
}

func (cmd *Command) Say(char *character.Character, msg string) {
	if msg == "" {
		char.SendMessage("Cannot send empty message...")
	} else {
		char.SendMessage(fmt.Sprintf("You say, \"%s\"", msg))
		room := cmd.World.Rooms[char.Room.GetId()]
		room.BroadcastMessage(char.Name, fmt.Sprintf("%s said, \"%s\"", char.Name, msg))
	}
}

func (cmd *Command) Shout(char *character.Character, msg string) {
	if msg == "" {
		char.SendMessage("Cannot send empty message...")
	} else {
		char.SendMessage(fmt.Sprintf("You shout, \"%s\"", msg))
		cmd.World.BroadcastMessage(char.Name, fmt.Sprintf("%s shouted, \"%s\"", char.Name, msg))
	}
}

func Tell(char *character.Character, args string) {
	if args != "" {
		targetId, msg, _ := strings.Cut(args, " ")
		if msg != "" {
			if char.World.ContainsCharacter(targetId) {
				target := char.World.GetCharacters()[targetId]
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

func (cmd *Command) Wave(char *character.Character, args string) {
	room := cmd.World.Rooms[char.Room.GetId()]
	if args != "" {
		if char.Room.ContainsCharacter(args) {
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

func Who(char *character.Character, args string) {
	if args != "" {
		if args == "all" {
			char.SendMessage("/who all:")
			for _, name := range char.World.GetCharacterNames() {
				char.SendMessage(fmt.Sprintf("  %s", name))
			}
		}
	} else {
		char.SendMessage("/who:")
		for _, name := range char.Room.GetCharacterNames() {
			char.SendMessage(fmt.Sprintf("  %s", name))
		}
	}
}
