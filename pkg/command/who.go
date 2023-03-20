package command

import (
	"fmt"

	"github.com/Arugulamoon/gomud/pkg/character"
	"github.com/Arugulamoon/gomud/pkg/world"
)

type Who struct {
	World *world.World
}

func (cmd *Who) Perform(char *character.Character, args string) {
	room := cmd.World.Rooms[char.RoomId]
	if args != "" {
		if args == "all" {
			char.SendMessage("/who all:")
			for _, name := range cmd.World.Characters {
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
