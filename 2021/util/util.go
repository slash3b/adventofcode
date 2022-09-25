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
