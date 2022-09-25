package main

import (
	"adventofcode/util"
	"flag"
)

var partNum int
var inputType string

func init() {
	flag.IntVar(&partNum, "part", 1, "part to run, can be either 1 or 2")
	flag.StringVar(&inputType, "input", "sample", "input type to use, either sample or puzzle. It is 'sample' by default")
}

func main() {
	flag.Parse()

	// here do get your input
	data := util.FileToIntSlice(inputType)

	switch partNum {
	case 1:
		Part1(data)
	case 2:
		Part2(data)
	default:
		panic("wrong part number given")
	}
}

func Part1(in []int) int {
	return 0
}

func Part2(in []int) int {
	panic("todo")
}
