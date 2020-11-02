package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	res, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}



	length := len(res)
	quadRadius := len(res)*2+1
	fmt.Println("what is the radius: ", quadRadius)

	var santaMap [][]bool
	for i:=0; i <= quadRadius; i++ {
		var row []bool
		for j:= 0; j <= quadRadius; j++ {
			row = append(row, false)
		}

		santaMap = append(santaMap, row)
	}

	row := length+1
	column := length+1

	var counter int
	counter++
	santaMap[column][row] = true
	for x := range res {
		dir := string(res[x])

		if dir == ">" {
			row++
		} else if dir == "<" {
			row--
		} else if dir == "^" {
			column++
		} else {
			column--
		}
		if santaMap[column][row] == false {
			counter++
			santaMap[column][row] = true
		}
	}

	fmt.Println(counter)
}
