package util

import (
    "os"
    "strings"
)

func FileToStrings(input string) []string {
    b, err := os.ReadFile(input)
    if err != nil {
        panic(err)
    }

    in := strings.TrimRight(string(b), "\n")

    return strings.Split(in, "\n")
}
