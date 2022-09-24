package day1_test

import (
	"adventofcode/day1"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestDay1_SampleTest(t *testing.T) {

	fmt.Println("SAMPLE")
	res := FileToIntSlice("./test_data")

	fmt.Println("Result", day1.Run(res))
}

func TestDay1_PuzzleTest(t *testing.T) {

	fmt.Println("PUZZLE")
	res := FileToIntSlice("./puzzle_data")

	fmt.Println("Result", day1.Run(res))
}

func TestDay2_SampleTest(t *testing.T) {

	fmt.Println("SAMPLE day2")
	res := FileToIntSlice("./test_data")

	fmt.Println("Result", day1.Run2(res))
}

func TestDay2_PuzzleTest(t *testing.T) {

	fmt.Println("PUZZLE day2")
	res := FileToIntSlice("./puzzle_data")

	fmt.Println("Result", day1.Run2(res))
}

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
