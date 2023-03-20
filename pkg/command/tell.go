package command

import (
	"fmt"
	"strings"

	"github.com/Arugulamoon/gomud/pkg/character"
	"github.com/Arugulamoon/gomud/pkg/world"
)

type Tell struct {
	World *world.World
}

func (cmd *Tell) Perform(char *character.Character, args string) {
	if args != "" {
		targetName, msg, _ := strings.Cut(args, " ")
		if msg != "" {
			if cmd.World.ContainsCharacter(targetName) {
				target := cmd.World.GetCharacters()[targetName]
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
