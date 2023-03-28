package world

import (
	"log"
	"strings"

	"github.com/Arugulamoon/gomud/pkg/character"
	"github.com/Arugulamoon/gomud/pkg/room"
	"github.com/Arugulamoon/gomud/pkg/session"
)

type SessionHandler struct {
	World        *World
	EventChannel <-chan session.SessionEvent
}

func NewSessionHandler(w *World, ch <-chan session.SessionEvent) *SessionHandler {
	return &SessionHandler{
		World:        w,
		EventChannel: ch,
	}
}

func (h *SessionHandler) Start() {
	for sessionEvent := range h.EventChannel {
		char := sessionEvent.Session.Character
		r := h.World.Rooms[char.RoomId]

		switch event := sessionEvent.Event.(type) {

		case *session.SessionCreateEvent:
			go h.World.HandleCharacterJoined(char)
			go r.HandleCharacterJoined(char)

		case *session.SessionDisconnectEvent:
			go h.World.HandleCharacterLeft(char)
			go r.HandleCharacterLeft(char)

		case *session.SessionInputEvent:
			h.handleCharacterInput(char, event.Input)
		}
	}
}

func (h *SessionHandler) handleCharacterInput(char *character.Character, inp string) {
	cmd, args := splitCommandAndArgs(char, inp)

	var err error
	switch cmd {

	// World
	case GOTO:
		err = h.World.goTo(char, args)

	case MOTD:
		h.World.motd(char)

	case SHOUT:
		err = h.World.shout(char, args)

	case TELL:
		err = h.World.tell(char, args)

	// World and Room
	case WHO:
		if args == "all" {
			h.World.who(char)
		}
		r := h.World.Rooms[char.RoomId]
		r.Who(char)

	// Room
	case room.PICKUP:
		r := h.World.Rooms[char.RoomId]
		err = r.PickUp(char, args)

	case room.SAY:
		r := h.World.Rooms[char.RoomId]
		err = r.Say(char, args)

	case room.WAVE:
		r := h.World.Rooms[char.RoomId]
		err = r.Wave(char, args)

		// Character
	case character.CHAR:
		char.Char()

	}

	if err != nil {
		log.Println(err)
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
