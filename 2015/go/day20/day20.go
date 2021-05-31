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

	answers := make(chan [2]int)

	for {
		select {
		case ans := <-answers:
			if ans[0] == presentsNum {
				fmt.Println("ANSWER: ", ans[1])
				return ans[1]
			}
			fmt.Println("Presents #:", ans[0], "House #:", ans[1])

		default:
			go func(house int) {
				presents := 0
				for i := house; i > 0; i-- {
					if house%i == 0 {
						elfPresents := i * 10
						presents += elfPresents
					}
				}
				answers <- [2]int{presents, house}
			}(house)

			house++
		}

	}

	//34000000
}
