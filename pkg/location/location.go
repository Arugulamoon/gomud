package location

import (
	"github.com/Arugulamoon/gomud/pkg/avatar"
)

type Location struct {
	Desc  string
	Chars map[string]*avatar.Avatar
}

func New(desc string) *Location {
	return &Location{
		Desc:  desc,
		Chars: make(map[string]*avatar.Avatar),
	}
}

func (r *Location) Description() string {
	return r.Desc
}

func (r *Location) Avatars() map[string]*avatar.Avatar {
	return r.Chars
}

func (r *Location) AddAvatar(c *avatar.Avatar) {
	r.Chars[c.Id] = c
}

// func (r *Location) WelcomeAvatar(name string) {
// 	a := avatar.New(name, r)
// 	a.Welcome()
// }
