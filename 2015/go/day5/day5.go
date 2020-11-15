package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	lines, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	strs := strings.Split(string(lines), "\n")

	var counter int
	var sCounter int
	for x := range strs {
		s := strs[x]
		if isVowels(s) && isTwice(s) && isNotException(s) {
			counter++
		}
		if isShortDouble(s) && isBetweenRepeated(s) {
			sCounter++
		}
	}

	fmt.Println("Total of nice strings: ", counter)
	fmt.Println("Total of a very nice strings: ", sCounter)
}

// It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
func isVowels(s string) bool {
	vowelsCnt := 0
	for x := range s {
		if string(s[x]) == "a" ||
			string(s[x]) == "e" ||
			string(s[x]) == "i" ||
			string(s[x]) == "o" ||
			string(s[x]) == "u" {
			vowelsCnt++
		}
	}

	if vowelsCnt < 3 {
		return false
	}

	return true
}

// It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.
func isNotException(s string) bool {
	for i := 0; i < len(s)-1; i += 1 {

		part := s[i : i+2]
		if part == "ab" || part == "cd" || part == "pq" || part == "xy" {
			return false
		}
	}

	return true
}

// It contains at least one letter that appears twice in a
// row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
func isTwice(s string) bool {
	for i := 0; i < len(s)-1; i += 1 {
		//fmt.Println("aa,dd,xx comparison:", s[i], s[i+1])
		if s[i] == s[i+1] {
			return true
		}
	}

	return false
}

/*
It contains a pair of any two letters that appears at least twice in
the string without overlapping, like xyxy (xy) or aabcdefgaa (aa),
but not like aaa (aa, but it overlaps).
*/
func isShortDouble(s string) bool {
	if len(s) < 4 {
		return false
	}

	for i := 0; i < len(s)-3; i++ {
		one := s[i : i+2]

		for j := i + 2; j < len(s)-1; j++ {
			two := s[j : j+2]
			if one == two {
				return true
			}
		}
	}

	return false
}

/*
It contains at least one letter which repeats with exactly one letter between them,
like xyx, abcdefeghi (efe), or even aaa.
*/
func isBetweenRepeated(s string) bool {

	for i := 1; i < len(s)-1; i++ {
		before := i - 1
		after := i + 1

		if s[before] == s[after] {
			return true
		}
	}

	return false
}
