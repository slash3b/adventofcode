package main

import (
	"adventofcode/util"
	"flag"
	"fmt"
	"strconv"
	"strings"
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

func Part1(input []string) int {
	tick := 0
	X := 1
	cycles := []int{20, 40}
	currCycle := cycles[0]
	totFreq := 0
	for _, v := range input {

		tick++

		if tick%currCycle == 0 {
			freq := tick * X
			totFreq += freq
			fmt.Printf("Tick #%d: %d, freq:%d\n", tick, X, freq)
			if len(cycles) > 1 {
				cycles = cycles[1:]
			}

			currCycle += cycles[0]
		} else {
			fmt.Printf("Tick #%d: %d\n", tick, X)
		}

		if v == "noop" {
			continue
		}

		tick++

		res := strings.Split(v, " ")
		v, err := strconv.Atoi(res[1])
		if err != nil {
			panic(err)
		}

		if tick%currCycle == 0 {
			freq := tick * X
			totFreq += freq
			fmt.Printf("Tick #%d: %d, freq:%d\n", tick, X, freq)
			if len(cycles) > 1 {
				cycles = cycles[1:]
			}

			currCycle += cycles[0]
		} else {
			fmt.Printf("Tick #%d: %d\n", tick, X)
		}

		X += v
	}

	fmt.Println("ticked:", tick, "times")
	fmt.Println("final value:", X)
	fmt.Println("total Freq:", totFreq)

	return 0
}

func Part2(input []string) int {
	return 0
}
