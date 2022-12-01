package util

import (
	"bufio"
	"bytes"
	"fmt"
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

		if elems := strings.Split(v, ","); len(elems) > 1 {
			for _, el := range elems {
				intRes, err := strconv.Atoi(el)
				if err != nil {
					panic(err)
				}

				res = append(res, intRes)
			}
			continue
		}

		intRes, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		res = append(res, intRes)
	}

	return res
}

func FileContentsSplitByNewLine(filePath string) []string {
	bs, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(bs), "\n")
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

func ScanFile(filePath string) {
	bs, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(bs))
	i := 0
	for scanner.Scan() {
		fmt.Println(i)
		fmt.Printf("%#v\n", scanner.Text())
		fmt.Printf("%#v\n", scanner.Bytes())
		i++
	}
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
