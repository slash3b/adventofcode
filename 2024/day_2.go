package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/slash3b/aoc/util"
)

var (
	input = flag.String("input", "", "file name")
)

func main() {
	flag.Parse()

	lines := util.FileToStrings(*input)

	reports := make([][]int, len(lines))

	for i, ln := range lines {
		for _, lvl := range strings.Split(ln, " ") {
			reports[i] = append(reports[i], util.MustAtoi(lvl))
		}
	}

	totalSafe := 0
	totalSafe2 := 0

	for _, report := range reports {
		if isSafe(report) {
			totalSafe++
			totalSafe2++
			continue
		}

		for i := range report {
			newreport := []int{}
			for j, v := range report {
				if i == j {
					continue
				}
				newreport = append(newreport, v)
			}

			if isSafe(newreport) {
				totalSafe2++
				break
			}
		}
	}

	fmt.Println("totalSafe", totalSafe)
	fmt.Println("totalSafe2", totalSafe2)
}

func isSafe(r []int) bool {
	isdesc, isasc := true, true

	if len(r) <= 1 {
		return true
	}

	for i := 1; i < len(r); i++ {
		if r[i-1] < r[i] && isdesc {
			isdesc = false
		}
		if r[i-1] > r[i] && isasc {
			isasc = false
		}

		if !isdesc && !isasc {
			return false
		}

		df := util.Diff(r[i], r[i-1])
		if df < 1 || df > 3 {
			return false
		}
	}

	return true
}
