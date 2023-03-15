package character

import (
	"fmt"
	"math/rand"
)

type session interface {
	WriteLine(msg string) error
}

type world interface {
	GetCharacterNames() []string
	BroadcastMessage(speaker, msg string)
}

type room interface {
	GetId() string
	GetDescription() string
	GetCharacterNames() []string
	BroadcastMessage(speaker, msg string)
}

type Character struct {
	Id, Name string
	Session  session
	World    world
	Room     room
}

func New(s session) *Character {
	return &Character{
		Id:      generateId(),
		Name:    generateName(),
		Session: s,
	}
}

func (c *Character) SendMessage(msg string) {
	c.Session.WriteLine(msg)
}

func (c *Character) Shout(msg string) {
	c.Session.WriteLine(fmt.Sprintf("You shout, \"%s\"", msg))
	c.World.BroadcastMessage(c.Name, fmt.Sprintf("%s shouted, \"%s\"", c.Name, msg))
}

func (c *Character) Say(msg string) {
	c.Session.WriteLine(fmt.Sprintf("You say, \"%s\"", msg))
	c.Room.BroadcastMessage(c.Name, fmt.Sprintf("%s said, \"%s\"", c.Name, msg))
}

func (c *Character) Wave() {
	c.Session.WriteLine("You wave.")
	c.Room.BroadcastMessage("%s waved.", c.Name)
}

func (c *Character) WaveAtTarget(target string) {
	c.Session.WriteLine(fmt.Sprintf("You wave at %s.", target))
	c.Room.BroadcastMessage(c.Name, fmt.Sprintf("%s waved at %s.", c.Name, target))
}

func (c *Character) WhoAll() {
	c.Session.WriteLine("/who all:")
	for _, name := range c.World.GetCharacterNames() {
		c.Session.WriteLine(fmt.Sprintf("  %s", name))
	}
}

func (c *Character) Who() {
	c.Session.WriteLine("/who:")
	for _, name := range c.Room.GetCharacterNames() {
		c.Session.WriteLine(fmt.Sprintf("  %s", name))
	}
}

func (c *Character) Tell(target *Character, msg string) {
	c.Session.WriteLine(fmt.Sprintf("You tell %s, \"%s\"", target, msg))
	target.SendMessage(fmt.Sprintf("%s tells you, \"%s\"", c.Name, msg))
}

var nextId = 1

func generateId() string {
	var id = nextId
	nextId++
	return fmt.Sprintf("%d", id)
}

func generateName() string {
	return fmt.Sprintf("Character %d", rand.Intn(100)+1)
}
