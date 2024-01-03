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

    // part 2

    type cardstat struct {
        num int // total amount of cards
        wins int // this card's wind, hence copies of cards
    }


    cards := make([]cardstat, len(lines))

    for k, line := range lines {
        win, nums := game(line)

        wins := 0
        for _, v := range nums {
            _, ok := win[v]
            if ok {
                wins++
            }
        }

        cards[k] = cardstat{
            num: 1,
            wins: wins,
        }
    }

    for k,v := range cards {
        // for each card
        for i:= v.num; i>0;i-- {
            // makes copies if any
            for i:= 0; i < v.wins; i++ {
                cards[k+i+1].num++
            }
        }
    }

    for _,v := range cards {
        res2 += v.num
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

