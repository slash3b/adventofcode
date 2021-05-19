package main

import (
	"fmt"
	"github.com/slash3b/permutation"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	//fmt.Println(permutation.Strings([]string{"a", "b", "c", "d", "e", "f", "g"}))

	readAndPrepare("/home/slash3b/Projects/aoc/2015/go/day13/input")
	readAndPrepareSecondPart("/home/slash3b/Projects/aoc/2015/go/day13/input")

}

func readAndPrepare(s string) {
	res, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(res), "\n")

	names := make(map[string]bool)
	scores := make(map[string]int)

	for _, v := range inputs {
		v = strings.TrimRight(v, ".")
		line := strings.Split(v, " ")
		if _, exists := names[line[0]]; !exists {
			names[line[0]] = true
		}

		// [Bob would gain 89 happiness units by sitting next to David]
		score, _ := strconv.Atoi(line[3])
		if line[2] == "lose" {
			score = score * -1
		}
		key := line[0] + line[len(line)-1]
		scores[key] = score
	}
	var nameSet []string
	for k, _ := range names {
		nameSet = append(nameSet, k)
	}

	perms := permutation.Strings(nameSet)
	//fmt.Println("Names", names)
	//fmt.Println("Scores: ", scores)
	//fmt.Println("Permutations: ", perms)

	calcHappiness := func(names []string) int {
		var init int
		for i := 0; i < len(names)-1; i++ {
			init += scores[names[i]+names[i+1]]
			init += scores[names[i+1]+names[i]]
		}
		// last sitting to first
		init += scores[names[len(names)-1]+names[0]]
		init += scores[names[0]+names[len(names)-1]]

		return init
	}

	maxHappiness := 0
	//maxNames := []string{}
	for _, permNames := range perms {
		currentHappiness := calcHappiness(permNames)
		//fmt.Println(permNames, "-->", currentHappiness)
		if currentHappiness > maxHappiness {
			maxHappiness = currentHappiness
			//maxNames = permNames
		}
	}

	fmt.Println("ANSWER: ", maxHappiness)
	//fmt.Println("CHOSEN names ", maxNames)

}

func readAndPrepareSecondPart(s string) {
	res, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(res), "\n")

	names := make(map[string]bool)
	scores := make(map[string]int)

	for _, v := range inputs {
		v = strings.TrimRight(v, ".")
		line := strings.Split(v, " ")
		if _, exists := names[line[0]]; !exists {
			names[line[0]] = true
		}

		// [Bob would gain 89 happiness units by sitting next to David]
		score, _ := strconv.Atoi(line[3])
		if line[2] == "lose" {
			score = score * -1
		}
		key := line[0] + line[len(line)-1]
		scores[key] = score
	}

	// include myself in the list!
	names["myself"] = true
	var nameSet []string
	for k, _ := range names {
		nameSet = append(nameSet, k)
	}

	perms := permutation.Strings(nameSet)
	//fmt.Println("Names", names)
	//fmt.Println("Scores: ", scores)
	//fmt.Println("Permutations: ", perms)

	calcHappiness := func(names []string) int {
		var init int
		for i := 0; i < len(names)-1; i++ {
			if names[i] == "myself" || names[i+1] == "myself" {
				continue
			}
			init += scores[names[i]+names[i+1]]
			init += scores[names[i+1]+names[i]]
		}

		if names[len(names)-1] == "myself" || names[0] == "myself" {
			return init
		}
		// last sitting to first
		init += scores[names[len(names)-1]+names[0]]
		init += scores[names[0]+names[len(names)-1]]

		return init
	}

	maxHappiness := 0
	maxNames := []string{}
	for _, permNames := range perms {
		currentHappiness := calcHappiness(permNames)
		//fmt.Println(permNames, "-->", currentHappiness)
		if currentHappiness > maxHappiness {
			maxHappiness = currentHappiness
			maxNames = permNames
		}
	}

	fmt.Println("ANSWER: ", maxHappiness)
	fmt.Println("CHOSEN names ", maxNames)

}
