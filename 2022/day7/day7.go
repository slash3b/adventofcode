package main

import (
	"adventofcode/util"
	"flag"
	"fmt"
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

	root := &Dir{Name: "/"}
	current := root
	lines = lines[1:]
	for _, line := range lines {
		if line[0] == '$' {
			line = line[2:]
		}

		if line == "ls" {
			continue
		}

		if line == "cd .." {
			current = current.Parent
			continue
		}

		res := strings.Split(line, " ")

		// cd <dir
		// size file

		if len(res) == 2 {
			if res[0] == "dir" {
				current.Dirs = append(current.Dirs, &Dir{
					Name:   res[1],
					Parent: current,
				})
				continue
			}

			if res[0] == "cd" { // that means go to specific directory
				for _, d := range current.Dirs {
					if d.Name == res[1] {
						current = d
						break
					}
				}
				continue
			}

			// and that leaves us with file
			bytesSize, err := strconv.Atoi(res[0])
			if err != nil {
				panic(err)
			}

			current.Files = append(current.Files, &File{Name: res[1], Size: bytesSize})
		}
	}
	withSize(root)

	switch p {
	case 1:
		fmt.Println(Part1(root))
	case 2:
		fmt.Println(Part2(root))
	default:
		panic("wrong part number given")
	}
}

func withSize(dr *Dir) {
	size := 0
	for _, d := range dr.Dirs {
		withSize(d)
	}

	for _, f := range dr.Files {
		size += f.Size
	}
	for _, d := range dr.Dirs {
		size += d.Size
	}

	dr.Size = size
}

type Dir struct {
	Name   string
	Parent *Dir
	Dirs   []*Dir
	Files  []*File
	Size   int
}

type File struct {
	Name string
	Size int
}

func part1(r *Dir) int {
	size := 0

	for _, d := range r.Dirs {
		size += part1(d)
	}

	if r.Size < 100000 {
		size += r.Size
	}

	return size
}

func Part1(root *Dir) int {
	fmt.Println("RES", part1(root))

	return 0
}

func Part2(lines *Dir) int {
	// 70000000
	// 30000000
	return 0
}

func debug(root *Dir, lvl int) {
	fmt.Printf("%s%s %d\n", mrg(lvl), root.Name, root.Size)
	fmt.Println(mrg(lvl), "Files:")
	fmt.Print(mrg(lvl), " ")
	for _, f := range root.Files {
		fmt.Printf("%s %d, ", f.Name, f.Size)
	}
	fmt.Print("\n")

	if len(root.Dirs) == 0 {
		return
	}
	fmt.Println(mrg(lvl), "Dirs")
	for _, d := range root.Dirs {
		debug(d, lvl+2)
	}
}

func mrg(n int) string {
	out := ""
	for i := 0; i < n; i++ {
		out += " "
	}

	return out
}
