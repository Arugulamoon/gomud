package world

import (
	"testing"

	"github.com/Arugulamoon/gomud/pkg/character"
	"github.com/Arugulamoon/gomud/pkg/session"
)

func TestContainsCharacter(t *testing.T) {
	char1 := character.NewCharacter()
	char2 := character.NewCharacter()

	emptyRoom := Room{
		Sessions:   make(map[string]*session.Session),
		Characters: make(map[string]*character.Character),
	}

	roomWithChar1 := Room{
		Sessions:   make(map[string]*session.Session),
		Characters: make(map[string]*character.Character),
	}
	roomWithChar1.Sessions["1"] = &session.Session{
		Id:        "1",
		Character: char1,
	}
	roomWithChar1.Characters["1"] = char1

	type given struct {
		room Room
		name string
	}

	tests := []struct {
		given
		want bool
	}{
		{given{emptyRoom, char1.Name}, false},
		{given{roomWithChar1, char1.Name}, true},
		{given{roomWithChar1, char2.Name}, false},
	}
	for _, tt := range tests {
		got := tt.given.room.ContainsCharacter(tt.name)
		if got != tt.want {
			t.Errorf("given: %+v, \"%s\"\ngot: \"%t\"\nwant: \"%t\"",
				tt.room, tt.name, got, tt.want)
		}
	}
}
