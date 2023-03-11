package world

import (
	"fmt"
	"strings"

	"github.com/Arugulamoon/gomud/pkg/character"
	"github.com/Arugulamoon/gomud/pkg/input"
	"github.com/Arugulamoon/gomud/pkg/session"
)

type World struct {
	Characters map[string]*character.Character
	Rooms      map[string]*Room
}

func New() *World {
	return &World{
		Characters: make(map[string]*character.Character),
	}
}

// TODO: Move into yaml
func (w *World) Load() {
	w.Rooms = map[string]*Room{
		"Bedroom": {
			Id:          "Bedroom",
			Description: "You have entered your bedroom. There is a door leading out! (type \"/open door\" to leave the bedroom)",
			Links: []*RoomLink{
				{
					Verb:   "/open door",
					RoomId: "Hallway",
				},
			},
			Sessions:   make(map[string]*session.Session),
			Characters: make(map[string]*character.Character),
		},
		"Hallway": {
			Id:          "Hallway",
			Description: "You have entered a hallway with doors at either end. (type \"/open north door\" to enter the living room or \"/open south door\" to enter the bedroom)",
			Links: []*RoomLink{
				{
					Verb:   "/open north door",
					RoomId: "LivingRoom",
				},
				{
					Verb:   "/open south door",
					RoomId: "Bedroom",
				},
			},
			Sessions:   make(map[string]*session.Session),
			Characters: make(map[string]*character.Character),
		},
		"LivingRoom": {
			Id:          "LivingRoom",
			Description: "You have entered the living room. (type \"/open door\" to enter the hallway)",
			Links: []*RoomLink{
				{
					Verb:   "/open door",
					RoomId: "Hallway",
				},
			},
			Sessions:   make(map[string]*session.Session),
			Characters: make(map[string]*character.Character),
		},
	}
}

func (w *World) HandleCharacterJoined(s *session.Session) {
	w.Characters[s.Character.Id] = s.Character
	w.Rooms["Bedroom"].AddCharacter(s)

	s.WriteLine(fmt.Sprintf("Welcome %s!", s.Character.Name))
	s.WriteLine("")
	s.WriteLine(s.Character.Room.GetDescription())
}

func (w *World) HandleCharacterLeft(s *session.Session) {
	room := w.Rooms[s.Character.Room.GetId()]
	room.RemoveCharacter(s)
	delete(w.Characters, s.Character.Id)
}

func (w *World) getRoomById(id string) *Room {
	for _, r := range w.Rooms {
		if r.Id == id {
			return r
		}
	}
	return nil
}

func (w *World) HandleCharacterInput(s *session.Session, inp string) {
	subject := s.Character.Name

	roomId := s.Character.Room.GetId()
	room := w.Rooms[roomId]
	for _, link := range room.RoomLinks() {
		if link.Verb == inp {
			target := w.getRoomById(link.RoomId)
			if target != nil {
				w.MoveCharacter(s, target)
				return
			}
		}
	}

	verb := "say"
	args := inp
	hasArgs := true
	if inp[0:1] == "/" {
		var cmd string
		cmd, args, hasArgs = strings.Cut(inp, " ")
		verb = cmd[1:]
	}

	if verb != "say" && hasArgs && !room.ContainsCharacter(args) {
		s.WriteLine("There is no one around with that name...")
	} else {
		s.WriteLine(input.ProcessInput(subject, verb, args, subject, hasArgs))

		for id, other := range room.ConnectedSessions() {
			if id != s.Id {
				observer := other.Character.Name
				other.WriteLine(input.ProcessInput(subject, verb, args, observer, hasArgs))
			}
		}
	}
}

func (w *World) MoveCharacter(s *session.Session, targetRoom *Room) {
	// Update Rooms
	currentRoom := w.Rooms[s.Character.Room.GetId()]
	currentRoom.RemoveCharacter(s)
	targetRoom.AddCharacter(s)

	// Update Character
	s.WriteLine(s.Character.Room.GetDescription())
}
