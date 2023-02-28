package session

import (
	"testing"
)

func TestProcessInput(t *testing.T) {
	type given struct {
		subject  string
		verb     string
		args     string
		observer string
		hasArgs  bool
	}
	helloTests := []struct {
		given
		want string
	}{
		{given{"Nicholas", "say", "Hello World!", "Nicholas", true}, "You say, \"Hello World!\""},
		{given{"Nicholas", "say", "Hello World!", "Deirdre", true}, "Nicholas said, \"Hello World!\""},
		{given{"Nicholas", "say", "Hello", "Nicholas", true}, "You say, \"Hello\""},
		{given{"Nicholas", "say", "", "Nicholas", false}, "You open your mouth to speak, but nothing comes out."},
		{given{"Nicholas", "say", "", "Deirdre", false}, "Nicholas opened their mouth to speak, but nothing came out."},

		{given{"Nicholas", "wave", "", "Nicholas", false}, "You wave."},
		{given{"Nicholas", "wave", "", "Deirdre", false}, "Nicholas waved."},

		{given{"Nicholas", "wave", "Nicholas", "Nicholas", true}, "You wave at yourself."},
		{given{"Nicholas", "wave", "Dylan", "Nicholas", true}, "You wave at Dylan."},
		{given{"Nicholas", "wave", "Nicholas", "Deirdre", true}, "Nicholas waved at themselves."},

		{given{"Nicholas", "wave", "Deirdre", "Deirdre", true}, "Nicholas waved at you."},

		{given{"Nicholas", "wave", "Dylan", "Deirdre", true}, "Nicholas waved at Dylan."},
	}
	for _, tt := range helloTests {
		got := ProcessInput(tt.subject, tt.verb, tt.args, tt.observer, tt.hasArgs)
		if got != tt.want {
			t.Errorf("given: %+v\ngot: \"%s\"\nwant: \"%s\"", tt.given, got, tt.want)
		}
	}
}
