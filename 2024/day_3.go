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
	reg   = regexp.MustCompile(`(?m)\(\d{1,3},\d{1,3}\)`)
)

func main() {
	flag.Parse()

	ln := util.FileToStrings(*input)
	inp := ln[0]

	total := 0

	fmt.Println("matches count", len(reg.FindAllString(inp, -1)))
	for _, v := range reg.FindAllString(inp, -1) {
		total += mul(v)
	}

	fmt.Println("total", total)

	return

	for i := 0; i < len(inp); {
		if string(inp[i]) == "(" {
			for j := i + 4; j < i+9; j++ {
				if string(inp[j]) == ")" {
					total += mul(string(inp[i+1 : j]))
					i = j
					break
				}
			}
		}
		i++
	}

	fmt.Println("total", total)
}

func mul(s string) int {

	s = s[1 : len(s)-1]
	nums := strings.Split(s, ",")
	if len(nums) != 2 {
		return 0
		panic(fmt.Sprintf("unexpected string %s", s))
	}

	return util.MustAtoi(nums[0]) * util.MustAtoi(nums[1])
}
