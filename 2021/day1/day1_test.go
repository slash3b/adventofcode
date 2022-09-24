package day1_test

import (
	"adventofcode/day1"
	"adventofcode/util"
	"fmt"
	"testing"
)

func TestDay1_SampleTest(t *testing.T) {

	fmt.Println("SAMPLE")
	res := util.FileToIntSlice("./test_data")

	fmt.Println("Result", day1.Run(res))
}

func TestDay1_PuzzleTest(t *testing.T) {

	fmt.Println("PUZZLE")
	res := util.FileToIntSlice("./puzzle_data")

	fmt.Println("Result", day1.Run(res))
}

func TestDay1_SampleTest2(t *testing.T) {

	fmt.Println("SAMPLE day2")
	res := util.FileToIntSlice("./test_data")

	fmt.Println("Result", day1.Run2(res))
}

func TestDay1_PuzzleTest2(t *testing.T) {

	fmt.Println("PUZZLE day2")
	res := util.FileToIntSlice("./puzzle_data")

	fmt.Println("Result", day1.Run2(res))
}
