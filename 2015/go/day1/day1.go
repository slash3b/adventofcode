package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	b, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	var counter int
	var wentToBasement bool
	for x := range b {
		chr := string(b[x])
		if chr == "(" {
			counter++
		} else {
			counter--
		}
		if counter == -1 && !wentToBasement {
			wentToBasement = true
			fmt.Println("First char that led to basement: ", x + 1)
		}
	}

	fmt.Println(counter)
}
