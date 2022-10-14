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

	stringFish := strings.Split(strings.Split(util.FileContents(inputType), "\n")[0], ",")
	fish := []int{}
	for _, v := range stringFish {
		if v == "" || v == " " {
			continue
		}
		res, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		fish = append(fish, res)
	}

	switch partNum {
	case 1:
		Part1(fish)
	case 2:
		Part2(fish)
	default:
		panic("wrong part number given")
	}
}

func Part1(fish []int) {
	i := 0

	// turns out this is a naive approach that
	// simply does not scale
	for i < 80 {
		for i := range fish {
			// ---------- CORRECTION -------
			if fish[i] == 0 {
				fish[i] = 6
				fish = append(fish, 8)
				continue
			}
			// ---------- NEXT DAY ---------
			fish[i]--
		}
		i++
	}

	fmt.Println("RESULT", len(fish))
}

func Part2(fish []int) {
	fmt.Println(fish)

	fishTime := map[int]int{}
	for _, v := range fish {
		fishTime[v]++ // one of each as in input
	}

	// so this is a "liz the grey" approach
	// main idea is that instead of using an array
	// we use fish timers/clocks
	// in eli5 terms, we count fish instead of directly producing
	// larger output like in Part1
	for i := 0; i < 256; i++ {
		newFish := map[int]int{}

		for k, v := range fishTime {
			if k == 0 {
				newFish[6] = v
				newFish[8] = v
				continue
			}
			newFish[k-1] += v
		}
		fishTime = newFish
	}

	res := 0
	for _, v := range fishTime {
		res += v
	}

	fmt.Println(res)
}
