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
	fmt.Println(len(strs))

	var counter int
	for x := range strs {
		if isNice(strs[x]) {
			counter++
		}
	}

	fmt.Println("Total of nice strings: ", counter)
}

func isNice(s string) bool {
	return isVowels(s) && isTwice(s) && isNotException(s)
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
		fmt.Println("part: ", part)
		if part == "ab" || part == "cd" || part == "pq" || part == "xy" {
			return false
		}
	}

	return true
}

// It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
func isTwice(s string) bool {
	for i := 0; i < len(s)-1; i += 1 {
		//fmt.Println("aa,dd,xx comparison:", s[i], s[i+1])
		if s[i] == s[i+1] {
			return true
		}
	}

	return false
}
