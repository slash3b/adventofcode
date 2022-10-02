package main

import (
	"adventofcode/util"
	"flag"
	"fmt"
)

var partNum int
var inputType string

func init() {
	flag.IntVar(&partNum, "part", 1, "part to run, can be either 1 or 2")
	flag.StringVar(&inputType, "input", "sample", "input type to use, either sample or puzzle. It is 'sample' by default")
}

func main() {
	flag.Parse()

	data := util.FileToStringSlice(inputType)

	switch partNum {
	case 1:
		Part1(data)
	case 2:
		Part2(data)
	default:
		panic("wrong part number given")
	}
}

func Part1(in []string) {
	pop := make([]int, len(in[0]))

	for _, v := range in {
		for i := 0; i < len(v); i++ {
			if v[i] == '1' {
				pop[i]++
				continue
			}
			pop[i]--
		}
	}

	var gamma int

	ln := len(pop)
	for i := 0; i < ln; i++ {
		if pop[i] > 0 {
			gamma += 1 << (ln - i - 1)
			continue
		}
	}

	mask := 1<<len(pop) - 1
	epsilon := gamma ^ mask
	result := gamma * epsilon
	fmt.Printf("gamma:\t\t%012b %[1]d \n", gamma)
	fmt.Printf("mask:\t\t%012b %[1]d\n", mask)
	fmt.Printf("epsilon:\t%012b %[1]d\n", epsilon)
	fmt.Printf("result:\t\t%012b %[1]d\n", result)
}

func Part2(in []string) {

	oxy := in[:]

	filter := func(num int, col int, s []string) []string {
		res := []string{}

		for i := 0; i < len(s); i++ {

			if num == 0 && s[i][col] == '1' {
				res = append(res, s[i])
				continue
			}

			if num > 0 && s[i][col] == '1' {
				res = append(res, s[i])
				continue
			}

			if num < 0 && s[i][col] == '0' {
				res = append(res, s[i])
				continue
			}
		}

		return res
	}

	col := 0
	for len(oxy) > 1 {
		pop := make([]int, len(oxy[0]))
		for _, v := range oxy {
			for i := 0; i < len(v); i++ {
				if v[i] == '1' {
					pop[i]++
					continue
				}
				pop[i]--
			}
		}
		val := pop[col]
		oxy = filter(val, col, oxy)
		col++
	}

	oxyDec := 0
	for k, v := range oxy[0] {
		if v == '1' {
			oxyDec += 1 << (len(oxy[0]) - 1 - k)
		}
	}

	fmt.Println("oxygen value: ", oxyDec)

	filter2 := func(num int, col int, s []string) []string {
		res := []string{}

		for i := 0; i < len(s); i++ {

			if num == 0 && s[i][col] == '0' {
				res = append(res, s[i])
				continue
			}

			if num > 0 && s[i][col] == '0' {
				res = append(res, s[i])
				continue
			}

			if num < 0 && s[i][col] == '1' {
				res = append(res, s[i])
				continue
			}
		}

		return res
	}

	co2 := in[:]

	col = 0
	for len(co2) > 1 {
		pop := make([]int, len(co2[0]))
		for _, v := range co2 {
			for i := 0; i < len(v); i++ {
				if v[i] == '1' {
					pop[i]++
					continue
				}
				pop[i]--
			}
		}
		val := pop[col]
		co2 = filter2(val, col, co2)
		col++
	}

	co2Dec := 0
	for k, v := range co2[0] {
		if v == '1' {
			co2Dec += 1 << (len(co2[0]) - 1 - k)
		}
	}
	fmt.Println("CO2 value:", co2Dec)

	fmt.Println("FINAL", oxyDec*co2Dec)
}
