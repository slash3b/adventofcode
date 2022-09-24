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

	for i := 1; i < len(in)-2; i++ {

		prev := in[i-1] + in[i] + in[i+1]
		curr := in[i] + in[i+1] + in[i+2]

		if curr > prev {
			res++
		}
	}

	return res
}
