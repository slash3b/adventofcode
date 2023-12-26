package main

import (
    "fmt"
    "flag"
    "os"
    "strings"
    "strconv"
)

var input = flag.String("input", "", "file name")

func main(){
    flag.Parse()

    b, err := os.ReadFile(*input)
    if err != nil {
        panic(err)
    }

    in := strings.TrimRight(string(b), "\n")

    day1res := 0

    for _, v := range strings.Split(in, "\n") {

        res := strings.Split(v, ":")
        if len(res) != 2 {
            panic(fmt.Sprintf("unexpected split result from %s", v ))
        }

        gameID := strings.TrimLeft(res[0], "Game ")
        fmt.Println("game ID", gameID)

        valid := true
        for _, set := range strings.Split(res[1], ";") {
            cubes := fromStr(set)
            valid = cubes.Valid()
            if !valid {
                fmt.Printf("\t%#v\n", cubes)
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
    // 12 red, 13 green, 14 blue

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


