package main

import (
	"fmt"
	"testing"
)

func TestIsShortDouble(t *testing.T) {
	strs := map[string]bool{
		// vowels
		"aaa":        false,
		"aab":        false,
		"abcdefgg":   false,
		"aabcdefgaa": true,
		"xyxy":       true,
		"xyaxy":      true,
		"ffcdefdb":   false,
		"ffcdeff":    true,
		"ffcdefff":   true,
		"fffcdeff":   true,
		"aaacdeaaa":  true,
	}

	for str, exp := range strs {
		res := isShortDouble(str)
		if res != exp {
			panic(fmt.Sprintln("expected: ", exp, "but received: ", res, " for  arg: ", str))
		}
	}
}

func TestIsBetweenRepeated(t *testing.T) {

	strs := map[string]bool{
		// vowels
		"aaa":              true,
		"xyx":              true,
		"abcdefeghi":       true,
		"ieodomkazucvgmuy": true,
		"ffcdefff":         true,
		"fffcdeff":         true,
		"aacdeaaa":         true,
		"aab":              false,
		"asdfghjk":         false,
	}

	for str, exp := range strs {
		res := isBetweenRepeated(str)
		if res != exp {
			panic(fmt.Sprintln("expected: ", exp, "but received: ", res, " for  arg: ", str))
		}
	}
}
