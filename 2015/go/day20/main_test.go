package main

import (
	"testing"
)

func TestFindHouseByPresents(t *testing.T) {
	cases := []struct {
		Num         int
		ExpectedSum int
	}{
		{
			1,
			1,
		},
		{
			2,
			3, // 1, 2
		},
		{
			3,
			4, // 1, 3
		},
		{
			4,
			7, // 1, 2, 4
		},
		{
			5,
			6, // 1, 5
		},
		{
			6,
			12, // 1, 2, 3, 6
		},
		{
			7,
			8, // 1, 7
		},
		{
			8,
			15, //
		},
		{
			9,
			13, // 1, 3, 9
		},
		{
			10,
			18, // 1, 2, 5, 10
		},
		{
			11,
			12, // 1, 11
		},
		{
			12,
			28, // 1, 2, 3, 4, 6, 12  |||| 1, 2, 3, 4, 6, 12
		},
		{
			14,
			24, // 1, 2, 7, 14
		},
		{
			15,
			24, // 1, 3, 5, 15
		},
	}

	for _, testCase := range cases {
		fSum := GetFactorsSum(testCase.Num)
		if testCase.ExpectedSum != fSum {
			t.Errorf("incorrect result for house=%d; expected=%d actual=%d", testCase.Num, testCase.ExpectedSum, fSum)
		}
	}
}

func BenchmarkGetFactorsSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFactorsSum(1000000)
	}
}
