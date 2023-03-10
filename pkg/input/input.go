package input

import (
	"fmt"
)

func ProcessInput(subject, verb, args, observer string, hasArgs bool) string {
	// TODO: Do something based on cmd
	// ie saying something aloud in a certain room will make something happen
	// ie open door will switch door state to open / close door
	if hasArgs {
		if verb == "say" {
			return Say(subject, args, observer)
		}
		return fmt.Sprintf("%s %s at %s.",
			ReflexiveSubject(subject, observer),
			ReflexiveVerb(subject, verb, observer),
			ReflexiveObject(subject, args, observer))
	}
	if verb == "say" {
		if subject == observer {
			return "You open your mouth to speak, but nothing comes out."
		}
		return fmt.Sprintf("%s opened their mouth to speak, but nothing came out.",
			ReflexiveSubject(subject, observer))
	}
	return fmt.Sprintf("%s %s.",
		ReflexiveSubject(subject, observer),
		ReflexiveVerb(subject, verb, observer))
}

func Say(subject, args, observer string) string {
	return fmt.Sprintf("%s %s, \"%s\"",
		ReflexiveSubject(subject, observer),
		ReflexiveVerb(subject, "say", observer),
		ReflexiveObject(subject, args, observer))
}
