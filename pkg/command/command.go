package command

import (
	"fmt"
	"strings"

	"github.com/Arugulamoon/gomud/pkg/character"
)

const SAY = "say"
const SHOUT = "shout"
const TELL = "tell"
const WAVE = "wave"
const WHO = "who"

type Command struct{}

func Say(char *character.Character, msg string) {
	if msg == "" {
		char.SendMessage("Cannot send empty message...")
	} else {
		char.SendMessage(fmt.Sprintf("You say, \"%s\"", msg))
		char.Room.BroadcastMessage(char.Name, fmt.Sprintf("%s said, \"%s\"", char.Name, msg))
	}
}

func Shout(char *character.Character, msg string) {
	if msg == "" {
		char.SendMessage("Cannot send empty message...")
	} else {
		char.SendMessage(fmt.Sprintf("You shout, \"%s\"", msg))
		char.World.BroadcastMessage(char.Name, fmt.Sprintf("%s shouted, \"%s\"", char.Name, msg))
	}
}

func Tell(char *character.Character, args string) {
	if args != "" {
		targetName, msg, _ := strings.Cut(args, " ")
		if msg != "" {
			if char.World.ContainsCharacter(targetName) {
				target := char.World.GetCharacters()[targetName]
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

func Wave(char *character.Character, args string) {
	if args != "" {
		if char.Room.ContainsCharacter(args) {
			char.SendMessage(fmt.Sprintf("You wave at %s.", args))
			char.Room.BroadcastMessage(char.Name, fmt.Sprintf("%s waved at %s.", char.Name, args))
		} else {
			char.SendMessage("There is no one around with that name...")
		}
	} else {
		char.SendMessage("You wave.")
		char.Room.BroadcastMessage(char.Name, fmt.Sprintf("%s waved.", char.Name))
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
