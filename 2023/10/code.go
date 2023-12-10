package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

var mappings = map[rune]string{
	'-': "━",
	'|': "┃",
	'7': "┓",
	'L': "┗",
	'J': "┛",
	'F': "┏",
	'S': "S",
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	grid := make([][]rune, 0)
	steps := make([][]int, 0)

	startV := -1
	startH := -1

	for v, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		gl := make([]rune, 0)
		st := make([]int, 0)
		for h := 0; h < len(line); h += 1 {
			gl = append(gl, rune(line[h]))
			st = append(st, -1)

			if rune(line[h]) == 'S' {
				if startV != -1 {
					panic("S already found!")
				}
				startV = v
				startH = h
			}
		}
		grid = append(grid, gl)
		steps = append(steps, st)
	}

	max := maxSteps(grid, steps, startV, startH, -1, -1, 0)

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		// contagion
		// for each cell on the edges that is -1 steps (not in the loop), set to -2

		for v := range steps {
			for h := range steps[v] {
				if steps[v][h] < 0 {
					grid[v][h] = '.'
				}
			}
		}

		// now, expand the grid to make more space to slip "between" the pipes
		bigSteps := make([][]int, 0)
		for v := 0; v < len(steps)*3; v++ {
			bigLine := make([]int, 0)
			for h := 0; h < len(steps[0])*3; h++ {
				bigLine = append(bigLine, -1)
			}
			bigSteps = append(bigSteps, bigLine)
		}

		for v := range grid {
			for h := range grid[v] {
				st := steps[v][h]
				if grid[v][h] == '|' {
					bigSteps = copySteps(bigSteps, v*3, h*3, [][]int{
						{-1, st, -1},
						{-1, st, -1},
						{-1, st, -1},
					})
				} else if grid[v][h] == '-' {
					bigSteps = copySteps(bigSteps, v*3, h*3, [][]int{
						{-1, -1, -1},
						{st, st, st},
						{-1, -1, -1},
					})
				} else if grid[v][h] == '7' {
					bigSteps = copySteps(bigSteps, v*3, h*3, [][]int{
						{-1, -1, -1},
						{st, st, -1},
						{-1, st, -1},
					})
				} else if grid[v][h] == 'F' {
					bigSteps = copySteps(bigSteps, v*3, h*3, [][]int{
						{-1, -1, -1},
						{-1, st, st},
						{-1, st, -1},
					})
				} else if grid[v][h] == 'L' {
					bigSteps = copySteps(bigSteps, v*3, h*3, [][]int{
						{-1, st, -1},
						{-1, st, st},
						{-1, -1, -1},
					})
				} else if grid[v][h] == 'J' {
					bigSteps = copySteps(bigSteps, v*3, h*3, [][]int{
						{-1, st, -1},
						{st, st, -1},
						{-1, -1, -1},
					})
				} else if grid[v][h] == 'S' {
					bigSteps = copySteps(bigSteps, v*3, h*3, [][]int{
						{-1, st, -1},
						{st, st, st},
						{-1, st, -1},
					})
				}
			}
		}

		// Recursively set the outside to "outside" flag (-2)
		for h := 0; h < len(bigSteps[0]); h++ {
			setOutside(bigSteps, 0, h)
			setOutside(bigSteps, len(bigSteps)-1, h)
		}
		for v := 0; v < len(bigSteps); v++ {
			setOutside(bigSteps, v, 0)
			setOutside(bigSteps, v, len(bigSteps[v])-1)
		}

		// Finally, reset these mappings to erase odd space around pipes for final calculation
		for v := range grid {
			for h := range grid[v] {
				st := steps[v][h]
				if grid[v][h] == '|' {
					bigSteps = copySteps(bigSteps, v*3, h*3, [][]int{
						{-3, st, -3},
						{-3, st, -3},
						{-3, st, -3},
					})
				} else if grid[v][h] == '-' {
					bigSteps = copySteps(bigSteps, v*3, h*3, [][]int{
						{-3, -3, -3},
						{st, st, st},
						{-3, -3, -3},
					})
				} else if grid[v][h] == '7' {
					bigSteps = copySteps(bigSteps, v*3, h*3, [][]int{
						{-3, -3, -3},
						{st, st, -3},
						{-3, st, -3},
					})
				} else if grid[v][h] == 'F' {
					bigSteps = copySteps(bigSteps, v*3, h*3, [][]int{
						{-3, -3, -3},
						{-3, st, st},
						{-3, st, -3},
					})
				} else if grid[v][h] == 'L' {
					bigSteps = copySteps(bigSteps, v*3, h*3, [][]int{
						{-3, st, -3},
						{-3, st, st},
						{-3, -3, -3},
					})
				} else if grid[v][h] == 'J' {
					bigSteps = copySteps(bigSteps, v*3, h*3, [][]int{
						{-3, st, -3},
						{st, st, -3},
						{-3, -3, -3},
					})
				} else if grid[v][h] == 'S' {
					bigSteps = copySteps(bigSteps, v*3, h*3, [][]int{
						{-3, st, -3},
						{st, st, st},
						{-3, st, -3},
					})
				}
			}
		}

		tilesContained := 0
		for v := range bigSteps {
			for h := range bigSteps[v] {
				if bigSteps[v][h] == -1 {
					tilesContained += 1
				}
			}
		}

		return tilesContained / 9
	}

	// solve part 1 here
	return max / 2
}

