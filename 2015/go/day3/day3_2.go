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

	fmt.Println(length)
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

	var (
		santaRow = length+1
		santaColumn = length+1
		roboRow = santaRow
		roboColumn = santaColumn
	)

	var counter int
	counter++
	santaMap[santaColumn][santaRow] = true
	for x := range res {
		dir := string(res[x])
		if x % 2 == 0 {
			if dir == ">" {
				santaRow++
			} else if dir == "<" {
				santaRow--
			} else if dir == "^" {
				santaColumn++
			} else {
				santaColumn--
			}
			if santaMap[santaColumn][santaRow] == false {
				counter++
				santaMap[santaColumn][santaRow] = true
			}
		} else {
			if dir == ">" {
				roboRow++
			} else if dir == "<" {
				roboRow--
			} else if dir == "^" {
				roboColumn++
			} else {
				roboColumn--
			}
			if santaMap[roboColumn][roboRow] == false {
				counter++
				santaMap[roboColumn][roboRow] = true
			}
		}
	}

	fmt.Println("Num of hourses with at least one present: ", counter)
}
