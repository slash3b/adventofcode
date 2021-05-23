package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//type auntProps struct {
//	children
//	cats
//	samoyeds
//	pomeranians
//	akitas
//	vizslas
//	goldfish
//	trees
//	cars
//	perfumes
//}

func main() {
	readAndPrepare("/home/slash3b/Projects/aoc/2015/go/day16/input")

	/*
		- figure which aunt Suie 1- 500
		- MFCSAM can detect kinds of compounds
		- your aunt has these
			children: 3
			cats: 7
			samoyeds: 2
			pomeranians: 3
			akitas: 0
			vizslas: 0
			goldfish: 5
			trees: 3
			cars: 2
			perfumes: 1
	*/
	// try to get all aunts from the list who has all the properties above and amount exactly of 0 from your side of things
	// lets see
}

type props map[string]int

func readAndPrepare(s string) {
	res, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}

	allAunts := make(map[int]props)

	inputs := strings.Split(string(res), "\n")

	for _, v := range inputs {
		line := strings.Split(v, " ")
		number, _ := strconv.Atoi(strings.TrimRight(line[1], ":"))
		prop1Name := strings.TrimRight(line[2], ":")
		prop1Value, _ := strconv.Atoi(strings.TrimRight(line[3], ","))
		prop2Name := strings.TrimRight(line[4], ":")
		prop2Value, _ := strconv.Atoi(strings.TrimRight(line[5], ","))
		prop3Name := strings.TrimRight(line[6], ":")
		prop3Value, _ := strconv.Atoi(strings.TrimRight(line[7], ":"))

		properties := make(props)
		properties[prop1Name] = prop1Value
		properties[prop2Name] = prop2Value
		properties[prop3Name] = prop3Value

		allAunts[number] = properties
	}

	target := make(map[string]int)
	target["children"] = 3
	target["cats"] = 7
	target["samoyeds"] = 2
	target["pomeranians"] = 3
	target["akitas"] = 0
	target["vizslas"] = 0
	target["goldfish"] = 5
	target["trees"] = 3
	target["cars"] = 2
	target["perfumes"] = 1

	for key, aunt := range allAunts {
		if isMatchingAunt(target, aunt) {
			fmt.Println("matching Aunt: ", key, aunt)
		}
		if isRealAunt(target, aunt) {
			fmt.Println("matching REAL Aunt: ", key, aunt)
		}
	}

	//fmt.Println(allAunts)
}

func isMatchingAunt(target, check map[string]int) bool {

	for prop, value := range check {
		targetValue, _ := target[prop]
		// target value == 0 and then value might be anything ok
		// or target should be exactly == to value
		//fmt.Println("comparing:", prop, targetValue, value)
		if targetValue != value {
			return false
		}
	}

	return true
}

/*
	In particular, the cats and trees readings indicates that
	there are greater than that many
	(due to the unpredictable nuclear decay of cat dander and tree pollen),
	while the pomeranians and goldfish readings indicate that there are
	fewer than that many
	(due to the modial interaction of magnetoreluctance).
*/

func isRealAunt(target, check map[string]int) bool {
	for prop, value := range check {
		targetValue, _ := target[prop]

		// cats
		// trees
		// is >=
		// pomeranians
		// goldfish
		// <=

		if prop == "cats" || prop == "trees" {
			if value < targetValue {
				return false
			}
		} else if prop == "pomeranians" || prop == "goldfish" {
			if value > targetValue {
				return false
			}
		} else {
			if targetValue != value {
				return false
			}
		}
	}

	return true
}
