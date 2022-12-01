package main

import (
	"adventofcode/util"
	"flag"
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

	util.ScanFile(inputType)

	switch p {
	case 1:
		Part1()
	case 2:
		Part2()
	default:
		panic("wrong part number given")
	}
}

func Part1() {
}

func Part2() {
}
