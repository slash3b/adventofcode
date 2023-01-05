package main

import (
	"adventofcode/util"
	"flag"
	"fmt"
	"strconv"
	"strings"
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
	plane := make([][]int, 0)
	for i := 0; i < 1000; i++ {
		plane = append(plane, make([]int, 1000))
	}

	x, y := 499, 499   // middle of a plane
	x2, y2 := 499, 499 // middle of a plane

	tailVisits := 0

	for _, v := range input {
		res := strings.Split(v, " ")
		move := res[0]
		steps, err := strconv.Atoi(res[1])
		if err != nil {
			panic(err)
		}

		switch move {
		case "U":
			for i := 0; i < steps; i++ {
				y--

				// adjust tail
				if diff(y, y2) > 1 || diff(x, x2) > 1 {
					y2 = y + 1
					x2 = x

					if plane[y2][x2] == 0 {
						tailVisits++
					}

					plane[y2][x2]++
				}

				// drawPlane(plane)
			}

		case "D":
			for i := 0; i < steps; i++ {
				y++

				// adjust tail
				if diff(y, y2) > 1 || diff(x, x2) > 1 {
					y2 = y - 1
					x2 = x

					if plane[y2][x2] == 0 {
						tailVisits++
					}

					plane[y2][x2]++
				}

				// drawPlane(plane)
			}

		case "R":
			for i := 0; i < steps; i++ {
				x++

				// adjust tail
				if diff(y, y2) > 1 || diff(x, x2) > 1 {
					y2 = y
					x2 = x - 1

					if plane[y2][x2] == 0 {
						tailVisits++
					}

					plane[y2][x2]++
				}

				// drawPlane(plane)
			}

		case "L":
			for i := 0; i < steps; i++ {
				x--

				// adjust tail
				if diff(y, y2) > 1 || diff(x, x2) > 1 {
					y2 = y
					x2 = x + 1

					if plane[y2][x2] == 0 {
						tailVisits++
					}

					plane[y2][x2]++
				}

				// drawPlane(plane)
			}
		}
	}

	fmt.Println(tailVisits + 1)

	return 0
}

func drawPlane(input [][]int) {
	time.Sleep(time.Millisecond * 5)
	fmt.Print("\033[H\033[2J")

	for _, v := range input {
		for _, v := range v {
			if v == 0 {
				fmt.Print(" ")
				continue
			}

			fmt.Print(v)
		}
		fmt.Println()
	}
}

func Part2(input []string) int {
	return 0
}

func diff(a, b int) int {
	res := a - b

	if res < 0 {
		return res * -1
	}

	return res
}
