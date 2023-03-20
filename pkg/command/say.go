package command

import (
	"fmt"

	"github.com/Arugulamoon/gomud/pkg/character"
	"github.com/Arugulamoon/gomud/pkg/world"
)

type Say struct {
	World *world.World
}

func (cmd *Say) Perform(char *character.Character, msg string) {
	if msg == "" {
		char.SendMessage("Cannot send empty message...")
	} else {
		char.SendMessage(fmt.Sprintf("You say, \"%s\"", msg))
		room := cmd.World.Rooms[char.RoomId]
		room.BroadcastMessage(char.Name, fmt.Sprintf("%s said, \"%s\"", char.Name, msg))
	}
}
