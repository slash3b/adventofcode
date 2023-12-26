package main

import (
    "fmt"
    "flag"
    "strings"
    "strconv"

    "aoc/util"
)

var input = flag.String("input", "", "file name")

func main(){
    flag.Parse()

    lines := util.FileToStrings(*input)

    day1res := 0

    for _, v := range lines {

        res := strings.Split(v, ":")
        if len(res) != 2 {
            panic(fmt.Sprintf("unexpected split result from %s", v ))
        }

        gameID := strings.TrimLeft(res[0], "Game ")

        valid := true
        for _, set := range strings.Split(res[1], ";") {
            cubes := fromStr(set)
            valid = cubes.Valid()
            if !valid {
                break
            }
        }

        if !valid {
            continue
        }

        id, _ := strconv.Atoi(gameID)
        day1res += id

    }

    fmt.Println("Day 1 res:", day1res)

    day2res := 0

    for _, v := range lines {

        res := strings.Split(v, ":")
        if len(res) != 2 {
            panic(fmt.Sprintf("unexpected split result from %s", v ))
        }

        fc := Cubes{}
        for _, set := range strings.Split(res[1], ";") {
            fc.Take(fromStr(set))
        }

        day2res += fc.red * fc.green * fc.blue
    }

    fmt.Println("Day 2 res:", day2res)
}

type Cubes struct {
    red int
    green int
    blue int
}


func (c Cubes) Valid() bool {
    if c.red > 12 {return false}
    if c.green > 13 {return false}
    if c.blue > 14 {return false}

    return true
}

func (c *Cubes) Take(from Cubes) {
    if from.red > c.red {c.red = from.red}
    if from.green > c.green {c.green = from.green}
    if from.blue > c.blue {c.blue = from.blue}
}

func fromStr(s string) Cubes {

    c := Cubes{}

    res := strings.Split(s, ",")

    for _, v := range res {
        props := strings.Split(strings.Trim(v, " "), " ")

        val, err := strconv.Atoi(props[0])
        if err != nil {
            panic(fmt.Sprintf("unable to convert to int %v", props[0]))
        }

        switch props[1] {
        case "blue":
            c.blue = val
        case "red":
            c.red = val
        case "green":
            c.green = val
        default:
            panic(fmt.Sprintf("unexpected property from %s", v))
        }
    }

    return c
}


