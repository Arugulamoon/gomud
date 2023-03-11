package avatar

import (
	"fmt"
)

type room interface {
	Description() string
	Avatars() map[string]*Avatar
	AddAvatar(*Avatar)
}

type Avatar struct {
	Id, Name string
	Room     room
}

func New(name string, r room) *Avatar {
	return &Avatar{
		Name: name,
		Room: r,
	}
}

func (c *Avatar) EnterRoom() string {
	c.Room.AddAvatar(c)
	return c.Look()
}

func (c *Avatar) Welcome() string {
	return fmt.Sprintf("Hello %s!\n", c.Name)
}

func (c *Avatar) Look() string {
	return c.Room.Description()
}

func (c *Avatar) Who() map[string]*Avatar {
	return c.Room.Avatars()
}
