package main

import (
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
		matchMap := make(map[string]int)
		cycleVal := make(map[int]int)
		cycleStart := -1
		cycleEnd := -1
		for i := 1; i <= 1000000000; i++ {
			grid = spin(grid)
			gridStr := gridString(grid)

			if at, ok := matchMap[gridStr]; ok {
				cycleVal[at] = northLoad(grid)
				if cycleStart == -1 {
					cycleStart = at
				} else if at > cycleStart {
					cycleEnd = at
				} else {
					// Found our full cycle, break
					break
				}
			} else {
				matchMap[gridStr] = i
			}
		}

		cycleLen := cycleEnd - cycleStart + 1
		// calculate which cycle will be the 1000000000th
		cycle := cycleStart + (1000000000-cycleStart)%cycleLen

		return cycleVal[cycle]
	}

	grid = tilt(grid, 'N')
	return northLoad(grid)
}

func northLoad(grid [][]rune) int {
	load := 0
	distToS := len(grid)
	for v := 0; v < len(grid); v++ {
		for h := 0; h < len(grid[v]); h++ {
			if grid[v][h] == 'O' {
				load += (distToS - v)
			}
		}
	}

	return load
}

func gridString(grid [][]rune) string {
	s := ""
	for v := 0; v < len(grid); v++ {
		for h := 0; h < len(grid[v]); h++ {
			s += string(grid[v][h])
		}
	}
	return s
}

func spin(grid [][]rune) [][]rune {
	grid = tilt(grid, 'N')
	grid = tilt(grid, 'W')
	grid = tilt(grid, 'S')
	grid = tilt(grid, 'E')
	return grid
}

func tilt(grid [][]rune, dir rune) [][]rune {
	if dir == 'N' {
		for v := 0; v < len(grid); v++ {
			for h := 0; h < len(grid[v]); h++ {
				if grid[v][h] == 'O' {
					// Move up as far as possible:
				rollingN:
					for mv := v - 1; mv >= 0; mv-- {
						if grid[mv][h] == '.' {
							// swap and continue
							grid[mv][h] = 'O'
							grid[mv+1][h] = '.'
						} else {
							break rollingN
						}
					}
				}
			}
		}

		return grid
	} else if dir == 'S' {
		for v := len(grid) - 1; v >= 0; v-- {
			for h := 0; h < len(grid[v]); h++ {
				if grid[v][h] == 'O' {
					// Move down as far as possible:
				rollingS:
					for mv := v + 1; mv < len(grid); mv++ {
						if grid[mv][h] == '.' {
							// swap and continue
							grid[mv][h] = 'O'
							grid[mv-1][h] = '.'
						} else {
							break rollingS
						}
					}
				}
			}
		}

		return grid
	} else if dir == 'E' {
		for h := len(grid[0]) - 1; h >= 0; h-- {
			for v := 0; v < len(grid); v++ {
				if grid[v][h] == 'O' {
					// Move right as far as possible:
				rollingE:
					for mh := h + 1; mh < len(grid); mh++ {
						if grid[v][mh] == '.' {
							// swap and continue
							grid[v][mh] = 'O'
							grid[v][mh-1] = '.'
						} else {
							break rollingE
						}
					}
				}
			}
		}
		return grid
	} else if dir == 'W' {
		for h := 0; h < len(grid[0]); h++ {
			for v := 0; v < len(grid); v++ {
				if grid[v][h] == 'O' {
					// Move right as far as possible:
				rollingW:
					for mh := h - 1; mh >= 0; mh-- {
						if grid[v][mh] == '.' {
							// swap and continue
							grid[v][mh] = 'O'
							grid[v][mh+1] = '.'
						} else {
							break rollingW
						}
					}
				}
			}
		}
		return grid
	}

	panic("Bad direction")
}
