package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

//func main() {
//	readAndPrepare("/home/slash3b/Projects/aoc/2015/go/day12/input")
//}

func readAndPrepare(s string) {
	res, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}

	//fmt.Println("DEBUG: ", string(res))

	ints := regexp.MustCompile(`\-?\d+`).FindAllString(string(res), -1)
	fmt.Println(ints)

	var sum int
	for _, value := range ints {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			panic(value)
		}
		sum += intValue
	}

	fmt.Println(sum)

}
