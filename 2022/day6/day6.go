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

func Part1(lines []string) int {
	for _, v := range lines {
		for i := 3; i < len(v); i++ {

			a := v[i-3] != v[i-2] && v[i-3] != v[i-1] && v[i-3] != v[i]
			b := v[i-2] != v[i-1] && v[i-2] != v[i]
			c := v[i-1] != v[i-0]
			if a && b && c {
				fmt.Println(v, i+1)
				break
			}

		}
	}

	return 0
}

func Part2(lines []string) int {
	return 0
}