func copySteps(bigSteps [][]int, v int, h int, copy [][]int) [][]int {
	for i := 0; i < len(copy); i++ {
		for j := 0; j < len(copy[i]); j++ {
			bigSteps[v+i][h+j] = copy[i][j]
		}
	}

	return bigSteps
}

func setOutside(steps [][]int, v int, h int) {
	if v < 0 || h < 0 || v >= len(steps) || h >= len(steps[v]) || steps[v][h] != -1 {
		return
	}

	steps[v][h] = -2

	setOutside(steps, v-1, h)
	setOutside(steps, v+1, h)
	setOutside(steps, v, h-1)
	setOutside(steps, v, h+1)
}

func isN(grid [][]rune, steps [][]int, v int, h int) bool {
	if v < 0 || h < 0 || v >= len(grid) || h >= len(grid[0]) || steps[v][h] > 0 {
		return false
	}

	d := grid[v][h]
	return d == '|' || d == '7' || d == 'F' || d == 'S'
}

func isS(grid [][]rune, steps [][]int, v int, h int) bool {
	if v < 0 || h < 0 || v >= len(grid) || h >= len(grid[0]) || steps[v][h] > 0 {
		return false
	}

	d := grid[v][h]
	return d == '|' || d == 'J' || d == 'L' || d == 'S'
}

func isE(grid [][]rune, steps [][]int, v int, h int) bool {
	if v < 0 || h < 0 || v >= len(grid) || h >= len(grid[0]) || steps[v][h] > 0 {
		return false
	}

	d := grid[v][h]
	return d == '-' || d == '7' || d == 'J' || d == 'S'
}

func isW(grid [][]rune, steps [][]int, v int, h int) bool {
	if v < 0 || h < 0 || v >= len(grid) || h >= len(grid[0]) || steps[v][h] > 0 {
		return false
	}

	d := grid[v][h]
	return d == '-' || d == 'F' || d == 'L' || d == 'S'
}

func maxSteps(grid [][]rune, steps [][]int, v int, h int, vFrom int, hFrom int, stepsTaken int) int {

	if v < 0 || h < 0 || v >= len(grid) || h > len(grid[0]) {
		return -1
	}

	current := grid[v][h]

	if stepsTaken > 0 && current == 'S' {
		return stepsTaken
	}

	//fmt.Printf("%v -> \n", string(current))
	steps[v][h] = stepsTaken

	n := false
	s := false
	e := false
	w := false

	switch current {
	case 'S':
		{
			// check all 4 coordinates
			n = true
			s = true
			e = true
			w = true
		}
	case '|':
		{
			// Check N and S
			n = true
			s = true
		}
	case '-':
		{
			// Check E and W
			e = true
			w = true
		}
	case 'L':
		{
			// Check N and E
			n = true
			e = true

		}
	case 'J':
		{
			// Check N and W
			n = true
			w = true
		}
	case '7':
		{
			// Check S and W
			s = true
			w = true
		}
	case 'F':
		{
			// Check S and E
			s = true
			e = true
		}
	}

	if n && isN(grid, steps, v-1, h) && vFrom != v-1 {
		return maxSteps(grid, steps, v-1, h, v, h, stepsTaken+1)
	}

	if s && isS(grid, steps, v+1, h) && vFrom != v+1 {
		return maxSteps(grid, steps, v+1, h, v, h, stepsTaken+1)
	}

	if e && isE(grid, steps, v, h+1) && hFrom != h+1 {
		return maxSteps(grid, steps, v, h+1, v, h, stepsTaken+1)
	}

	if w && isW(grid, steps, v, h-1) && hFrom != h-1 {
		return maxSteps(grid, steps, v, h-1, v, h, stepsTaken+1)
	}

	panic("Nowhere to go?")
}
