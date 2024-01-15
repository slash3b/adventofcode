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

type ConvMap []FromTo

func (c ConvMap) Dest(x int) int {
    for _, v := range c {
        if v.Contains(x) {
            return v.Dest(x)
        }
    }


    return x // in case of a miss destination equals to x
}

type FromTo struct {
    From [2]int
    To [2]int
}


// FromRange constructor
func FromRange(a,b,c int) FromTo {
    return FromTo{
        From: [2]int{a, a + c - 1},
        To: [2]int{b, b + c - 1},
    }
}

func (f FromTo) Dest(x int) int {
    offset := x - f.From[0]

    return f.To[0] + offset
}

func (f FromTo) Contains(x int) bool {
    return f.From[0] <= x && x <= f.From[1]
}

var (
    m = make(map[string]ConvMap)
)

func main() {
	flag.Parse()

	lines := util.FileToStrings(*input)


    // work on seeds
    sres, _ := strings.CutPrefix(lines[0], "seeds: ")

    seeds := []int{}

    for _, v := range strings.Split(sres, " ") {
        seeds = append(seeds, util.MustAtoi(v))
    }

    fmt.Println("seeds:", seeds)

    // now every map

    key := ""
    for i:= 1; i < len(lines);i++ {
        // deal with key
        if lines[i] == "" {
            key = strToKey(lines[i+1])

            fmt.Println("new key", key)
            i++

            // m[key] = make(map[int]int)

            continue
        }


        res := strings.Split(lines[i], " ")
        a,b,c := util.MustAtoi(res[0]), util.MustAtoi(res[1]), util.MustAtoi(res[2])
        m[key] = append(m[key], FromRange(b, a, c)) // have to reverse
    }

    fmt.Println("-------------------------------------")

    min := 0
    for _, seed := range seeds {
        fmt.Println("seed", seed)

        d := toDest(seed, m)
        if min == 0 || d < min {
            min = d
        }
    }

    res1 = min


    fmt.Println("--------------------")
	fmt.Println(res1, res2)

    // experimentation starts here

    // input := strings.Split(`50 98 2
// 52 50 48
// `, "\n")

    // for _, v := range input[:len(input)-1] {
    //     if len(res) != 3 {
    //         panic("unexpected len")
    //     }

    //     for i:=0;i < c; i++ {
    //         _ = FromTo{
    //             From: b,
    //             To: a,
    //         }


    //         seedToSoil[b] = a

    //         a++
    //         b++
    //     }

    //     for k, _ := range seedToSoil {
    //         fmt.Println(k)
    //     }
    //     fmt.Println("----------------")

    // }

    // the action

}

func toDest(seed int, m map[string]ConvMap) int {
    soil := m["seed-to-soil"].Dest(seed)

    fert := m["soil-to-fertilizer"].Dest(soil)

    water := m["fertilizer-to-water"].Dest(fert)

    light := m["water-to-light"].Dest(water)

    tmp := m["light-to-temperature"].Dest(light)

    hu := m["temperature-to-humidity"].Dest(tmp)

    loc := m["humidity-to-location"].Dest(hu)

    return loc
}

func strToKey(s string) string {
    if strings.Contains(s, "seed-to-soil") {return "seed-to-soil"}
    if strings.Contains(s, "soil-to-fertilizer") { return "soil-to-fertilizer"}
    if strings.Contains(s, "fertilizer-to-water") { return "fertilizer-to-water"}
    if strings.Contains(s, "water-to-light") { return "water-to-light"}
    if strings.Contains(s, "light-to-temperature") { return "light-to-temperature"}
    if strings.Contains(s, "temperature-to-humidity") { return "temperature-to-humidity"}
    if strings.Contains(s, "humidity-to-location") {return "humidity-to-location"}

    panic("unexpected string")
}
