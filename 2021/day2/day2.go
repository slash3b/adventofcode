package main

import (
	"adventofcode/util"
	"flag"
	"fmt"
	"strconv"
	"strings"
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
	data := util.FileToStringSlice(inputType)
	// fmt.Printf("%#v", data)

	switch partNum {
	case 1:
		Part1(data)
	case 2:
		Part2(data)
	default:
		panic("wrong part number given")
	}
}

func Part1(in []string) {
	depth := 0
	frw := 0

	for _, v := range in {
		res := strings.Split(v, " ")
		k := res[0]
		v, err := strconv.Atoi(res[1])
		if err != nil {
			panic(err)
		}

		switch k {
		case "forward":
			frw += v
		case "down":
			depth += v
		case "up":
			depth -= v
		default:
			panic("unexpected command")
		}
	}

	fmt.Println(depth, frw)

	fmt.Println("result: ", depth*frw)
}

func Part2(in []string) {
	depth := 0
	frw := 0 // horizontal position
	aim := 0

	for _, v := range in {
		res := strings.Split(v, " ")
		k := res[0]
		v, err := strconv.Atoi(res[1])
		if err != nil {
			panic(err)
		}

		switch k {
		case "forward":
			frw += v
			depth += aim * v
		case "down":
			aim += v
		case "up":
			aim -= v
		default:
			panic("unexpected command")
		}
	}

	fmt.Println(depth, frw)

	fmt.Println("result: ", depth*frw)
}
