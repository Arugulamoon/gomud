package world

import (
	"fmt"

	"github.com/Arugulamoon/gomud/pkg/character"
	"github.com/Arugulamoon/gomud/pkg/session"
)

type Room struct {
	Id, Description string
	Links           []*RoomLink
	Characters      map[string]*character.Character
}

func (r *Room) GetId() string {
	return r.Id
}

func (r *Room) GetDescription() string {
	return r.Description
}

func (r *Room) RoomLinks() []*RoomLink {
	return r.Links
}

func (r *Room) GetCharacters() map[string]*character.Character {
	return r.Characters
}

func (r *Room) SendMessage(s *session.Session, msg string) {
	for id, other := range r.Characters {
		if id != s.Character.Id {
			other.Session.WriteLine(msg)
		}
	}
}

func (r *Room) ContainsCharacter(name string) bool {
	for _, character := range r.Characters {
		if character.Name == name {
			return true
		}
	}
	return false
}

func (r *Room) AddCharacter(s *session.Session) {
	r.Characters[s.Character.Id] = s.Character
	s.Character.Room = r
	r.SendMessage(s, fmt.Sprintf("%s entered the room.", s.Character.Name))
}

func (r *Room) RemoveCharacter(s *session.Session) {
	delete(r.Characters, s.Character.Id)
	s.Character.Room = nil
	r.SendMessage(s, fmt.Sprintf("%s left the room.", s.Character.Name))
}

type RoomLink struct {
	Verb, RoomId string
}
