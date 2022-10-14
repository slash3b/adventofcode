package util

import (
	"os"
	"strconv"
	"strings"
)

func FileToIntSlice(filePath string) []int {
	bs, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	res := []int{}

	input := strings.Split(string(bs), "\n")

	lastElIndex := len(input) - 1
	if len(input[lastElIndex]) == 0 {
		input = input[:lastElIndex]
	}

	// now conversion to type

	for _, v := range input {
		intRes, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		res = append(res, intRes)

	}

	return res
}

func FileContents(filePath string) string {
	bs, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(bs)
}

func FileToStringSlice(filePath string) []string {
	bs, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(bs), "\n")

	lastElIndex := len(input) - 1
	if len(input[lastElIndex]) == 0 {
		input = input[:lastElIndex]
	}

	return input
}

func Min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Diff(a, b int) int {
	res := a - b
	if res < 0 {
		res *= -1
	}

	return res
}
