package main

import "fmt"

/*
Calibration involves determining the number of molecules
that can be generated in one step from a given starting point.


*/

func main() {

	fmt.Println("DEBUG: Answer part1: ", findHouseByPresentsNum(34000000))
}

func findHouseByPresentsNum(presentsNum int) int {
	house := 1
	for {
		presents := 0
		for i := house; i > 0; i-- {
			if house%i == 0 {
				elfPresents := i * 10
				presents += elfPresents
			}
		}
		if presents == presentsNum {
			return house
		}
		fmt.Println("House #:", house)
		fmt.Println("Presents #:", presents)
		house++
	}

	//34000000
}
