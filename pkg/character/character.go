package character

import (
	"fmt"
)

type room interface {
	Description() string
	Characters() []*Character
	AddCharacter(*Character)
}

type Character struct {
	Name string
	Room room
}

func New(name string, r room) *Character {
	return &Character{
		Name: name,
		Room: r,
	}
}

func (c *Character) EnterRoom() string {
	c.Room.AddCharacter(c)
	return c.Look()
}

func (c *Character) Welcome() string {
	return fmt.Sprintf("Hello %s!\n", c.Name)
}

func (c *Character) Look() string {
	return c.Room.Description()
}

func (c *Character) Who() []*Character {
	return c.Room.Characters()
}
