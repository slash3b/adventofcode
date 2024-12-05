package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/slash3b/aoc/util"
)

var (
	input = flag.String("input", "", "file name")
	reg   = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
)

func main() {
	flag.Parse()

	lines := util.FileToStrings(*input)
	total := 0
	cnt := 0

	for _, ln := range lines {
		for _, v := range reg.FindAllString(ln, -1) {
			cnt++
			total += mul(v)
		}
	}

	total2 := 0
	for _, ln := range lines {
		// ok := true
		res := strings.Split(ln, "don't()")

		for _, v := range reg.FindAllString(res[0], -1) {
			total2 += mul(v)
		}

		for i := 1; i < len(res); i++ {
			res2 := strings.Split(res[i], "do()")
			if len(res2) < 2 {
				continue
			}

			for _, v := range reg.FindAllString(res2[1], -1) {
				total2 += mul(v)
			}
		}
	}

	fmt.Println("cnt", cnt)
	fmt.Println("total", total)
	fmt.Println("total2", total2)
}

func mul(s string) int {
	s = s[4 : len(s)-1]
	nums := strings.Split(s, ",")

	if len(nums) != 2 {
		return 0
		panic(fmt.Sprintf("unexpected string %s", s))
	}

	return util.MustAtoi(nums[0]) * util.MustAtoi(nums[1])
}
