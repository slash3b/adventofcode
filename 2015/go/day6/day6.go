package main

import (
	"fmt"
	"image"
	"image/draw"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// dot[0] is x
// dot[1] is y
type dot [2]int

type command struct {
	direction string
	A         dot
	B         dot
}

func main() {
	grid := newGrid()
	cmds := grabCommands("back_input")

	img := image.NewRGBA(image.Rect(0, 0, 999, 999))
	draw.Draw(img, img.Bounds(), &image.Uniform{image.Black}, image.Point{}, draw.Src)

	//part1(cmds, grid)
	part2(cmds, grid)
}

func part1(cmds []command, grid [][]int) {
	var aCounter int
	for _, v := range cmds {
		switch v.direction {
		case "turn off":
			for rowIndex := v.A[0]; rowIndex <= v.B[0]; rowIndex++ {
				for colIndex := v.A[1]; colIndex <= v.B[1]; colIndex++ {
					if grid[rowIndex][colIndex] == 1 {
						grid[rowIndex][colIndex] = 0
						aCounter--
					}
				}
			}
		case "turn on":
			for rowIndex := v.A[0]; rowIndex <= v.B[0]; rowIndex++ {
				for colIndex := v.A[1]; colIndex <= v.B[1]; colIndex++ {
					if grid[rowIndex][colIndex] == 0 {
						grid[rowIndex][colIndex] = 1
						aCounter++
					}
				}
			}
		case "toggle":
			for rowIndex := v.A[0]; rowIndex <= v.B[0]; rowIndex++ {
				for colIndex := v.A[1]; colIndex <= v.B[1]; colIndex++ {
					if grid[rowIndex][colIndex] == 0 {
						grid[rowIndex][colIndex] = 1
						aCounter++
						continue
					}
					if grid[rowIndex][colIndex] == 1 {
						grid[rowIndex][colIndex] = 0
						aCounter--
					}
				}
			}
		}
	}

	// uncomment to get a picture : )
	/*
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == 1 {
					draw.Draw(img, image.Rect(i, j, i+1, j+1), &image.Uniform{image.White}, image.Point{}, draw.Src)
				}
			}
		}
		file, _ := os.Create("lamps.png")

		png.Encode(file, img)
	*/

	fmt.Println("amount of lamps.png turned on: ", aCounter)
}

func part2(cmds []command, grid [][]int) {
	var brightnessCounter int
	for _, v := range cmds {
		switch v.direction {

		case "turn off":
			for rowIndex := v.A[0]; rowIndex <= v.B[0]; rowIndex++ {
				for colIndex := v.A[1]; colIndex <= v.B[1]; colIndex++ {
					if grid[rowIndex][colIndex] > 0 {
						grid[rowIndex][colIndex]--
						brightnessCounter--
					}
				}
			}
		case "turn on":
			for rowIndex := v.A[0]; rowIndex <= v.B[0]; rowIndex++ {
				for colIndex := v.A[1]; colIndex <= v.B[1]; colIndex++ {
					grid[rowIndex][colIndex]++
					brightnessCounter++
				}
			}
		case "toggle":
			for rowIndex := v.A[0]; rowIndex <= v.B[0]; rowIndex++ {
				for colIndex := v.A[1]; colIndex <= v.B[1]; colIndex++ {
					grid[rowIndex][colIndex] += 2
					brightnessCounter += 2
				}
			}
		}
	}

	fmt.Println("total brightness: ", brightnessCounter)
}

func grabCommands(file string) []command {
	res, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	commandRe := regexp.MustCompile(`[a-z ]*`)
	coordinatesRe := regexp.MustCompile(`(\d{1,3},\d{1,3})`)

	var result []command
	commands := strings.Split(string(res), "\n")
	for _, cmdStr := range commands {
		coordinates := coordinatesRe.FindAllString(cmdStr, 2)
		res := strings.Split(coordinates[0], ",")
		x1, _ := strconv.Atoi(res[0])
		y1, _ := strconv.Atoi(res[1])
		res = strings.Split(coordinates[1], ",")
		x2, _ := strconv.Atoi(res[0])
		y2, _ := strconv.Atoi(res[1])

		name := strings.TrimRight(commandRe.FindString(cmdStr), " ")

		result = append(result, command{
			direction: name,
			A:         dot{x1, y1},
			B:         dot{x2, y2},
		})
	}

	return result
}

func newGrid() [][]int {
	var grid [][]int

	for i := 0; i < 1000; i++ {
		var line []int
		for i := 0; i < 1000; i++ {
			line = append(line, 0)
		}
		grid = append(grid, line)
	}

	return grid
}
