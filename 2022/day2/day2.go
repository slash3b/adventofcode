package main

import (
	"adventofcode/util"
	"flag"
	"fmt"
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

	lines := util.FileToStringSlice(inputType)

	switch p {
	case 1:
		fmt.Println(Part1(lines))
	case 2:
		fmt.Println(Part2(lines))
	default:
		panic("wrong part number given")
	}
}

// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock
// A for Rock, B for Paper, and C for Scissors
// X for Rock, Y for Paper, and Z for Scissors
var m = map[byte]int{
	'A': 1,
	'B': 2,
	'C': 3,
	//
	'X': 1,
	'Y': 2,
	'Z': 3,
}

var (
	loss = 0
	draw = 3
	win  = 6
)

func result(input string) int {
	p1 := m[input[0]]
	p2 := m[input[2]]

	if p1 == p2 {
		fmt.Printf("DRAW %s %s: %d %d: res %d \n", string(input[0]), string(input[2]), p1, p2, draw+p2)
		return draw + p2
	}

	if p1 > p2 {
		fmt.Printf("LOSS: %s %s: %d %d: res %d \n", string(input[0]), string(input[2]), p1, p2, loss+p2)
		return loss + p2
	}

	fmt.Printf("WIN: %s %s: %d %d: res %d \n", string(input[0]), string(input[2]), p1, p2, win+p2)
	return win + p2
}

func Part1(lines []string) int {
	fmt.Printf("%v\n\n", m)
	res := 0
	for _, v := range lines {
		res += result(v)
	}

	return res
}

func Part2(lines []string) int {
	return 0
}
