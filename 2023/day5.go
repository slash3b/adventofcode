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

type FromTo struct {
    From, To int
}

var (
    m = make(map[string]map[int]int)
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

            m[key] = make(map[int]int)

            continue
        }


        res := strings.Split(lines[i], " ")
        a,b,c := util.MustAtoi(res[0]), util.MustAtoi(res[1]), util.MustAtoi(res[2])
        for i:=0;i < c; i++ {

            m[key][b] = a

            a++
            b++
        }

    }

    /*
    for k, v := range m {
        fmt.Println(k)

        for j, k := range v {
            fmt.Printf("\t %d:%d\n", j,k)
        }
    }
    */

    min := 0
    for _, seed := range seeds {
        fmt.Println("seed", seed)

        d :=  toDest(seed, m)
        if min == 0 || d < min {
            min = d
        }

        return 

    }
    res1 = min


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

func toDest(seed int, m map[string]map[int]int) int {
    soil, ok := m["seed-to-soil"][seed]
    if !ok {soil = seed}

    fert, ok := m["soil-to-fertilizer"][soil]
    if !ok {fert = soil}

    water, ok := m["fertilizer-to-water"][fert]
    if !ok {water = fert}

    light, ok := m["water-to-light"][water]
    if !ok {light = water}

    tmp, ok := m["light-to-temperature"][light]
    if !ok {tmp = light}

    hu, ok := m["temperature-to-humidity"][tmp]
    if !ok {hu = tmp}

    loc, ok := m["humidity-to-location"][hu]
    if !ok {loc = hu}

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
