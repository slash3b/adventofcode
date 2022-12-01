package main

import (
	"adventofcode/util"
	"flag"
	"fmt"
	"strconv"
)

var test bool
var p int

func init() {
	flag.BoolVar(&test, "test", false, "run with test data. Just pass `-test` to the run command to use test data")
	flag.IntVar(&p, "p", 1, "part to run, 1 by default")
}

func main() {
	flag.Parse()

	inputType := "puzzle"
	if test {
		inputType = "sample"
	}

	lines := util.FileContentsSplitByNewLine(inputType)

	switch p {
	case 1:
		fmt.Println(Part1(lines))
	case 2:
		fmt.Println(Part2(lines))
	default:
		panic("wrong part number given")
	}
}

func Part1(lines []string) int {
	lastElIndex := len(lines) - 1
	if len(lines[lastElIndex]) == 0 {
		lines = lines[:lastElIndex]
	}

	max := 0
	current := 0
	for _, v := range lines {
		if v == "" {
			if current > max {
				max = current
			}
			current = 0
			continue
		}
		intRes, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		current += intRes
	}

	return max
}

func Part2(lines []string) int {
	lastElIndex := len(lines) - 1
	if len(lines[lastElIndex]) == 0 {
		lines = lines[:lastElIndex]
	}

	a, b, c := 0, 0, 0

	current := 0
	for _, v := range lines {
		if v == "" {
			a, b, c = update(a, b, c, current)

			current = 0
			continue
		}
		intRes, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		current += intRes
	}

	fmt.Println(a, b, c)

	return a + b + c
}

// todo: figure how to find smallest and returns other three ints
// order here does not matter
func update(a, b, c, current int) (int, int, int) {

	if a == 0 || (a < b && a < c && a < current) {
		a = current
		return a, b, c
	}

	if b == 0 || (b < c && b < a && b < current) {
		b = current
		return a, b, c
	}

	if c == 0 || (c < b && c < a && c < current) {
		c = current
		return a, b, c
	}

	return a, b, c
}
