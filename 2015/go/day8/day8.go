package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	readAndPrepare("/home/slash3b/Projects/aoc/2015/go/day8/input")
}

func readAndPrepare(s string) {
	res, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(res), "\n")

	totalChrs := 0
	totalChrsInMemory := 0
	totalEscaped := 0

	for _, inputStr := range inputs {
		totalChrs += len(inputStr)
		totalChrsInMemory += countInMemory(inputStr)
		totalEscaped += countEscaped(inputStr)
	}

	fmt.Println(totalChrs)
	fmt.Println(totalChrsInMemory)
	fmt.Println(totalEscaped)
	fmt.Println(totalChrs - totalChrsInMemory)
	fmt.Println(totalEscaped - totalChrs)
}

func countEscaped(s string) int {
	fmt.Println("String escaped: ", strconv.Quote(s))

	return len(strconv.Quote(s))
}

// find in memory length of a string
func countInMemory(s string) int {
	fmt.Println("Incoming: ", s, " ", len(s))
	// get rid of double quotes on both side
	s = s[1 : len(s)-1]

	hexRegexp := regexp.MustCompile(`(\\x[0-9a-f]{2}|\\"|[\\]{2})`)

	s = hexRegexp.ReplaceAllString(s, ".")

	fmt.Println("Outgoing escaped: ", s, " ", len(s))

	return len(s)
}
