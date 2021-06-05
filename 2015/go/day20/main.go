package main

import "fmt"

func main() {

	////var houses []int
	minHouse := 1<<63 - 1
	house := 100000
	for {
		presents := GetFactorsSum(house)
		//fmt.Println("going through house and presents", house, presents)
		if presents > 3_500_000 {
			break
		}

		if presents >= 3_400_000 && house < minHouse {
			fmt.Println("FOUND min:", house)
			minHouse = house
		}
		house++

		fmt.Println(" house #", house, " presents:", presents, " min house: ", minHouse)
	}

	fmt.Println(">>>>", minHouse)

	/// -----------------------------
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
	//	default:
	//		if totalPresents > 3500000 {
	//			close(answers)
	//			continue
	//		}
	//
	//		go func(house int) {
	//			presents := GetFactorsSum(house)
	//			answers <- [2]int{presents, house}
	//		}(house)
	//
	//		house++
	//	case ans := <-answers:
	//		ps := ans[0]
	//
	//		if ps == 3400000 {
	//			if ans[1] < minHouse {
	//				minHouse = ans[1]
	//			}
	//		}
	//		go func() {
	//			m.Lock()
	//			if ps > totalPresents {
	//				totalPresents = ps
	//			}
	//			m.Unlock()
	//		}()
	//
	//		fmt.Println("Presents #:", ans[0], "House #:", ans[1], "Min House with 34M+:", minHouse)
	//
	//		//default:
	//		// I always get the same number -- 1700864
	//
	//	}
	//}
}

// GetFactorsSum return sum of factors for the number
func GetFactorsSum(num int) int {
	// out is for debug purposes basically
	out := []int{1, num}
	sum := 0

	if num == 1 {
		return 1
	}

	limit := num / 2

	for i := limit; i > 1; i-- {
		if num%i == 0 {
			out = append(out, i)
		}
	}

	for _, v := range out {
		sum += v
	}

	return sum
}
