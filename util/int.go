package util

import (
    "strconv"
    "fmt"
)

func Diff(a, b int) int {
    if a > b {
        return a - b
    }

    return b - a
}

func MustAtoi(s string) int {
    i, err := strconv.Atoi(s)
    if err != nil {
        panic(fmt.Sprintf("could not convert %s, err: %v", s, err))
    }

    return i
}
