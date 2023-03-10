package main

import (
	"fmt"

	"github.com/Arugulamoon/gomud/pkg/character"
	"github.com/Arugulamoon/gomud/pkg/room"
)

func main() {
	r := room.New("You have entered your bedroom. There is a door leading out! (type \"/open door\" to leave the bedroom)")

	c := character.New("Arugulamoon", r)
	// r.WelcomeCharacter("Arugulamoon")
	fmt.Println(c.Welcome())

	fmt.Println(c.EnterRoom())
	// r.Description()

	fmt.Println("Characters:")
	for _, character := range c.Who() {
		fmt.Println(character.Name)
	}
}
