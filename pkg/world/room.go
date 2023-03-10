package world

import "fmt"

type Room struct {
	Id, Desc string
	Links    []*RoomLink
	Sessions map[string]*Session
}

func (r *Room) Description() string {
	return r.Desc
}

func (r *Room) RoomLinks() []*RoomLink {
	return r.Links
}

func (r *Room) ConnectedSessions() map[string]*Session {
	return r.Sessions
}

func (r *Room) SendMessage(s *Session, msg string) {
	for id, other := range r.Sessions {
		if id != s.Id {
			other.WriteLine(msg)
		}
	}
}

func (r *Room) ContainsCharacter(name string) bool {
	for _, s := range r.Sessions {
		if s.Character.Name == name {
			return true
		}
	}
	return false
}

func (r *Room) AddCharacter(s *Session) {
	r.Sessions[s.Id] = s
	s.Character.Room = r
	r.SendMessage(s, fmt.Sprintf("%s entered the room.", s.Character.Name))
}

func (r *Room) RemoveCharacter(s *Session) {
	delete(r.Sessions, s.Id)
	s.Character.Room = nil
	r.SendMessage(s, fmt.Sprintf("%s left the room.", s.Character.Name))
}

type RoomLink struct {
	Verb, RoomId string
}
