package main

import (
	"adventofcode/util"
	"flag"
	"fmt"
	"regexp"
	"strconv"
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

func Part1(lines []string) int {
	sum := 0
	reg := regexp.MustCompile(`\d+`)
	for _, v := range lines {
		ints := reg.FindAllString(v, -1)
		fmt.Println("pairs", ints)
		a, err := strconv.Atoi(ints[0])
		if err != nil {
			panic(err)
		}

		b, err := strconv.Atoi(ints[1])
		if err != nil {
			panic(err)
		}

		c, err := strconv.Atoi(ints[2])
		if err != nil {
			panic(err)
		}

		d, err := strconv.Atoi(ints[3])
		if err != nil {
			panic(err)
		}

		if c >= a && d <= b {
			sum++
			continue
		}

		if a >= c && b <= d {
			sum++
		}
	}

	return sum
}

func Part2(lines []string) int {
	sum := 0
	reg := regexp.MustCompile(`\d+`)
	for _, v := range lines {
		ints := reg.FindAllString(v, -1)
		a, err := strconv.Atoi(ints[0])
		if err != nil {
			panic(err)
		}

		b, err := strconv.Atoi(ints[1])
		if err != nil {
			panic(err)
		}

		c, err := strconv.Atoi(ints[2])
		if err != nil {
			panic(err)
		}

		d, err := strconv.Atoi(ints[3])
		if err != nil {
			panic(err)
		}

		if c >= a && c <= b {
			sum++
			continue
		}

		if d >= a && d <= b {
			sum++
			continue
		}

		// check the other side
		// ab, cd
		if a >= c && a <= d {
			sum++
			continue
		}

		if b >= c && b <= d {
			sum++
		}

	}

	return sum
}
