package main

import (
	"adventofcode/util"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var test bool
var p int

func init() {
	flag.BoolVar(&test, "test", false, "run with test data. Just pass `-test` to the run command to use test data")
	flag.IntVar(&p, "p", 1, "part to run, 1 by default")
}

func main() {
	flag.Parse()

	inputType := "puzzle"
	if test {
		inputType = "sample"
	}

	lines := util.FileToStringSlice(inputType)

	stacks := make([]*List, 0)
	commands := make([]string, 0)

	makeStacks := func(input []string) []*List {
		// figure indexes
		idxs := []int{}
		for i, v := range input[len(input)-1] {
			if v != ' ' {
				idxs = append(idxs, i)
			}
		}

		// discard last item
		input = input[:len(input)-1]

		// make a slice of lists
		lists := make([]*List, len(idxs))
		for i := 0; i < len(idxs); i++ {
			lists[i] = &List{}
		}

		for i := len(input) - 1; i >= 0; i-- {
			for j, v := range idxs {
				if input[i][v] != ' ' {
					lists[j].Append(&Node{Val: rune(input[i][v])})
				}
			}
		}

		return lists
	}

	for i, v := range lines {
		if v == "" {
			stacks = makeStacks(lines[:i])
			commands = lines[i+1:]
			break
		}
	}

	switch p {
	case 1:
		fmt.Println(Part1(stacks, commands))
	case 2:
		fmt.Println(Part2(stacks, commands))
	default:
		panic("wrong part number given")
	}
}

type Node struct {
	Val  rune
	Prev *Node
}

type List struct {
	Tail *Node
}

func (l *List) Append(n *Node) {
	if l.Tail == nil {
		l.Tail = n
		l.Tail.Prev = nil
		return
	}

	prevTail := l.Tail
	l.Tail = n

	n.Prev = prevTail
}

func (l *List) Move(to *List) {
	curr := l.Tail
	if l.Tail.Prev == nil {
		l.Tail = nil
	} else {
		l.Tail = l.Tail.Prev
	}

	to.Append(curr)
}

func (l *List) MoveMany(n int, to *List) {

	//find node we need to detach
	oldTail := l.Tail
	curr := l.Tail
	for i := 1; i < n; i++ {
		curr = curr.Prev
	}

	//fmt.Println("old tail:", string(oldTail.Val))
	// fmt.Println("curr node:", string(curr.Val))

	l.Tail = curr.Prev

	curr.Prev = to.Tail
	to.Tail = oldTail
}

func debug(title string, stacks []*List) {
	fmt.Printf("%s:-------------\n", strings.Title(title))
	for i := range stacks {
		fmt.Print(i, ": ")
		fmt.Print("\t")
		curr := stacks[i].Tail
		for curr != nil {
			fmt.Printf("%q", string(curr.Val))
			if curr.Prev != nil {
				fmt.Printf("-> %q", string(curr.Prev.Val))
			} else {
				fmt.Print("-> nil")
			}
			fmt.Printf(" | ")
			if curr == curr.Prev {
				return
			}

			curr = curr.Prev

		}
		fmt.Println("")
	}
}

func Part1(stacks []*List, commands []string) int {
	debug("init", stacks)

	fmt.Println("-----------------------")

	reg := regexp.MustCompile(`\d+`)
	for _, v := range commands {
		fmt.Println("command", v)
		res := reg.FindAllString(v, -1)
		move, _ := strconv.Atoi(res[0])
		from, _ := strconv.Atoi(res[1])
		to, _ := strconv.Atoi(res[2])

		fromList := stacks[from-1]
		toList := stacks[to-1]

		for move > 0 {
			fromList.Move(toList)
			move--
		}
		debug("____", stacks)
	}

	for _, v := range stacks {
		fmt.Print(string(v.Tail.Val))
	}
	fmt.Println("")

	return 0
}

func Part2(stacks []*List, commands []string) int {
	debug("init", stacks)

	fmt.Println("-----------------------")

	reg := regexp.MustCompile(`\d+`)
	for _, v := range commands {
		res := reg.FindAllString(v, -1)
		move, _ := strconv.Atoi(res[0])
		from, _ := strconv.Atoi(res[1])
		to, _ := strconv.Atoi(res[2])

		fromList := stacks[from-1]
		toList := stacks[to-1]

		fromList.MoveMany(move, toList)
	}

	for _, v := range stacks {
		fmt.Print(string(v.Tail.Val))
	}
	fmt.Println("")

	return 0
}
