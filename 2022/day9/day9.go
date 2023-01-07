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
	x2, y2 := 499, 150 // middle of a plane

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
	time.Sleep(time.Second)
	fmt.Print("\033[H\033[2J")

	for _, v := range input {
		for _, v := range v {
			if v == -1 {
				fmt.Print("H")
				continue
			}

			if v == 0 {
				fmt.Print(".")
				continue
			}

			fmt.Print(v)
		}
		fmt.Println()
	}
}

func drawPlane2(rope map[int]*Coord) {
	time.Sleep(time.Second)
	// 	fmt.Print("\033[H\033[2J")

	plane := make([][]int, 0)
	for i := 0; i < 50; i++ {
		plane = append(plane, make([]int, 50))
	}

	for k, v := range rope {
		if k == 0 {
			k = -1
		}
		plane[v.Y][v.X] = k
	}

	for _, v := range plane {
		for _, v := range v {
			switch v {
			case -1:
				fmt.Print("H")
			case 9:
				fmt.Print("T")
			case 0:
				fmt.Print(".")
			default:
				fmt.Printf("%d", v)
			}
		}
		fmt.Println()
	}
}

func drawCoords(m map[int]*Coord) {
	time.Sleep(time.Second)
	fmt.Print("\033[H\033[2J")

	for i := 0; i < 10; i++ {
		v, _ := m[i]
		fmt.Printf("%d: Y:%d, X:%d \n", i, v.Y, v.X)
	}
}

type Coord struct {
	Y int
	X int
}

func (c *Coord) Move(other *Coord) bool {
	yDiff := other.Y - c.Y
	xDiff := other.X - c.X

	fmt.Printf("parent: %#v\n", other)
	fmt.Printf("child: %#v\n", c)
	fmt.Println(yDiff, xDiff)

	/*
				12345

				HHHHH
		  16	HabcH	6
		  15	HhTdH	7
		  14	HgfeH	8
				HHHHH

				11119
				3210
	*/

	// A corner
	// 1
	if yDiff == -2 && xDiff == -2 {
		c.X--
		c.Y--
		return true
	}
	// 2
	if yDiff == -2 && xDiff == -1 {
		c.X--
		c.Y--
		return true
	}
	// 16
	if yDiff == -1 && xDiff == -2 {
		c.X--
		c.Y--
		return true
	}

	// B spot
	// 3
	if yDiff == -2 && xDiff == 0 {
		c.Y++
		return true
	}

	// C corner
	// 4
	if yDiff == -2 && xDiff == 1 {
		c.Y--
		c.X++
		return true
	}
	// 5
	if yDiff == -2 && xDiff == 2 {
		c.Y--
		c.X++
		return true
	}
	// 6
	if yDiff == -1 && xDiff == 2 {
		c.Y--
		c.X++
		return true
	}

	// D corner
	// 7
	if yDiff == 0 && xDiff == 2 {
		c.X++
		return true
	}

	/*
				12345

				HHHHH
		  16	HabcH	6
		  15	HhTdH	7
		  14	HgfeH	8
				HHHHH

				11119
				3210
	*/
	// 8
	if yDiff == 1 && xDiff == 2 {
		c.Y++
		c.X++
		return true
	}
	// 9
	if yDiff == 2 && xDiff == 2 {
		c.Y++
		c.X++
		return true
	}
	// 10
	if yDiff == 2 && xDiff == 1 {
		c.Y++
		c.X++
		return true
	}

	// F corner
	// 11
	if yDiff == 2 && xDiff == 0 {
		c.Y--
		return true
	}

	// G corner
	// 12
	if yDiff == 2 && xDiff == -1 {
		c.Y++
		c.X--
		return true
	}
	// 13
	if yDiff == 2 && xDiff == -2 {
		c.Y++
		c.X--
		return true
	}
	// 14
	if yDiff == 1 && xDiff == -2 {
		c.Y++
		c.X--
		return true
	}

	// H corner
	// 15
	if yDiff == 0 && xDiff == -2 {
		c.X--
		return true
	}

	return false
}

func Part2(input []string) int {
	y, x := 10, 10 // head

	m := make(map[int]*Coord)
	for i := 0; i < 2; i++ {
		m[i] = &Coord{Y: y, X: x}
	}

	tailVisits := 0

	for _, v := range input {
		res := strings.Split(v, " ")

		move := res[0]
		steps, err := strconv.Atoi(res[1])
		if err != nil {
			panic(err)
		}

		for i := 0; i < steps; i++ {
			switch move {
			case "U":
				m[0].Y--
			case "D":
				m[0].Y++
			case "R":
				m[0].X++
			case "L":
				m[0].X--
			}

			for i := 1; i < 2; i++ {
				prev := m[i-1]
				curr := m[i]
				if !curr.Move(prev) {
				}
			}

			drawPlane2(m)
			// drawCoords(m)
		}
	}

	fmt.Println(tailVisits)

	return 0
}

func diff(a, b int) int {
	res := a - b

	if res < 0 {
		return res * -1
	}

	return res
}
