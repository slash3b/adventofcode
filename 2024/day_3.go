package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
    "slices"

	"github.com/slash3b/aoc/util"
)

var (
	input = flag.String("input", "", "file name")
	reg   = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	dont   = regexp.MustCompile(`don't\(\)`)
	do   = regexp.MustCompile(`do\(\)`)
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

	for _, ln := range lines{
        muls := reg.FindAllStringIndex(ln, -1)
        donts := dont.FindAllStringIndex(ln, -1)
        dos := do.FindAllStringIndex(ln, -1)

        muls = append(muls, donts...)
        muls = append(muls, dos...)

        slices.SortFunc(muls, func(a,b []int) int {
            if a[0] > b[0] {return 1}
            if a[0] < b[0] {return -1}
            return 0
        })

        enabled := true
        res := 0
        for _, v := range muls {
            res, enabled = mul2(ln[v[0]:v[1]], enabled)
            if enabled {
                fmt.Println("added ", res, "from", ln[v[0]:v[1]])
            } else {
                fmt.Println("ignoring", ln[v[0]:v[1]])
            }


            total2 += res
        }

        break
	}


	fmt.Println("cnt", cnt)
	fmt.Println("total", total)
	fmt.Println("total2", total2)
}

func mul2(s string, enabled bool) (int, bool) {
    switch s {
    case "don't()":
        return 0, false
    case "do()":
        return 0, true
    }

    if enabled {
        return mul(s), true
    }

    return 0, false
}

func mul(s string) int {
    fmt.Println("incoming:",s)
	s = s[4 : len(s)-1]
	nums := strings.Split(s, ",")

	if len(nums) != 2 {
		return 0
		panic(fmt.Sprintf("unexpected string %s", s))
	}

    fmt.Println("parsed:",util.MustAtoi(nums[0]), util.MustAtoi(nums[1]))
	s = s[4 : len(s)-1]
	return util.MustAtoi(nums[0]) * util.MustAtoi(nums[1])
}
