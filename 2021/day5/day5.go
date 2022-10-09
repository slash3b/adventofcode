package main

import (
	"adventofcode/util"
	"flag"
	"fmt"
	"regexp"
	"strconv"
)

var partNum int
var inputType string

func init() {
	flag.IntVar(&partNum, "part", 1, "part to run, can be either 1 or 2")
	flag.StringVar(&inputType, "input", "sample", "input type to use, either sample or puzzle. It is 'sample' by default")
}

type Point struct {
	X int
	Y int
}

type Line struct {
	A Point
	B Point
}

func (l *Line) IsOrthogonal() bool {
	if l.A.X == l.B.X {
		return true
	}

	if l.A.Y == l.B.Y {
		return true
	}

	return false
}

func (l *Line) GenVector() [][]int {
	a := l.A
	b := l.B
	vec := [][]int{}

	if a.X == b.X {
		x := a.X
		yMin := min(a.Y, b.Y)
		yMax := max(a.Y, b.Y)

		for i := yMin; i <= yMax; i++ {
			vec = append(vec, []int{x, i})
		}

		return vec
	}

	if a.Y == b.Y {
		y := a.Y
		xMin := min(a.X, b.X)
		xMax := max(a.X, b.X)

		for i := xMin; i <= xMax; i++ {
			vec = append(vec, []int{i, y})
		}

		return vec
	}

	// down below is too.... cumbersome
	// is there a smarter approach?

	xses := []int{}
	yses := []int{}

	if a.X > b.X {
		for ; a.X >= b.X; a.X-- {
			xses = append(xses, a.X)
		}
	} else {
		for ; a.X <= b.X; a.X++ {
			xses = append(xses, a.X)
		}
	}

	if a.Y > b.Y {
		for ; a.Y >= b.Y; a.Y-- {
			yses = append(yses, a.Y)
		}
	} else {
		for ; a.Y <= b.Y; a.Y++ {
			yses = append(yses, a.Y)
		}
	}

	for k, x := range xses {
		vec = append(vec, []int{x, yses[k]})
	}

	return vec
}

func main() {
	flag.Parse()

	lines := []Line{}

	for _, v := range util.FileToStringSlice(inputType) {
		reg := regexp.MustCompile(`\d+`)

		res := reg.FindAllString(v, -1)
		if len(res) != 4 {
			panic(fmt.Sprintf("unexpected nub of matches, expected 4, actual %d", len(res)))
		}

		a := Point{}
		a.X, _ = strconv.Atoi(res[0])
		a.Y, _ = strconv.Atoi(res[1])

		b := Point{}
		b.X, _ = strconv.Atoi(res[2])
		b.Y, _ = strconv.Atoi(res[3])

		l := Line{
			A: a,
			B: b,
		}

		lines = append(lines, l)
	}

	switch partNum {
	case 1:
		Part1(lines)
	case 2:
		Part2(lines)
	default:
		panic("wrong part number given")
	}
}

func Part1(lines []Line) {
	var f = [1000][1000]int{}

	for _, v := range lines {
		// fmt.Println(v)
		if !v.IsOrthogonal() {
			continue
		}
		for _, crs := range v.GenVector() {
			f[crs[0]][crs[1]]++
		}
	}

	res := 0
	for _, row := range f {
		for _, col := range row {
			if col > 1 {
				res++
			}
		}
	}
	fmt.Println(">>>>>>>>>>>>>>>>>", res)
}

func Part2(lines []Line) {
	var f = [1000][1000]int{}

	for _, v := range lines {
		for _, crs := range v.GenVector() {
			f[crs[1]][crs[0]]++
		}
	}

	res := 0
	for _, row := range f {
		for _, col := range row {
			if col > 1 {
				res++
			}
		}
	}

	fmt.Println(">>>>>>>>>>>>>>>>>", res)
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func diff(a, b int) int {
	res := a - b
	if res < 0 {
		res *= -1
	}

	return res
}
