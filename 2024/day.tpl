package main

import (
	"flag"
	"fmt"

	"aoc/util"
)

var (
	input      = flag.String("input", "", "file name")
	res1, res2 = 0, 0
)

func main() {
	flag.Parse()

	lines := util.FileToStrings(*input)
	_ = lines

	fmt.Println(res1, res2)
}
