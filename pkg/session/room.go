package session

import "fmt"

type Room struct {
	Id, Desc string
	Links    []*RoomLink
	Sessions []*Session // TODO: Change to map
}

func (r *Room) SendMessage(character *Character, msg string) {
	for _, other := range r.Sessions {
		if other.User.Character != character {
			other.WriteLine(msg)
		}
	}
}

func (r *Room) AddCharacter(s *Session) {
	r.Sessions = append(r.Sessions, s)
	s.User.Character.Room = r

	r.SendMessage(s.User.Character,
		fmt.Sprintf("%s entered the room.", s.User.Character.Name))
}

// TODO: Optimize?
func (r *Room) RemoveCharacter(s *Session) {
	s.User.Character.Room = nil

	var sessions []*Session
	for _, sess := range r.Sessions {
		if sess.SessionId() != s.SessionId() {
			sessions = append(sessions, sess)
		}
	}
	r.Sessions = sessions

	r.SendMessage(s.User.Character,
		fmt.Sprintf("%s left the room.", s.User.Character.Name))
}

type RoomLink struct {
	Verb, RoomId string
}
