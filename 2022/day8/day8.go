package main

import (
	"adventofcode/util"
	"flag"
	"fmt"
	"strconv"
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

	input := make([][]int, 0)
	for _, v := range lines {
		var row []int
		for k, _ := range v {
			val, err := strconv.Atoi(string(v[k]))
			if err != nil {
				panic(err)
			}
			row = append(row, val)
		}
		input = append(input, row)
	}

	switch p {
	case 1:
		fmt.Println(Part1(input))
	case 2:
		fmt.Println(Part2(input))
	default:
		panic("wrong part number given")
	}
}

func Part1(input [][]int) int {
	init := len(input)*2 + len(input[0])*2 - 4

	colLen := len(input)
	rowLen := len(input[0])

	var other int
	for y := 1; y < colLen-1; y++ {
		for x := 1; x < rowLen-1; x++ {
			if isVisible(y, x, input) {
				other++
			}
		}
	}
	fmt.Println("INIT", init)
	fmt.Println("OTHER", other)

	return init + other
}

func isVisible(y, x int, input [][]int) bool {
	tree := input[y][x]
	rowLen := len(input[0])

	// right
	right := true
	for i := x + 1; i < rowLen; i++ {
		if input[y][i] >= tree {
			right = false
			break
		}
	}

	// left
	left := true
	for i := x - 1; i >= 0; i-- {
		if input[y][i] >= tree {
			left = false
			break
		}
	}

	// up
	up := true
	for i := y - 1; i >= 0; i-- {
		if input[i][x] >= tree {
			up = false
			break
		}
	}

	// down
	colLen := len(input)
	down := true
	for i := y + 1; i < colLen; i++ {
		if input[i][x] >= tree {
			down = false
			break
		}
	}

	return right || left || up || down
}

func Part2(input [][]int) int {

	var res int

	colLen := len(input)
	rowLen := len(input[0])

	for y := 1; y < colLen-1; y++ {
		for x := 1; x < rowLen-1; x++ {
			if f := score(y, x, input); f > res {
				res = f
			}
		}
	}

	return res
}

func score(y, x int, input [][]int) int {
	tree := input[y][x]
	rowLen := len(input[0])

	// right
	right := 0
	for i := x + 1; i < rowLen; i++ {
		right++
		if input[y][i] >= tree {
			break
		}
	}

	// left
	left := 0
	for i := x - 1; i >= 0; i-- {
		left++
		if input[y][i] >= tree {
			break
		}
	}

	// up
	up := 0
	for i := y - 1; i >= 0; i-- {
		up++
		if input[i][x] >= tree {
			break
		}
	}

	// down
	colLen := len(input)
	down := 0
	for i := y + 1; i < colLen; i++ {
		down++
		if input[i][x] >= tree {
			break
		}
	}

	fmt.Printf("tree %d \n right %d left %d up %d down %d \n", tree, right, left, up, down)

	return right * left * up * down
}
