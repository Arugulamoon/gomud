package main

import (
	"fmt"

	"github.com/Arugulamoon/gomud/pkg/character"
	"github.com/Arugulamoon/gomud/pkg/room"
)

func main() {
	r := room.New("You have entered your bedroom. There is a door leading out! (type \"/open door\" to leave the bedroom)")
	c := character.New("Arugulamoon", r)
	r.AddCharacter(c)
	c.Welcome()

	fmt.Println(c.Look())

	fmt.Println("Characters:")
	for _, character := range c.Who() {
		fmt.Println(character.Name)
	}
	// r.WelcomeCharacter("Arugulamoon")
	// r.Description()
}
