package room

import (
	"github.com/Arugulamoon/gomud/pkg/character"
)

type Room struct {
	Desc  string
	Chars map[string]*character.Character
}

func New(desc string) *Room {
	return &Room{
		Desc:  desc,
		Chars: make(map[string]*character.Character),
	}
}

func (r *Room) Description() string {
	return r.Desc
}

func (r *Room) Characters() map[string]*character.Character {
	return r.Chars
}

func (r *Room) AddCharacter(c *character.Character) {
	r.Chars[c.Id] = c
}

// func (r *Room) WelcomeCharacter(name string) {
// 	char := character.New(name, r)
// 	char.Welcome()
// }
