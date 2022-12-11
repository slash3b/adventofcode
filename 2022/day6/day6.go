package main

import (
	"adventofcode/util"
	"flag"
	"fmt"
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

func Part1(lines []string) int {
	for _, v := range lines {
		for i := 3; i < len(v); i++ {

			a := v[i-3] != v[i-2] && v[i-3] != v[i-1] && v[i-3] != v[i]
			b := v[i-2] != v[i-1] && v[i-2] != v[i]
			c := v[i-1] != v[i-0]
			if a && b && c {
				fmt.Println(v, i+1)
				break
			}

		}
	}

	return 0
}

type Detector struct {
	DupN int
	m    map[rune]int
	buf  []rune
}

func NewDetector() *Detector {
	m := make(map[rune]int)

	return &Detector{
		m: m,
	}
}

func (d *Detector) AddAndCheck(r rune) bool {
	// fmt.Println("dup:", d.DupN, "buf state:", string(d.buf))

	d.m[r]++
	if d.m[r] > 1 {
		d.DupN++
	}

	d.buf = append(d.buf, r)

	if len(d.buf) > 14 {

		headPop := d.buf[0]
		d.buf = d.buf[1:]

		if d.m[headPop] > 1 {
			d.DupN--
		}
		d.m[headPop]--

		// res := d.DupN == 0
		// if res {
		// 	fmt.Println("MAP content")
		// 	for k, v := range d.m {
		// 		fmt.Printf("k: %s, v: %d \n", string(k), v)
		// 	}

		// }
		return d.DupN == 0
	}

	return false
}

func Part2(lines []string) int {

	for _, input := range lines {
		fmt.Println("-------------------", input)
		dt := NewDetector()
		for i, v := range input {
			if dt.AddAndCheck(v) {
				fmt.Println("line:", input, "answer:", i+1)
				break
			}
		}
	}

	return 0
}
