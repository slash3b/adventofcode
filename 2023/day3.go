package main

import (
    "flag"

    "fmt"
    "aoc/util"
)

var input = flag.String("input", "", "file name")

func main(){
    flag.Parse()

    lines := util.FileToStrings(*input)

    respart1 := 0

    fmt.Println("len lines", len(lines))

    n := []int{}
    for x, line := range lines {
        fmt.Println("line: ", line)

        if len(n) > 0 {
        panic("fooo")
    }
        for y:= 0; y < len(line); y++ {
            v, ok := util.ByteToDecimal(line[y])
            if ok {
                n = append(n, v)
                if y != len(line)-1 {
                    continue
                }
            }

            // check number
            if len(n) > 0 {

                fmt.Printf("DEBUG: slice %v\n", n)
                valid := false
                y2 := y
                for i:= len(n) -1; i >= 0; i-- {
                    // fmt.Printf("DEBUG: checking number %d\n", n[i])
                    // tail
                    if i == len(n) - 1 {
                        if hasSymbol(x-1, y2, lines) {
                            valid=true
                            break
                        }
                        if hasSymbol(x, y2, lines) {
                            valid=true
                            break
                        }
                        if hasSymbol(x+1, y2, lines) {
                            valid=true
                            break
                        }
                    }

                    // middle
                    if hasSymbol(x+1, y2-1, lines) || hasSymbol(x-1, y2-1, lines) {
                        valid=true
                        break
                    }

                    // head
                    if i == 0 {
                        if hasSymbol(x-1, y2-2, lines) {
                            valid=true
                            break
                        }
                        if hasSymbol(x, y2-2, lines) {
                            valid=true
                            break
                        }
                        if hasSymbol(x+1, y2-2, lines) {
                            valid=true
                            break
                        }
                    }

                    y2--
                }

                // check every position
                // y, x
                if valid {
                    fmt.Printf("\tvalid %v\n", n)
                    // assemble whole interger value out of int slice
                    res := 0
                    m := 1
                    for i:= len(n) -1; i >= 0; i-- {
                        res += n[i] * m
                        m = m * 10
                    }
                    fmt.Println("adding", res)

                    respart1 += res
                }

                ///fmt.Printf("\t%#v\n",n)
                // clean up
                n = []int{}
            }

        }

    }

    fmt.Println("Part 1 result:", respart1)
}

func hasSymbol(x, y int, lines []string) bool {
    // fmt.Printf("\t coords x: %d, y: %d", x, y)

    // simple check
    if y < 0 || y > 139 { fmt.Println();return false }
    if x < 0 || x > 139 { fmt.Println();return false }

    b := lines[x][y]

    // fmt.Printf(">>foundValue %s", string(b))

    _, ok := util.ByteToDecimal(b)

    isDot := b == 0x2e

    // fmt.Printf(">>isSymbol %v\n", !(ok || isDot))

    return !(ok || isDot)
}
