package session

import (
	"fmt"
)

// English Sentence Structure
// Subject + Verb + Object
// Incorporate Observer View
// See Tests

// Subject / Writer
func ReflexiveSubject(subject, observer string) string {
	if observer == subject {
		return "You"
	}
	return subject
}

// Verb
func ReflexiveVerb(subject, verb, observer string) string {
	if observer == subject {
		return verb
	}
	if verb == "say" {
		return "said"
	}
	if verb[len(verb)-1:] == "e" {
		return fmt.Sprintf("%sd", verb)
	}
	return fmt.Sprintf("%sed", verb)
}

// Object / Target / Reader
func ReflexiveObject(subject, object, observer string) string {
	if subject == object {
		if observer == subject {
			return "yourself"
		}
		return "themselves"
	} else if object == observer {
		return "you"
	} else {
		return object
	}
}
