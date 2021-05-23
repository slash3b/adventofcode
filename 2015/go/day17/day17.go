package main

import "fmt"

func main() {
	containers := []int{
		33,
		14,
		18,
		20,
		45,
		35,
		16,
		35,
		1,
		13,
		18,
		13,
		50,
		44,
		48,
		6,
		24,
		41,
		30,
		42,
	}

	powerSetResults := PowerSetWrapper(containers)

	var counter int
	minContainers := make(map[int]int)

	for _, set := range powerSetResults {
		sum := 0
		for _, value := range set {
			sum += value
		}
		if sum == 150 {
			minContainers[len(set)]++
			fmt.Println(set)
			counter++
		}
	}

	// now lets find the freaking minimum of containers and a number of ways we can fill in
	theMostMin := 1<<63 - 1
	for k, _ := range minContainers {
		if k < theMostMin {
			theMostMin = k
		}
	}

	fmt.Println("RESULTS: ", len(powerSetResults))
	fmt.Println("counter: ", counter)
	fmt.Println("the most min ways: ", minContainers[theMostMin])

	//readAndPrepare("/home/slash3b/Projects/aoc/2015/go/day17/input")
}

func readAndPrepare(s string) {
	//res, err := ioutil.ReadFile(s)
	//if err != nil {
	//	panic(err)
	//}
	//
	//inputs := strings.Split(string(res), "\n")

	//for _, v := range inputs {
	//}
}
