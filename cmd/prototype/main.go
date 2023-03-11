package main

import (
	"fmt"

	"github.com/Arugulamoon/gomud/pkg/avatar"
	"github.com/Arugulamoon/gomud/pkg/location"
)

func main() {
	r := location.New("You have entered your bedroom. There is a door leading out! (type \"/open door\" to leave the bedroom)")

	c := avatar.New("Arugulamoon", r)
	// r.WelcomeAvatar("Arugulamoon")
	fmt.Println(c.Welcome())

	fmt.Println(c.EnterRoom())
	// r.Description()

	fmt.Println("Avatars:")
	for _, avatar := range c.Who() {
		fmt.Println(avatar.Name)
	}
}
