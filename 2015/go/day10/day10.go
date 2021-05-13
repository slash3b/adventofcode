package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var cached map[string]string

func main() {
	input := "1321131112"

	cached = make(map[string]string)

	for i := 0; i < 50; i++ {
		input = nextString(input)
	}

	fmt.Println("Result: ", len(input))
}

func nextString(s string) string {
	parts := regexp.MustCompile(`1+|2+|3+`).FindAllString(s, -1)

	var result = []string{}
	for _, part := range parts {
		result = append(result, convert(part))
	}

	return strings.Join(result, "")
}

func convert(s string) string {
	if value, exists := cached[s]; exists {
		return value
	}

	strLen := len(s)
	if strLen == 1 {
		return "1" + s
	}

	result := strconv.Itoa(len(s)) + string(s[0])

	cached[s] = result

	return result
}
