package world

import "testing"

func TestContainsCharacter(t *testing.T) {
	char1 := NewCharacter()
	char2 := NewCharacter()

	emptyRoom := Room{
		Sessions: make(map[string]*Session),
	}

	roomWithChar1 := Room{
		Sessions: make(map[string]*Session),
	}
	roomWithChar1.Sessions["1"] = &Session{
		Id:        "1",
		Character: char1,
	}

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
