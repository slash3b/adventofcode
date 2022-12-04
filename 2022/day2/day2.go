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

// A for Rock, B for Paper, and C for Scissors
// X for Rock, Y for Paper, and Z for Scissors

var (
	loss = 0
	draw = 3
	win  = 6
)

// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock
var m1 = map[string]int{
	"A X": 1 + draw, // Rock == Rock
	"A Y": 2 + win,  // Rock < Paper
	"A Z": 3 + loss, // Rock > Scissors
	//
	"B X": 1 + loss, // Paper > Rock
	"B Y": 2 + draw, // Paper == Paper
	"B Z": 3 + win,  // Paper < Scissors
	//
	"C X": 1 + win,  // Scissors < Rock
	"C Y": 2 + loss, // Scissors > Paper
	"C Z": 3 + draw, // Scissors == Scissors
}

// 'A' Rock
// 'B' Paper
// 'C' Scissors
// 'X' Rock
// 'Y' Paper
// 'Z' Scissors

// X means you need to loss, Y means you need to end the round in a draw, and Z means you need to win
// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock
// 1 for Rock, 2 for Paper, and 3 for Scissors
var m2 = map[string]int{
	"A X": 3 + loss, // Rock > Scissors
	"A Y": 1 + draw, // Rock == Rock
	"A Z": 2 + win,  // Rock < Paper
	//
	"B X": 1 + loss, // Paper > Rock
	"B Y": 2 + draw, // Paper == Paper
	"B Z": 3 + win,  // Paper < Scissors
	//
	"C X": 2 + loss, // Scissors > Paper
	"C Y": 3 + draw, // Scissors == Scissors
	"C Z": 1 + win,  // Scissors < Rock
}

func Part1(lines []string) int {
	res := 0
	for _, v := range lines {
		mv, ok := m1[v]
		if !ok {
			panic(fmt.Sprintf("unknown key %s", v))
		}

		res += mv
	}

	return res
}

func Part2(lines []string) int {
	res := 0
	for _, v := range lines {
		mv, ok := m2[v]
		if !ok {
			panic(fmt.Sprintf("unknown key %s", v))
		}
		res += mv
	}

	return res
}
