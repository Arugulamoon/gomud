package command

import (
	"github.com/Arugulamoon/gomud/pkg/character"
	"github.com/Arugulamoon/gomud/pkg/world"
)

type GoTo struct {
	World *world.World
}

func (cmd *GoTo) Perform(char *character.Character, targRoomId string) {
	currRoom := cmd.World.Rooms[char.RoomId]
	targRoom := currRoom.Links[targRoomId]
	if targRoom == nil {
		char.SendMessage("There is no one around with that name...")
	} else {
		currRoom.RemoveCharacter(char)
		targRoom.AddCharacter(char)
		char.RoomId = targRoom.Id
		// This done in wrong place
		char.SendMessage(targRoom.Description)
	}
}
