package command

import (
	"fmt"

	"github.com/Arugulamoon/gomud/pkg/character"
	"github.com/Arugulamoon/gomud/pkg/world"
)

type Wave struct {
	World *world.World
}

func (cmd *Wave) Perform(char *character.Character, args string) {
	room := cmd.World.Rooms[char.RoomId]
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
