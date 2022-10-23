package main

import (
	"adventofcode/util"
	"flag"
	"fmt"
	"time"
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

	ints := util.FileToIntSlice(inputType)

	start := time.Now()
	switch p {
	case 1:
		Part1(ints)
	case 2:
		Part2(ints)
	default:
		panic("wrong part number given")
	}
	fmt.Println("Runtime:", time.Now().Sub(start))
}

func Part1(input []int) {
	j, k := input[0], input[0]

	for _, v := range input {
		if v < j {
			j = v
		}

		if v > k {
			k = v
		}
	}

	minFuel := 0

	for ; j <= k; j++ {
		fu := 0
		for _, crab := range input {
			fu += util.Diff(crab, j)
		}

		if minFuel == 0 && fu > 0 {
			minFuel = fu
			continue
		}

		if fu < minFuel {
			minFuel = fu
		}
	}

	fmt.Println("RES:", minFuel)
}

func Part2(input []int) {
	j, k := input[0], input[0]

	for _, v := range input {
		if v < j {
			j = v
		}

		if v > k {
			k = v
		}
	}

	minFuel := 0

	for ; j <= k; j++ {
		fu := 0
		for _, crab := range input {
			fu += calcFuel(util.Diff(crab, j))
		}

		if minFuel == 0 && fu > 0 {
			minFuel = fu
			continue
		}

		if fu < minFuel {
			minFuel = fu
		}
	}

	fmt.Println("RES:", minFuel)
}

var fuels = make(map[int]int)

func calcFuel(steps int) int {
	if v, ok := fuels[steps]; ok {
		return v
	}

	res := 0
	for i := 1; i <= steps; i++ {
		res += i
	}

	fuels[steps] = res

	return res
}
