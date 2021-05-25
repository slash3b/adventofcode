package main

import (
	"testing"
)

func TestIsBorderline(t *testing.T) {

	cases := []struct {
		Input    [][]int
		Y, X     int
		Expected bool
	}{
		{
			Input: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			Y:        0,
			X:        0,
			Expected: true,
		},
		{
			Input: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			Y:        0,
			X:        1,
			Expected: true,
		},
		{
			Input: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			Y:        1,
			X:        0,
			Expected: true,
		},
		{
			Input: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			Y:        1,
			X:        1,
			Expected: false,
		},
		{
			Input: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			Y:        2,
			X:        1,
			Expected: true,
		},
	}

	for _, testCase := range cases {
		if testCase.Expected != isBorderline(testCase.Input, testCase.Y, testCase.X) {
			t.Error("ooops")
		}
	}

}

func TestCalculateStateForNormalCell(t *testing.T) {

	cases := []struct {
		Input    [][]int
		Y, X     int
		Expected int
		Comment  string
	}{
		{
			Input: [][]int{
				{1, 1, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			Y:        1,
			X:        1,
			Expected: 1, //
			Comment:  "2 lights so stays on",
		},
		{
			Input: [][]int{
				{1, 1, 1},
				{0, 1, 0},
				{0, 0, 0},
			},
			Y:        1,
			X:        1,
			Expected: 1, //
			Comment:  "3 lights so stays on",
		},
		{
			Input: [][]int{
				{1, 1, 1},
				{0, 1, 1},
				{0, 0, 0},
			},
			Y:        1,
			X:        1,
			Expected: 0, //
			Comment:  "4 lights so turns off",
		},

		{
			Input: [][]int{
				{1, 1, 1},
				{0, 0, 0},
				{0, 0, 0},
			},
			Y:        1,
			X:        1,
			Expected: 1, //
			Comment:  "turns on since there are 3 light in the neighbourhood",
		},

		{
			Input: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			Y:        1,
			X:        1,
			Expected: 0, //
			Comment:  "stays as is",
		},

		{
			Input: [][]int{
				{1, 1, 1},
				{0, 0, 0},
				{1, 1, 1},
			},
			Y:        1,
			X:        1,
			Expected: 0, //
			Comment:  "stays as is",
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.Comment, func(t *testing.T) {
			if testCase.Expected != lookAtNeighbours(testCase.Input, testCase.Y, testCase.X) {
				t.Error("ooops")
			}
		})
	}
}

func TestCalculateStateForBorderlineCell(t *testing.T) {

	cases := []struct {
		Input    [][]int
		Y, X     int
		Expected int
		Comment  string
	}{
		{
			Input: [][]int{
				{1, 1, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			Y:        1,
			X:        1,
			Expected: 1, //
			Comment:  "2 lights so stays on",
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.Comment, func(t *testing.T) {
			if testCase.Expected != lookAtNeighbours(testCase.Input, testCase.Y, testCase.X) {
				t.Error("ooops")
			}
		})
	}
}
