package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("--------------------")
}

func getInput(s string) []byte {
	b, err := os.ReadFile(s)
	if err != nil {
		panic(err)
	}

	return b
}

var reg = regexp.MustCompile("[a-z]*")

func day1_1() string {
	input := string(getInput("inputs/day1_puzzle"))

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
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	input := string(getInput("inputs/day1_puzzle"))

	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	var ans int

	for _, line := range lines {
		var (
			a      int
			afound bool
			b      int
		)

		for i := 0; i < len(line); i++ {
			v, ok := byteToDecimal(line[i])
			if ok {
				if !afound {
					a = v
					afound = true
				}
				b = v
			}

			for k, v := range m {
				// if rest of the line has sufficient lenght
				// then try to compare
				if len(k) <= len(line[i:]) && k == line[i:i+len(k)] {
					fmt.Println("comparing", k, line[i:i+len(k)])

					if !afound {
						a = v
						afound = true
					}

					b = v
				}
			}

		}

		fmt.Println("LINE:", line, "a", a, "b", b, ":", a*10+b)
		ans += a*10 + b
	}

	// part2 answer is: 54591
	return fmt.Sprintf("PART 2 ANS: %d", ans)
}

var asciiInts = []byte{0x30: 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0x7a: 0}

func byteToDecimal(b byte) (int, bool) {
	if asciiInts[b] == 1 {
		return int(b - 0x30), true
	}

	return 0, false
}
