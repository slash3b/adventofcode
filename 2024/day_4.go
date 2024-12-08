package main

import (
	"flag"
	"fmt"

	"github.com/slash3b/aoc/util"
)

var (
	input      = flag.String("input", "", "file name")
)

func main() {
	flag.Parse()

	lines := util.FileToStrings(*input)
    fmt.Println(len(lines))
}
