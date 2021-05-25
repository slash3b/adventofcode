package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	readAndPrepare("/home/slash3b/Projects/aoc/2015/go/day18/input")
}

func readAndPrepare(s string) {
	res, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(res), "\n")

	var inputLines [][]string
	for _, v := range inputs {
		inputLines = append(inputLines, strings.Split(v, ""))
	}
	// some magic to clear terminal
	fmt.Print("\033[H\033[2J")

	grid := newGrid(inputLines)

	for i := 0; i < 100; i++ {
		grid = processGridLights(grid)
		showLights(grid)
		time.Sleep(time.Millisecond * 100)
	}

	counter := 0
	for _, row := range grid {
		for _, colValue := range row {
			counter += colValue
		}
	}

	showLights(grid)
	fmt.Println("Number of light on: ", counter)
}

func newGrid(input [][]string) [][]int {
	var grid [][]int

	for i := 0; i < 100; i++ {
		var line []int
		for j := 0; j < 100; j++ {
			state := 0
			if input[i][j] == "#" {
				state = 1
			}
			line = append(line, state)
		}
		grid = append(grid, line)
	}
	/*
		1 1 1
		1 1 1
		1 1 1
	*/

	// light up 4 corners >,<
	grid[0][0] = 1
	grid[0][99] = 1
	grid[99][0] = 1
	grid[99][99] = 1

	return grid
}

func newDarkGrid() [][]int {
	var grid [][]int

	for i := 0; i < 100; i++ {
		var line []int
		for j := 0; j < 100; j++ {
			line = append(line, 0)
		}
		grid = append(grid, line)
	}

	return grid
}

// All of the lights update simultaneously;
// they all consider the same current state before moving to the next.
func processGridLights(grid [][]int) [][]int {
	newGrid := newDarkGrid()

	/*
		y
		y
		y
		y
		 x x x x x x
	*/
	for y, row := range grid {
		for x := range row {
			newGrid[y][x] = lookAtNeighbours(grid, x, y)
		}
	}
	// light up 4 corners >,<
	newGrid[0][0] = 1
	newGrid[0][99] = 1
	newGrid[99][0] = 1
	newGrid[99][99] = 1

	return newGrid
}

// lookAtNeighbours returns final result for new light state
func lookAtNeighbours(grid [][]int, y int, x int) int {

	// here we have two major cases
	// - simple one when all neighbours are present
	// - borderline

	/*
		The state a light should have next is based on its current state (on or off) plus the number of neighbors that are on:

		    A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
		    A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.

	*/

	// assuming it is not empty
	getGridValue := func(grid [][]int, y, x int) int {
		gridLength := len(grid) - 1
		rowLength := len(grid[0]) - 1

		if (y < 0 || x < 0) || (y > gridLength || x > rowLength) {
			return 0
		}

		return grid[y][x]
	}

	upperRow := getGridValue(grid, y-1, x-1) + getGridValue(grid, y-1, x) + getGridValue(grid, y-1, x+1)
	sides := getGridValue(grid, y, x-1) + getGridValue(grid, y, x+1)
	lowerRow := getGridValue(grid, y+1, x-1) + getGridValue(grid, y+1, x) + getGridValue(grid, y+1, x+1)

	neighbourLightsCounter := upperRow + sides + lowerRow

	isCurrentGridValueOn := grid[y][x] == 1

	if isCurrentGridValueOn {
		if neighbourLightsCounter == 2 || neighbourLightsCounter == 3 {
			return 1
		} else {
			return 0
		}
	} else if neighbourLightsCounter == 3 {
		return 1
	} else {
		return grid[y][x] // return as is
	}
}

// as in [y][x]int
func isBorderline(grid [][]int, y int, x int) bool {

	if (y == 0 || y == len(grid)-1) || (x == 0 || x == len(grid[0])-1) {
		return true
	}

	return false
}

func showLights(grid [][]int) {
	fmt.Print("\033[H\033[2J")

	for _, row := range grid {
		fmt.Println(row)
	}
}
