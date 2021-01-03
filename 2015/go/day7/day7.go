package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	RSHIFT = "RSHIFT"
	LSHIFT = "LSHIFT"
	AND    = "AND"
	OR     = "OR"
	NOT    = "NOT"
)

var cmds map[string]string
var computed map[string]int

func main() {
	cmds = make(map[string]string)
	computed = make(map[string]int)
	readAndPrepare("/home/slash3b/Projects/aoc/2015/go/day7/input")

	firstResult := calc("a")
	fmt.Println("FirstPart RESULT: ", calc("a"))
	computed = make(map[string]int)
	computed["b"] = firstResult
	fmt.Println("FINAL RESULT: ", calc("a"))
}
func readAndPrepare(file string) {
	res, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(res), "\n")

	for _, v := range inputs {
		splitRes := strings.Split(v, "->")
		cmd := strings.TrimSpace(splitRes[0])
		assignable := strings.TrimSpace(splitRes[1])
		cmds[assignable] = cmd
	}
}
func calc(key string) int {
	key = strings.TrimSpace(key)

	if value, exists := computed[key]; exists {
		return value
	}

	command, exists := cmds[key]
	if !exists {
		panic("empty command")
	}

	if strings.Contains(command, OR) {
		res := strings.Split(command, OR)

		a, err := strconv.Atoi(strings.TrimSpace(res[0]))
		if err != nil {
			a = calc(res[0])
			computed[strings.TrimSpace(res[0])] = a
		}

		b, err := strconv.Atoi(strings.TrimSpace(res[1]))
		if err != nil {
			b = calc(res[1])
			computed[strings.TrimSpace(res[1])] = b
		}

		return a | b
	} else if strings.Contains(command, NOT) {
		res := strings.Split(command, NOT)

		return ^calc(res[1])
	} else if strings.Contains(command, AND) {
		res := strings.Split(command, AND)

		a, err := strconv.Atoi(strings.TrimSpace(res[0]))
		if err != nil {
			a = calc(res[0])
			computed[strings.TrimSpace(res[0])] = a
		}

		b, err := strconv.Atoi(strings.TrimSpace(res[1]))
		if err != nil {
			b = calc(res[1])
			computed[strings.TrimSpace(res[1])] = b
		}

		return a & b
	} else if strings.Contains(command, RSHIFT) {
		res := strings.Split(command, RSHIFT)
		a, err := strconv.Atoi(strings.TrimSpace(res[1]))
		if err != nil {
			panic(err)
		}

		result := calc(res[0])
		computed[strings.TrimSpace(res[0])] = result
		return result >> a
	} else if strings.Contains(command, LSHIFT) {
		res := strings.Split(command, LSHIFT)
		a, err := strconv.Atoi(strings.TrimSpace(res[1]))
		if err != nil {
			panic(err)
		}

		result := calc(res[0])
		computed[strings.TrimSpace(res[0])] = result
		return result << a
	} else if strings.Contains(command, "lx") {
		return calc(command)
	}

	res, err := strconv.Atoi(strings.TrimSpace(command))
	if err != nil {
		panic(err)
	}

	return res
}
