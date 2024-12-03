package main

import (
	"flag"
	"fmt"
    "strings"

	"github.com/slash3b/aoc/util"
)

var (
	input      = flag.String("input", "", "file name")
)

func main() {
	flag.Parse()

	ln := util.FileToStrings(*input)
    inp := ln[0]

    total := 0

    // it is late
    // I am tired
    // and I pretend I do
    // not know regular expressions
    // as if I do not have enough
    // suffering

    for i:= 0; i < len(inp); {
        if string(inp[i]) == "(" {
            for j := i+4; j < i+9; j++ {
                if string(inp[j]) == ")" {
                    total += mul(string(inp[i+1:j]))
                    i = j
                    break
                }
            }
        }
        i++
    }

    fmt.Println("total", total)
}

func mul( s string) int {
    nums := strings.Split(s, ",")
    if len(nums) != 2  {
        return 0
        panic(fmt.Sprintf("unexpected string %s", s))
    }

    return util.MustAtoi(nums[0]) * util.MustAtoi(nums[1])
}
