package main

import (
	"flag"
	"fmt"

    "strings"

	"aoc/util"
)

var (
	input      = flag.String("input", "", "file name")
	res1, res2 = 0, 0
)

func main() {
	flag.Parse()

	lines := util.FileToStrings(*input)
    for _, line := range lines {
        win, nums := game(line)

        cnt := 0
        found := false

        for _, v := range nums {
            _, ok := win[v]
            if ok {
                if !found { found = true; continue;}
                cnt++
            }
        }

        if found {
            res1 += 1 << cnt
        }
    }

	fmt.Println(res1, res2)
}

func game(line string) (map[string]struct{}, []string) {
        res := strings.Split(line, ":")
        res = strings.Split(res[1], "|")

        win := make(map[string]struct{})
        for _, wv := range strings.Split(strings.TrimSpace(res[0]), " ") {
            if wv == "" { continue }
            win[wv] = struct{}{}
        }

        have := make([]string, 0)
        for _, num := range strings.Split(strings.TrimSpace(res[1]), " ") {
            if num == "" { continue }
            have = append(have, num)
        }

        return win, have
}

