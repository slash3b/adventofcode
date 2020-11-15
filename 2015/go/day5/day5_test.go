package main

import (
	"testing"
)

/*
A nice string is one with all of the following properties:

    It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
    It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
    It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.

For example:

    ugknbfddgicrmopn is nice because it has at least three vowels (u...i...o...), a double letter (...dd...), and none of the disallowed substrings.
    aaa is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.
    jchzalrnumimnmhp is naughty because it has no double letter.
    haegwjzuvuyypxyu is naughty because it contains the string xy.
    dvszwmarrgswjxmb is naughty because it contains only one vowel.

*/

func TestWords(t *testing.T) {
	t.SkipNow()

	strs := map[string]bool{
		// vowels
		"aaa":             true,
		"aa":              true,
		"aeiouaeiouaeiou": false,
		"xazegov":         true,
		"bbb":             false,
		"b":               false,

		// those do have forbidden
		"haegwjzuvuyypxyu": false,
		"haegwjzuvuyypxy":  false,
		"xyegwjzuvuyyp":    false,
		"xyaaadd":          false,
	}

	for str, exp := range strs {
		if isNice(str) != exp {
			panic(str)
		}
	}
}
