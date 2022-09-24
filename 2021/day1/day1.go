package day1

func Run(in []int) int {
	c := 0

	for i := 1; i < len(in); i++ {
		if in[i] > in[i-1] {
			c++
		}
	}

	return c
}

func Run2(in []int) int {

	res := 0

	wind := []int{}
	sum := 0

	// 1,2,3,4,5,6
	//       ^ this is where we start, from the right most border of a 3 int window

	prevSum := 0
	for i, v := range in {
		wind = append(wind, v)
		sum += v

		if i < 3 {
			continue
		}

		// ready to do calculation stage

		sum -= wind[0]
		wind = wind[1:]

		if sum > prevSum {
			res++
		}

		prevSum = sum

	}

	return res
}
