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
	grid := make([][]rune, 0)
	v, h := 0, 0
	for y, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}

		gridLine := make([]rune, 0)
		for x, s := range strings.Split(l, "") {
			if s == "S" {
				v, h = y, x
			}
			gridLine = append(gridLine, rune(s[0]))
		}
		grid = append(grid, gridLine)
	}

	steps := 64

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here

	possibleCells := [][]int{{v, h}}

	for i := 0; i < steps; i++ {
		nextStepCells := make([][]int, 0)
		alreadyStepped := make(map[string]bool)

		for _, pc := range possibleCells {

			if pc[0] > 0 && grid[pc[0]-1][pc[1]] != '#' {
				if _, ok := alreadyStepped[fmt.Sprintf("%v;%v", pc[0]-1, pc[1])]; !ok {
					nextStepCells = append(nextStepCells, []int{pc[0] - 1, pc[1]})
					alreadyStepped[fmt.Sprintf("%v;%v", pc[0]-1, pc[1])] = true
				}
			}

			if pc[1] > 0 && grid[pc[0]][pc[1]-1] != '#' {
				if _, ok := alreadyStepped[fmt.Sprintf("%v;%v", pc[0], pc[1]-1)]; !ok {
					nextStepCells = append(nextStepCells, []int{pc[0], pc[1] - 1})
					alreadyStepped[fmt.Sprintf("%v;%v", pc[0], pc[1]-1)] = true
				}
			}

			if pc[0] < len(grid)-1 && grid[pc[0]+1][pc[1]] != '#' {
				if _, ok := alreadyStepped[fmt.Sprintf("%v;%v", pc[0]+1, pc[1])]; !ok {
					nextStepCells = append(nextStepCells, []int{pc[0] + 1, pc[1]})
					alreadyStepped[fmt.Sprintf("%v;%v", pc[0]+1, pc[1])] = true
				}
			}

			if pc[1] < len(grid[0])-1 && grid[pc[0]][pc[1]+1] != '#' {
				if _, ok := alreadyStepped[fmt.Sprintf("%v;%v", pc[0], pc[1]+1)]; !ok {
					nextStepCells = append(nextStepCells, []int{pc[0], pc[1] + 1})
					alreadyStepped[fmt.Sprintf("%v;%v", pc[0], pc[1]+1)] = true
				}
			}

		}

		possibleCells = nextStepCells
		//fmt.Printf("Possible steps: %v\n", possibleCells)
	}

	return len(possibleCells)
}
