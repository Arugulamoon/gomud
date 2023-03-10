package room

import (
	"github.com/Arugulamoon/gomud/pkg/character"
)

type Room struct {
	Desc  string
	Chars []*character.Character
}

func New(desc string) *Room {
	return &Room{
		Desc:  desc,
		Chars: make([]*character.Character, 0),
	}
}

func (r *Room) Description() string {
	return r.Desc
}

func (r *Room) Characters() []*character.Character {
	return r.Chars
}

func (r *Room) AddCharacter(c *character.Character) {
	r.Chars = append(r.Chars, c)
}

// func (r *Room) WelcomeCharacter(name string) {
// 	char := character.New(name, r)
// 	char.Welcome()
// }
