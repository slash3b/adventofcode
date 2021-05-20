package main

import (
	"fmt"
	"github.com/slash3b/permutation"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	readAndPrepare("/home/slash3b/Projects/aoc/2015/go/day9/input")
}

func readAndPrepare(s string) {
	res, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(res), "\n")

	names := make(map[string]bool)
	distances := make(map[string]int)

	for _, v := range inputs {
		line := strings.Split(v, " ")
		nameA := line[0]
		if _, exists := names[nameA]; !exists {
			names[nameA] = true
		}

		nameB := line[2]
		if _, exists := names[nameB]; !exists {
			names[nameB] = true
		}

		distance, err := strconv.Atoi(line[4])
		if err != nil {
			panic(err)
		}

		key1 := nameA + nameB
		if _, exists := distances[key1]; !exists {
			distances[key1] = distance
		}
		key2 := nameB + nameA
		if _, exists := distances[key2]; !exists {
			distances[key2] = distance
		}
	}

	//fmt.Println(names)
	//fmt.Println(distances)

	cities := []string{}
	for k, _ := range names {
		cities = append(cities, k)
	}

	routes := permutation.Strings(cities)
	//fmt.Println(routes)

	calcDistance := func(r []string) int {
		res := 0

		for i := 0; i < len(r)-1; i++ {
			key := r[i] + r[i+1]
			v, exists := distances[key]
			if !exists {
				return 0
			}
			res += v
		}

		return res
	}

	minDistance := 1<<63 - 1
	maxDistance := 0

	for _, route := range routes {
		distResponse := calcDistance(route)
		if distResponse > 0 && distResponse < minDistance {
			minDistance = distResponse
		}
		if distResponse > 0 && distResponse > maxDistance {
			maxDistance = distResponse
		}
	}

	fmt.Println("ANSWER: min", minDistance)
	fmt.Println("ANSWER: max", maxDistance)
}
