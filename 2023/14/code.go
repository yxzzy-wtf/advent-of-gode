package main

import (
	"fmt"
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

	grid := make([][]rune, 0)
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}

		gridLine := make([]rune, 0)
		for _, r := range strings.Split(l, "") {
			gridLine = append(gridLine, rune(r[0]))
		}

		grid = append(grid, gridLine)
	}

	if part2 {
		return "not implemented"
	}
	// solve part 1 here

	printGrid(grid)
	grid = tilt(grid, 'N')
	printGrid(grid)

	distToS := len(grid)
	load := 0
	for v := 0; v < len(grid); v++ {
		for h := 0; h < len(grid[v]); h++ {
			if grid[v][h] == 'O' {
				load += (distToS - v)
			}
		}
	}

	return load
}

func printGrid(grid [][]rune) {
	print := false

	if print {
		fmt.Println("-----")
		for v := 0; v < len(grid); v++ {
			for h := 0; h < len(grid[v]); h++ {
				fmt.Printf("%v", string(grid[v][h]))
			}
			fmt.Println()
		}
		fmt.Println("-----")
	}
}

func tilt(grid [][]rune, dir rune) [][]rune {
	if dir == 'N' {
		for v := 0; v < len(grid); v++ {
			for h := 0; h < len(grid[v]); h++ {
				if grid[v][h] == 'O' {
					// Move up as far as possible:
				rolling:
					for mv := v - 1; mv >= 0; mv-- {
						if grid[mv][h] == '.' {
							// swap and continue
							grid[mv][h] = 'O'
							grid[mv+1][h] = '.'
						} else {
							break rolling
						}
					}
				}
			}
		}

		return grid
	}

	panic("Bad direction")
}
