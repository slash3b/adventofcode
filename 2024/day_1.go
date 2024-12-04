package main

import (
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/slash3b/aoc/util"
)

var (
	input      = flag.String("input", "", "file name")
	res1, res2 = 0, 0
)

func main() {
	flag.Parse()

	lines := util.FileToStrings(*input)

	a := make([]int, 0, len(lines))
	b := make([]int, 0, len(lines))

	m := make(map[int]int, len(lines))

	for _, v := range lines {
		res := strings.Split(v, " ")

		a = append(a, util.MustAtoi(res[0]))

		bvalue := util.MustAtoi(res[3])
		b = append(b, bvalue)
		m[bvalue]++
	}

	slices.Sort(a)
	slices.Sort(b)

	dist := 0
	simscore := 0
	for i := range a {
		dist += util.Diff(a[i], b[i])

		if v, ok := m[a[i]]; ok {
			simscore += a[i] * v
		}
	}

	fmt.Println("total dist", dist)
	fmt.Println("similarity score", simscore)
}
