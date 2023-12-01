package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getInput(s string) []byte {
	b, err := os.ReadFile(s)
	if err != nil {
		panic(err)
	}

	return b
}

var reg = regexp.MustCompile("[a-z]*")

func day1_1() string {
	input := string(getInput("day1_puzzle"))

	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	var ans int

	for _, v := range lines {
		fmt.Println("orig", v)
		in := reg.ReplaceAllString(v, "")
		res, err := strconv.Atoi(string(in[0]) + string(in[len(in)-1]))
		if err != nil {
			panic(err)
		}

		ans += res

	}

	return fmt.Sprintf("PART 1 ANS: %d", ans)
}

func day1_2() string {

	m := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	_ = m

	input := string(getInput("day1_puzzle_2_test"))

	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	var ans int

	for _, v := range lines {
		fmt.Println("orig", v)

		for k, j := range m {
			v = strings.ReplaceAll(v, k, j)
		}

		fmt.Println("normalized", v)

		in := reg.ReplaceAllString(v, "")
		fmt.Println("replaced", in)

		res, err := strconv.Atoi(string(in[0]) + string(in[len(in)-1]))
		if err != nil {
			panic(err)
		}

		ans += res
		fmt.Println("----")

	}

	return fmt.Sprintf("PART 2 ANS: %d", ans)
}
