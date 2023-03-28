package world

import (
	"fmt"
	"strings"

	"github.com/Arugulamoon/gomud/pkg/character"
	"github.com/Arugulamoon/gomud/pkg/item"
	"github.com/Arugulamoon/gomud/pkg/room"
)

type World struct {
	Id         string
	Rooms      map[string]*room.Room
	Characters map[string]*character.Character
}

func New(id string) *World {
	return &World{
		Id:         id,
		Rooms:      make(map[string]*room.Room),
		Characters: make(map[string]*character.Character),
	}
}

func (w *World) Load() {
	bedroom := room.New("Bedroom", "You have entered your bedroom. There is a door leading out! (type \"/goto Hallway\" to leave the bedroom)")
	hallway := room.New("Hallway", "You have entered a hallway with doors at either end. (type \"/goto LivingRoom\" to enter the living room or \"/goto Bedroom\" to enter the bedroom)")
	livingRoom := room.New("LivingRoom", "You have entered the living room. (type \"/goto Hallway\" to enter the hallway)")
	bedroom.Paths[hallway.Id] = hallway
	hallway.Paths[bedroom.Id] = bedroom
	livingRoom.Paths[hallway.Id] = hallway
	hallway.Paths[livingRoom.Id] = livingRoom
	book := item.New("Book")
	bedroom.Items[book.Id] = book

	w.Rooms[bedroom.Id] = bedroom
	w.Rooms[hallway.Id] = hallway
	w.Rooms[livingRoom.Id] = livingRoom
}

// Characters
func (w *World) GetCharacters() map[string]*character.Character {
	return w.Characters
}

func (w *World) FindCharacter(name string) *character.Character {
	for _, char := range w.Characters {
		if char.Name == name {
			return char
		}
	}
	return nil
}

func (w *World) addCharacter(char *character.Character) {
	w.Characters[char.Id] = char
}

func (w *World) removeCharacter(char *character.Character) {
	delete(w.Characters, char.Id)
}

// Messaging
func (w *World) BroadcastMessage(speaker, msg string) {
	for _, char := range w.Characters {
		if char.Name != speaker {
			char.SendMessage(msg)
		}
	}
}

// Events
func (w *World) join(char *character.Character) {
	w.addCharacter(char)
	w.BroadcastMessage(char.Name, fmt.Sprintf("%s entered the world.", char.Name))
	char.WorldId = w.Id
}

func (w *World) leave(char *character.Character) {
	w.removeCharacter(char)
	w.BroadcastMessage(char.Name, fmt.Sprintf("%s left the world.", char.Name))
}

// Handlers
func (w *World) HandleCharacterJoined(char *character.Character) {
	w.join(char)
	w.motd(char)
}

func (w *World) HandleCharacterLeft(char *character.Character) {
	w.leave(char)
}

// Available User Commands
const GOTO = "goto"
const MOTD = "motd"
const SHOUT = "shout"
const TELL = "tell"
const WHO = "who"

func (w *World) goTo(char *character.Character, targRoomId string) error {
	currRoom := w.Rooms[char.RoomId]
	targRoom := currRoom.Paths[targRoomId]
	if targRoom == nil {
		char.SendMessage("There is no one around with that name...")
		return fmt.Errorf("target room not found")
	}

	currRoom.Leave(char)
	targRoom.Join(char)
	targRoom.Look(char)
	return nil
}

func (w *World) motd(char *character.Character) {
	char.SendMessage(fmt.Sprintf("Welcome %s!", char.Name))
}

func (w *World) shout(char *character.Character, msg string) error {
	if msg == "" {
		char.SendMessage("Cannot send empty message...")
		return fmt.Errorf("message empty")
	}

	char.SendMessage(fmt.Sprintf("You shout, \"%s\"", msg))
	w.BroadcastMessage(char.Name, fmt.Sprintf("%s shouted, \"%s\"", char.Name, msg))
	return nil
}

func (w *World) tell(char *character.Character, args string) error {
	if args == "" {
		char.SendMessage("Cannot send empty target and message...")
		return fmt.Errorf("target character not found and message empty")
	}

	targetName, msg, _ := strings.Cut(args, " ")
	if msg == "" {
		char.SendMessage("Cannot send empty message...")
		return fmt.Errorf("message empty")
	}
	targ := w.FindCharacter(targetName)
	if targ == nil {
		char.SendMessage("There is no one around with that name...")
		return fmt.Errorf("target character not found")
	}

	char.SendMessage(fmt.Sprintf("You tell %s, \"%s\"", targ.Name, msg))
	targ.SendMessage(fmt.Sprintf("%s tells you, \"%s\"", char.Name, msg))
	return nil
}

func (w *World) who(char *character.Character) {
	char.SendMessage("/who all:")
	for _, c := range w.Characters {
		char.SendMessage(fmt.Sprintf("  %s", c.Name))
	}
}
