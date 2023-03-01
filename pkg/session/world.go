package session

import (
	"fmt"
	"strings"
)

type World struct {
	Rooms []*Room
}

func NewWorld() *World {
	return &World{}
}

// TODO: Move into yaml
func (w *World) Init() {
	w.Rooms = []*Room{
		{
			Id:   "Bedroom",
			Desc: "You have entered your bedroom. There is a door leading out! (type \"/open door\" to leave the bedroom)",
			Links: []*RoomLink{
				{
					Verb:   "/open door",
					RoomId: "Hallway",
				},
			},
			Sessions: make(map[string]*Session),
		},
		{
			Id:   "Hallway",
			Desc: "You have entered a hallway with doors at either end. (type \"/open north door\" to enter the living room or \"/open south door\" to enter the bedroom)",
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
			Sessions: make(map[string]*Session),
		},
		{
			Id:   "LivingRoom",
			Desc: "You have entered the living room. (type \"/open door\" to enter the hallway)",
			Links: []*RoomLink{
				{
					Verb:   "/open door",
					RoomId: "Hallway",
				},
			},
			Sessions: make(map[string]*Session),
		},
	}
}

func (w *World) HandleCharacterJoined(s *Session) {
	w.Rooms[0].AddCharacter(s)

	s.WriteLine(fmt.Sprintf("Welcome %s!", s.User.Character.Name))
	s.WriteLine("")
	s.WriteLine(s.User.Character.Room.Desc)
}

func (w *World) HandleCharacterLeft(s *Session) {
	s.User.Character.Room.RemoveCharacter(s) // Weird; char removing self from room
}

func (w *World) getRoomById(id string) *Room {
	for _, r := range w.Rooms {
		if r.Id == id {
			return r
		}
	}
	return nil
}

func (w *World) HandleCharacterInput(s *Session, input string) {
	subject := s.User.Character.Name

	room := s.User.Character.Room
	for _, link := range room.Links {
		if link.Verb == input {
			target := w.getRoomById(link.RoomId)
			if target != nil {
				w.MoveCharacter(s, target)
				return
			}
		}
	}

	verb := "say"
	args := input
	hasArgs := true
	if input[0:1] == "/" {
		var cmd string
		cmd, args, hasArgs = strings.Cut(input, " ")
		verb = cmd[1:]
	}

	if verb != "say" && hasArgs && !room.ContainsCharacter(args) {
		s.WriteLine("There is no one around with that name...")
	} else {
		s.WriteLine(ProcessInput(subject, verb, args, subject, hasArgs))

		for id, other := range s.User.Character.Room.Sessions {
			if id != s.Id {
				observer := other.User.Character.Name
				other.WriteLine(ProcessInput(subject, verb, args, observer, hasArgs))
			}
		}
	}
}

func (w *World) MoveCharacter(s *Session, targetRoom *Room) {
	// Update Rooms
	currentRoom := s.User.Character.Room
	currentRoom.RemoveCharacter(s)
	targetRoom.AddCharacter(s)

	// Update Character
	s.WriteLine(s.User.Character.Room.Desc)
}
