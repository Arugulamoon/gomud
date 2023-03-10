package character

import (
	"fmt"
)

type room interface {
	Description() string
	Characters() []*Character
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

func (c *Character) Welcome() {
	fmt.Printf("Hello %s!\n", c.Name)
}

func (c *Character) Look() string {
	return c.Room.Description()
}

func (c *Character) Who() []*Character {
	return c.Room.Characters()
}
