package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var total int

func main() {
	readAndPrepare_part2("/home/slash3b/Projects/aoc/2015/go/day12/input")
}

func readAndPrepare_part2(s string) {
	res, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}

	unfold(res, &total)

	fmt.Println("RESULT:", total) // 65402 is the answer
}

func unfold(s []byte, total *int) {

	var sliceData []interface{}
	err := json.Unmarshal(s, &sliceData)
	if err == nil {
		for _, data := range sliceData {
			if v, success := data.(float64); success {
				*total += int(v)
				continue
			}
			if _, success := data.(string); success {
				continue
			}

			f, _ := json.Marshal(data)
			unfold(f, total)
		}
	}

	var structData map[string]interface{}

	err = json.Unmarshal(s, &structData)

	if err == nil {
		// quickly check if it is there
		if !isRed(structData) {
			for _, data := range structData {
				//fmt.Printf("%#v %[1]T \n", data)
				if v, success := data.(float64); success {
					*total += int(v)
					continue
				}
				if _, success := data.(string); success {
					continue
				}

				f, _ := json.Marshal(data)
				unfold(f, total)
			}
		}
	}
}

func isRed(structData map[string]interface{}) bool {
	for _, v := range structData {
		if stringValue, success := v.(string); success && stringValue == "red" {
			return true
		}

	}

	return false
}
