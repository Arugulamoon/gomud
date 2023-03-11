package location

import (
	"github.com/Arugulamoon/gomud/pkg/avatar"
)

type Room struct {
	Desc  string
	Chars map[string]*avatar.Avatar
}

func New(desc string) *Room {
	return &Room{
		Desc:  desc,
		Chars: make(map[string]*avatar.Avatar),
	}
}

func (r *Room) Description() string {
	return r.Desc
}

func (r *Room) Avatars() map[string]*avatar.Avatar {
	return r.Chars
}

func (r *Room) AddAvatar(c *avatar.Avatar) {
	r.Chars[c.Id] = c
}

// func (r *Room) WelcomeAvatar(name string) {
// 	a := avatar.New(name, r)
// 	a.Welcome()
// }
