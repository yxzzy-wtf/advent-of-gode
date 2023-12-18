package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		// More complex
		return "not implemented"
	}
	// solve part 1 here

	// Get max dimensions (sumR + sumD)

	// or just make a big ass grid

	maxV := 1200
	maxH := 1200

	grid := make([][]string, 0)
	for v := 0; v < maxV; v++ {
		gridLine := make([]string, 0)
		for h := 0; h < maxH; h++ {
			gridLine = append(gridLine, ".")
		}
		grid = append(grid, gridLine)
	}

	v := maxV / 2
	h := maxH / 2

	// Dig out trench
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}

		spl := strings.Split(l, " ")

		order := spl[0]
		dist, _ := strconv.Atoi(spl[1])
		colour := spl[2]

		for i := 0; i < dist; i++ {
			if order == "R" {
				h += 1
			} else if order == "L" {
				h -= 1
			} else if order == "U" {
				v -= 1
			} else if order == "D" {
				v += 1
			} else {
				panic("bad direction")
			}
			grid[v][h] = colour
		}
	}

	// Fill out outside
	outside := [][]int{{0, 0}}
	for len(outside) > 0 {
		o := outside[0]
		outside = outside[1:]

		if grid[o[0]][o[1]] == "." {
			grid[o[0]][o[1]] = " "

			v = o[0]
			h = o[1]

			if v > 0 {
				outside = append(outside, []int{v - 1, h})
			}

			if h > 0 {
				outside = append(outside, []int{v, h - 1})
			}

			if v < len(grid)-1 {
				outside = append(outside, []int{v + 1, h})
			}

			if h < len(grid[0])-1 {
				outside = append(outside, []int{v, h + 1})
			}
		}
	}

	// Finally count fill-in
	digout := 0
	for v := range grid {
		for h := range grid {
			if grid[v][h] != " " {
				digout += 1
			}
		}
	}

	print := false
	if print {
		for v := range grid {
			for h := range grid[v] {
				if len(grid[v][h]) == 1 {
					fmt.Print(grid[v][h])
				} else {
					fmt.Print("#")
				}
			}
			fmt.Println()
		}
	}

	return digout
}
