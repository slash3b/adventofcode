package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	readAndPrepare_part2("/home/slash3b/Projects/aoc/2015/go/day12/input")
}

func readAndPrepare_part2(s string) {
	res, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}

	var data interface{}

	json.Unmarshal(res, &data)

	fmt.Println("RESULT:", unfold(data)) // 65402 is the answer
}

func unfold(s interface{}) int {

	var total int

outer:
	switch data := s.(type) {
	case []interface{}:
		for _, value := range data {
			if v, success := value.(float64); success {
				total += int(v)
				continue
			}
			if _, success := value.(string); success {
				continue
			}

			total += unfold(data)
		}
	case map[string]interface{}:
		if isRed(data) {
			break outer
		}

		for _, value := range data {
			if v, success := value.(float64); success {
				total += int(v)
				continue
			}
			if _, success := value.(string); success {
				continue
			}

			total += unfold(data)
		}
	}

	return total
}

func isRed(structData map[string]interface{}) bool {
	for _, v := range structData {
		if stringValue, success := v.(string); success && stringValue == "red" {
			return true
		}

	}

	return false
}
