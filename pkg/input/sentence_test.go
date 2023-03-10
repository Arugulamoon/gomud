package input

import (
	"testing"
)

func TestReflexiveSubject(t *testing.T) {
	type given struct {
		subject  string
		observer string
	}
	tests := []struct {
		given
		want string
	}{
		{given{"Nicholas", "Nicholas"}, "You"},
		{given{"Nicholas", "Deirdre"}, "Nicholas"},
		{given{"Deirdre", "Deirdre"}, "You"},
		{given{"Deirdre", "Nicholas"}, "Deirdre"},
	}
	for _, tt := range tests {
		got := ReflexiveSubject(tt.subject, tt.observer)
		if got != tt.want {
			t.Errorf("given: \"%s\", \"%s\"\ngot: \"%s\"\nwant: \"%s\"",
				tt.subject, tt.observer, got, tt.want)
		}
	}
}

func TestReflexiveVerb(t *testing.T) {
	type given struct {
		subject  string
		verb     string
		observer string
	}
	tests := []struct {
		given
		want string
	}{
		{given{"Nicholas", "say", "Nicholas"}, "say"},       // You say
		{given{"Nicholas", "say", "Deirdre"}, "said"},       // Nicholas said
		{given{"Nicholas", "wave", "Nicholas"}, "wave"},     // You wave
		{given{"Nicholas", "wave", "Deirdre"}, "waved"},     // Nicholas waved
		{given{"Nicholas", "bow", "Nicholas"}, "bow"},       // You bow
		{given{"Nicholas", "bow", "Deirdre"}, "bowed"},      // Nicholas bowed
		{given{"Nicholas", "salute", "Nicholas"}, "salute"}, // You salute
		{given{"Nicholas", "salute", "Deirdre"}, "saluted"}, // Nicholas saluted
		// TODO: other verbs that aren't "ed"
	}
	for _, tt := range tests {
		got := ReflexiveVerb(tt.subject, tt.verb, tt.observer)
		if got != tt.want {
			t.Errorf("given: \"%s\", \"%s\", \"%s\"\ngot: \"%s\"\nwant: \"%s\"",
				tt.subject, tt.verb, tt.observer, got, tt.want)
		}
	}
}

func TestReflexiveObject(t *testing.T) {
	type given struct {
		subject  string
		object   string
		observer string
	}
	tests := []struct {
		given
		want string
	}{
		{given{"Nicholas", "Nicholas", "Nicholas"}, "yourself"},
		{given{"Nicholas", "Nicholas", "Deirdre"}, "themselves"},
		{given{"Nicholas", "Deirdre", "Deirdre"}, "you"},
		{given{"Nicholas", "Dylan", "Deirdre"}, "Dylan"},
	}
	for _, tt := range tests {
		got := ReflexiveObject(tt.subject, tt.object, tt.observer)
		if got != tt.want {
			t.Errorf("given: \"%s\", \"%s\", \"%s\"\ngot: \"%s\"\nwant: \"%s\"",
				tt.subject, tt.object, tt.observer, got, tt.want)
		}
	}
}
