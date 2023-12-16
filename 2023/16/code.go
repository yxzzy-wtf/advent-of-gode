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
type BeamSect struct {
	Direction rune
	Y         int
	X         int
}

type Beam struct {
	Direction      rune
	Path           []BeamSect
	EnergizedTiles int
}

type Tile struct {
	Type   rune
	ShineN bool
	ShineE bool
	ShineS bool
	ShineW bool
	X      int
	Y      int
}

func calcEnergized(grid [][]Tile, start Tile) int {

	tilesChanged := []Tile{start}

	for len(tilesChanged) > 0 {
		// Contagion
		stillChanged := make([]Tile, 0)
		for _, t := range tilesChanged {
			if t.ShineE && t.X < len(grid[t.Y])-1 {
				eX := t.X + 1
				eY := t.Y

				shineInto := &grid[eY][eX]

				if (shineInto.Type == '.' || shineInto.Type == '-') && !shineInto.ShineE {
					// Change!
					shineInto.ShineE = true
					stillChanged = append(stillChanged, *shineInto)
				} else if shineInto.Type == '|' && (!shineInto.ShineN || !shineInto.ShineS) {
					shineInto.ShineN = true
					shineInto.ShineS = true
					stillChanged = append(stillChanged, *shineInto)
				} else if shineInto.Type == '/' && !shineInto.ShineN {
					shineInto.ShineN = true
					stillChanged = append(stillChanged, *shineInto)
				} else if shineInto.Type == '\\' && !shineInto.ShineS {
					shineInto.ShineS = true
					stillChanged = append(stillChanged, *shineInto)
				}
			}

			if t.ShineW && t.X > 0 {
				wX := t.X - 1
				wY := t.Y

				shineInto := &grid[wY][wX]

				if (shineInto.Type == '.' || shineInto.Type == '-') && !shineInto.ShineW {
					// Change!
					shineInto.ShineW = true
					stillChanged = append(stillChanged, *shineInto)
				} else if shineInto.Type == '|' && (!shineInto.ShineN || !shineInto.ShineS) {
					shineInto.ShineN = true
					shineInto.ShineS = true
					stillChanged = append(stillChanged, *shineInto)
				} else if shineInto.Type == '/' && !shineInto.ShineS {
					shineInto.ShineS = true
					stillChanged = append(stillChanged, *shineInto)
				} else if shineInto.Type == '\\' && !shineInto.ShineN {
					shineInto.ShineN = true
					stillChanged = append(stillChanged, *shineInto)
				}
			}

			if t.ShineN && t.Y > 0 {
				nX := t.X
				nY := t.Y - 1

				shineInto := &grid[nY][nX]

				if (shineInto.Type == '.' || shineInto.Type == '|') && !shineInto.ShineN {
					// Change!
					shineInto.ShineN = true
					stillChanged = append(stillChanged, *shineInto)
				} else if shineInto.Type == '-' && (!shineInto.ShineE || !shineInto.ShineW) {
					shineInto.ShineE = true
					shineInto.ShineW = true
					stillChanged = append(stillChanged, *shineInto)
				} else if shineInto.Type == '/' && !shineInto.ShineE {
					shineInto.ShineE = true
					stillChanged = append(stillChanged, *shineInto)
				} else if shineInto.Type == '\\' && !shineInto.ShineW {
					shineInto.ShineW = true
					stillChanged = append(stillChanged, *shineInto)
				}

			}

			if t.ShineS && t.Y < len(grid)-1 {
				sX := t.X
				sY := t.Y + 1

				shineInto := &grid[sY][sX]

				if (shineInto.Type == '.' || shineInto.Type == '|') && !shineInto.ShineS {
					// Change!
					shineInto.ShineS = true
					stillChanged = append(stillChanged, *shineInto)
				} else if shineInto.Type == '-' && (!shineInto.ShineE || !shineInto.ShineW) {
					shineInto.ShineE = true
					shineInto.ShineW = true
					stillChanged = append(stillChanged, *shineInto)
				} else if shineInto.Type == '/' && !shineInto.ShineW {
					shineInto.ShineW = true
					stillChanged = append(stillChanged, *shineInto)
				} else if shineInto.Type == '\\' && !shineInto.ShineE {
					shineInto.ShineE = true
					stillChanged = append(stillChanged, *shineInto)
				}

			}
		}
		tilesChanged = stillChanged
	}

	energized := 0
	for v := range grid {
		for h := range grid[v] {
			if grid[v][h].ShineE || grid[v][h].ShineN || grid[v][h].ShineS || grid[v][h].ShineW {
				energized += 1
			}
		}
	}

	return energized
}

func resetGrid(grid [][]Tile) {
	for v := range grid {
		for h := range grid[v] {
			grid[h][v].ShineE = false
			grid[h][v].ShineN = false
			grid[h][v].ShineS = false
			grid[h][v].ShineW = false
		}
	}
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	grid := make([][]Tile, 0)
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}

		gridLine := make([]Tile, 0)
		for r := range l {
			gridLine = append(gridLine, Tile{Type: rune(l[r]), X: len(gridLine), Y: len(grid)})
		}
		grid = append(grid, gridLine)
	}

	if part2 {
		maxEnergized := -1

		for v := range grid {
			testE := calcEnergized(grid, Tile{ShineE: true, X: -1, Y: v})
			resetGrid(grid)
			if testE > maxEnergized {
				maxEnergized = testE
			}

			testW := calcEnergized(grid, Tile{ShineW: true, X: len(grid[v]), Y: v})
			resetGrid(grid)
			if testW > maxEnergized {
				maxEnergized = testW
			}
		}

		for h := range grid[0] {
			testN := calcEnergized(grid, Tile{ShineN: true, X: h, Y: len(grid)})
			resetGrid(grid)
			if testN > maxEnergized {
				maxEnergized = testN
			}

			testS := calcEnergized(grid, Tile{ShineS: true, X: h, Y: -1})
			resetGrid(grid)
			if testS > maxEnergized {
				maxEnergized = testS
			}
		}

		return maxEnergized
	}
	// solve part 1 her

	return calcEnergized(grid, Tile{ShineE: true, X: -1, Y: 0})
}
