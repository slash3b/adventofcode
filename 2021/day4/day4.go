package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"adventofcode/util"
)

var partNum int
var inputType string

func init() {
	flag.IntVar(&partNum, "part", 1, "part to run, can be either 1 or 2")
	flag.StringVar(&inputType, "input", "sample", "input type to use, either sample or puzzle. It is 'sample' by default")
}

type Board struct {
	isWon   bool
	b       [][]int
	visited map[int]struct{}
	sum     int
}

func (b *Board) Receive(v int) int {
	b.visited[v] = struct{}{}

	for _, row := range b.b {
		found := true

		for _, val := range row {
			_, ok := b.visited[val]
			if !ok {
				found = false
				break
			}
		}

		if found {
			s := 0
			for _, row := range b.b {
				for _, val := range row {
					if _, ok := b.visited[val]; !ok {
						s += val
					}
				}
			}
			fmt.Printf("Horizon: %#v \n", row)

			return s * v
		}
	}

	for i := 0; i < len(b.b); i++ {
		found := true
		tmpDebug := []int{}
		for j := 0; j < len(b.b[i]); j++ {
			k := b.b[j][i]
			tmpDebug = append(tmpDebug, k)

			_, ok := b.visited[k]
			if !ok {
				found = false
				break
			}
		}

		if found {
			s := 0
			for _, row := range b.b {
				for _, val := range row {
					if _, ok := b.visited[val]; !ok {
						s += val
					}
				}
			}

			fmt.Printf("Vertical: %#v \n", tmpDebug)
			return s * v
		}
	}

	return -1
}

func (b *Board) ReceivePart2(v int) int {
	if b.isWon {
		return -1
	}

	b.visited[v] = struct{}{}

	for _, row := range b.b {
		found := true

		for _, val := range row {
			_, ok := b.visited[val]
			if !ok {
				found = false
				break
			}
		}

		if found {
			s := 0
			for _, row := range b.b {
				for _, val := range row {
					if _, ok := b.visited[val]; !ok {
						s += val
					}
				}
			}
			fmt.Printf("Horizon: %#v \n", row)
			b.isWon = true
			return s * v
		}
	}

	for i := 0; i < len(b.b); i++ {
		found := true
		tmpDebug := []int{}
		for j := 0; j < len(b.b[i]); j++ {
			k := b.b[j][i]
			tmpDebug = append(tmpDebug, k)

			_, ok := b.visited[k]
			if !ok {
				found = false
				break
			}
		}

		if found {
			s := 0
			for _, row := range b.b {
				for _, val := range row {
					if _, ok := b.visited[val]; !ok {
						s += val
					}
				}
			}

			b.isWon = true
			fmt.Printf("Vertical: %#v \n", tmpDebug)
			return s * v
		}
	}

	return -1
}

func main() {
	flag.Parse()

	data := util.FileContents(inputType)
	plainStrs := strings.Split(data, "\n\n")

	numInput := []int{}
	for _, v := range strings.Split(plainStrs[0], ",") {
		i, _ := strconv.Atoi(v)
		numInput = append(numInput, i)
	}

	boards := make([]*Board, 0)

	for _, v := range plainStrs[1:] {
		board := [][]int{}
		sum := 0
		for _, line := range strings.Split(v, "\n") {

			b := []int{}
			for _, iv := range strings.Split(line, " ") {
				i, err := strconv.Atoi(iv)
				// fixme: figure why there is an error here ?
				if err != nil {
					continue
				}
				sum += i
				b = append(b, i)
			}
			if len(b) == 0 {
				continue
			}
			board = append(board, b)

		}

		boards = append(boards, &Board{
			b:       board,
			sum:     sum,
			visited: make(map[int]struct{}),
		})
	}

	switch partNum {
	case 1:
		Part1(numInput, boards)
	case 2:
		Part2(numInput, boards)
	default:
		panic("wrong part number given")
	}
}

func Part1(in []int, boards []*Board) {
	for _, v := range in {
		fmt.Println("sending V to boards", v)
		for _, board := range boards {
			res := board.Receive(v)
			if res > 0 {
				fmt.Println(">>>>", res)
				return
			}
		}
	}

	fmt.Println("failed to find")
}

func Part2(in []int, boards []*Board) {
	for _, v := range in {
		fmt.Println("sending V to boards", v)
		for _, board := range boards {
			res := board.ReceivePart2(v)
			if res > 0 {
				fmt.Println(">>>>", res)
			}
		}
	}

}
