package main

import "testing"

func TestIsMatchingAunt(t *testing.T) {
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

	check := make(map[string]int)
	check["akitas"] = 0
	check["vizslas"] = 0
	check["samoyeds"] = 2

	if !isMatchingAunt(target, check) {
		t.Error("oops")
	}

	check2 := make(map[string]int)
	check2["akitas"] = 1
	check2["vizslas"] = 0
	check2["samoyeds"] = 2

	if isMatchingAunt(target, check2) {
		t.Error("oops")
	}
}
