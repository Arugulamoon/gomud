package room

import (
	"fmt"

	"github.com/Arugulamoon/gomud/pkg/character"
	"github.com/Arugulamoon/gomud/pkg/item"
)

type Room struct {
	Id, Description string

	Paths map[string]*Room
	Items map[string]*item.Item

	Characters map[string]*character.Character
}

func New(id, desc string) *Room {
	return &Room{
		Id:          id,
		Description: desc,
		Paths:       make(map[string]*Room),
		Items:       make(map[string]*item.Item),
		Characters:  make(map[string]*character.Character),
	}
}

// Characters
func (r *Room) GetCharacters() map[string]*character.Character {
	return r.Characters
}

func (r *Room) FindCharacter(name string) *character.Character {
	for _, char := range r.Characters {
		if char.Name == name {
			return char
		}
	}
	return nil
}

// TODO: Remove once new tests written and can delete test for this function
func (r *Room) ContainsCharacter(name string) bool {
	return r.FindCharacter(name) != nil
}

func (r *Room) addCharacter(char *character.Character) {
	r.Characters[char.Id] = char
}

func (r *Room) removeCharacter(char *character.Character) {
	delete(r.Characters, char.Id)
}

// Messaging
func (r *Room) BroadcastMessage(speaker, msg string) {
	for _, char := range r.Characters {
		if char.Name != speaker {
			char.SendMessage(msg)
		}
	}
}

// Events
func (r *Room) Join(char *character.Character) {
	r.addCharacter(char)
	r.BroadcastMessage(char.Name, fmt.Sprintf("%s entered the room.", char.Name))
	char.RoomId = r.Id
}

func (r *Room) Leave(char *character.Character) {
	r.removeCharacter(char)
	r.BroadcastMessage(char.Name, fmt.Sprintf("%s left the room.", char.Name))
}

// Handlers
func (r *Room) HandleCharacterJoined(char *character.Character) {
	r.Join(char)
	r.Look(char)
}

func (r *Room) HandleCharacterLeft(char *character.Character) {
	r.Leave(char)
}

// Available User Commands
const PICKUP = "pickup"
const SAY = "say"
const WAVE = "wave"
const WHO = "who" // not needed

func (r *Room) Look(char *character.Character) {
	char.SendMessage(r.Description)
	if len(r.Items) > 0 {
		char.SendMessage("Items:")
		for _, item := range r.Items {
			char.SendMessage(fmt.Sprintf("  %s", item.Id))
		}
	}
}

func (r *Room) PickUp(char *character.Character, itemId string) error {
	item := r.Items[itemId]
	if item == nil {
		char.SendMessage("There is no item around with that name...")
		return fmt.Errorf("item not found")
	}

	return char.PickUp(itemId)
}

func (r *Room) Say(char *character.Character, msg string) error {
	if msg == "" {
		char.SendMessage("Cannot send empty message...")
		return fmt.Errorf("message empty")
	}

	char.SendMessage(fmt.Sprintf("You say, \"%s\"", msg))
	r.BroadcastMessage(char.Name, fmt.Sprintf("%s said, \"%s\"", char.Name, msg))
	return nil
}

func (r *Room) Wave(char *character.Character, args string) error {
	if args == "" {
		char.SendMessage("You wave.")
		r.BroadcastMessage(char.Name, fmt.Sprintf("%s waved.", char.Name))
	}

	targ := r.FindCharacter(args)
	if targ == nil {
		char.SendMessage("There is no one around with that name...")
		return fmt.Errorf("target character not found")
	}

	char.SendMessage(fmt.Sprintf("You wave at %s.", targ.Name))
	r.BroadcastMessage(char.Name, fmt.Sprintf("%s waved at %s.", char.Name, targ.Name))
	return nil
}

func (r *Room) Who(char *character.Character) {
	char.SendMessage("/who:")
	for _, c := range r.Characters {
		char.SendMessage(fmt.Sprintf("  %s", c.Name))
	}
}
