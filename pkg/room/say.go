package room

import (
	"fmt"

	"github.com/Arugulamoon/gomud/pkg/character"
)

type Say struct {
	Room    *Room
	Speaker *character.Character
	Message string
	Hearers map[string]*character.Character
}

func NewSay(r *Room, spkr *character.Character, msg string) *Say {
	return &Say{
		Room:    r,
		Speaker: spkr,
		Message: msg,
		Hearers: r.GetCharacters(),
	}
}

func (cmd *Say) Perform() error {
	if cmd.Message == "" {
		cmd.Speaker.SendMessage("Cannot send empty message...")
		return fmt.Errorf("empty message")
	}

	cmd.Speaker.SendMessage(fmt.Sprintf("You say, \"%s\"", cmd.Message))
	cmd.Room.BroadcastMessage(cmd.Speaker.Name,
		fmt.Sprintf("%s said, \"%s\"", cmd.Speaker.Name, cmd.Message))

	return nil
}
