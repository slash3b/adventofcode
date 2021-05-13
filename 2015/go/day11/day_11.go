package main

import (
	"fmt"
)

/*
	1. Increase the rightmost letter one step; if it was z, it wraps around to a, and repeat with the next letter to the left until one doesn't wrap around.
	2. must include one increasing straight of at least three letters, like abc, bcd, cde, and so on, up to xyz
	3. may not contain the letters i, o, or l
	4. contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz
*/

func main() {

	// 97 - 122
	//input := []byte("hxbxwxba")
	input := []byte("hxbxxyzz")

	for {
		if input[0] == 122 {
			break
		}

		flip(input)

		if isValid(input) {
			fmt.Println("DONE: ", string(input))
			break
		}

		// check if valid
		// if not -- do the increase in a loop
		// with every iteration check if valid

		fmt.Println(string(input))
	}
}

func flip(input []byte) {

	var copyBytes []byte
	copyBytes = append(copyBytes, input...)

	inputLen := len(input)

	// rotate the very last element
	input[inputLen-1] = rotate(input[inputLen-1])

	// continue with all others
	for i := len(input) - 2; i >= 0; i-- {
		if input[i+1] == 97 && copyBytes[i+1] != 97 {
			input[i] = rotate(input[i])
		}
	}
}

func rotate(b byte) byte {
	// 97 - 122
	if b == 122 {
		return 97
	}

	b++
	return b
}

func isValid(s []byte) bool {
	return isIncreasing(s) && isLegalChars(s) && isPairs(s)
}

func isPairs(s []byte) bool {
	var count int
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			count++
			i++ // jump to the very next non-overlapping
		}
	}

	return count >= 2
}

func isLegalChars(s []byte) bool {
	for _, char := range s {
		// i, o, l
		// 105, 108, 111
		if char == 105 || char == 108 || char == 111 {
			return false
		}
	}

	return true
}

func isIncreasing(s []byte) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1]-1 && s[i] == s[i+2]-2 {
			return true
		}
	}

	return false
}
