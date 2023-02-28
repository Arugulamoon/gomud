package session

import "testing"

func TestContainsCharacter(t *testing.T) {
	emptyRoom := Room{
		Sessions: make(map[string]*Session),
	}

	roomWithDylan := Room{
		Sessions: make(map[string]*Session),
	}
	roomWithDylan.Sessions["1"] = &Session{
		Id: "1",
		User: &User{
			&Character{
				Name: "Dylan",
				Room: &roomWithDylan,
			},
		},
	}

	type given struct {
		room Room
		name string
	}

	tests := []struct {
		given
		want bool
	}{
		{given{emptyRoom, "Dylan"}, false},
		{given{roomWithDylan, "Deirdre"}, false},
		{given{roomWithDylan, "Dylan"}, true},
	}
	for _, tt := range tests {
		got := tt.given.room.ContainsCharacter(tt.name)
		if got != tt.want {
			t.Errorf("given: %+v, \"%s\"\ngot: \"%t\"\nwant: \"%t\"",
				tt.room, tt.name, got, tt.want)
		}
	}
}
