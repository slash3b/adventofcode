package main

import (
	"fmt"
)

func main() {

	//var houses []int
	minHouse := 1<<63 - 1
	house := 1
	for {
		presents := getFactorsSum(house)
		//fmt.Println("going through house and presents", house, presents)
		if presents > 35000000 {
			break
		}

		if presents >= 34000000 && presents <= 35000000 {
			if house < minHouse {
				minHouse = house
			}
		}
		house++
	}

	fmt.Println(minHouse)

	//sort.Slice(house, func(i, j int) bool {
	//	return i < j
	//})
	//
	//fmt.Println(houses[0], houses[len(houses)-1])

	//
	//house := 1000000
	//minHouse := 1<<63 - 1
	//
	//var m sync.Mutex
	//
	//totalPresents := 0
	//
	//answers := make(chan [2]int)
	//
	//for {
	//	select {
	//	case ans := <-answers:
	//		ps := ans[0]
	//
	//		if ps >= 34000000 && ps <= 35000000 {
	//			if ans[1] < minHouse {
	//				minHouse = ans[1]
	//			}
	//		}
	//		m.Lock()
	//		if ps > totalPresents {
	//			totalPresents = ps
	//		}
	//		m.Unlock()
	//		if totalPresents > 35000000 {
	//			break
	//		}
	//
	//		fmt.Println("Presents #:", ans[0], "House #:", ans[1], "Min House with 34M+:", minHouse)
	//
	//	default:
	//		go func(house int) {
	//			presents := getFactorsSum(house)
	//			answers <- [2]int{presents, house}
	//		}(house)
	//
	//		house++
	//	}
	//}
}

// kind of like prime factors
func getFactorsSum(houseNum int) int {
	out := make(map[int]bool)
	out[1] = true
	out[houseNum] = true

	if houseNum < 2 {
		return 10
	}

	num := houseNum
	i := 2
	for {
		// continue to divide until we find uneven number
		if num%i != 0 {
			i++
			continue
		}

		if houseNum%num == 0 {
			out[num] = true
		}

		out[i] = true

		if num == i {
			out[i] = true
			break
		}

		num = num / i
	}

	//fmt.Println(out)

	sum := 0
	for k, _ := range out {
		sum += k * 10
	}

	return sum
}
