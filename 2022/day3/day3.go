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

var Red = "\033[31m"
var Reset = "\033[0m"

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

func charToVal(chr rune) int {
	if chr > 64 && chr < 91 {
		return int(chr - 64 + 26)
	}

	return int(chr - 96)
}

func Part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		a, b := line[:len(line)/2], line[len(line)/2:]
		m := make(map[rune]struct{})
		for _, v := range a {
			if _, ok := m[v]; !ok {
				m[v] = struct{}{}
			}
		}

		res := make(map[rune]struct{})
		for _, v := range b {
			_, ok := m[v]
			_, vok := res[v]
			if ok && !vok {
				res[v] = struct{}{}
			}
		}

		for k, v := range line {
			if k == len(line)/2 {
				fmt.Print(" : ")
			}
			if _, ok := res[v]; ok {
				fmt.Print(Red + string(v) + Reset)
				continue
			}
			fmt.Print(string(v))
		}

		foo := 0
		for k, _ := range res {
			fmt.Print(string(k))
			val := charToVal(k)
			fmt.Printf("%d ", val)
			foo += val
		}
		fmt.Printf(" : total line %d", foo)
		sum += foo

		fmt.Println("\n-------------------------------------------")
	}

	return sum
}

func Part2(lines []string) int {
	return 0
}
