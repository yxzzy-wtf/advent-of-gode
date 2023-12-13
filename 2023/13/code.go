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
	grids := make([][][]rune, 0)

	thisGrid := make([][]rune, 0)
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			grids = append(grids, thisGrid)
			thisGrid = make([][]rune, 0)
			continue
		}

		thisGridLine := make([]rune, 0)
		for _, r := range strings.Split(l, "") {
			thisGridLine = append(thisGridLine, rune(r[0]))
		}

		thisGrid = append(thisGrid, thisGridLine)
	}

	if part2 {
		sum := 0
		for _, g := range grids {
			cv, ch := calcReflectionRemainder(g, -1, -1)
			found := false
		calcGrid:
			for v := range g {
				for h := range g[v] {
					//fmt.Printf("Smudging [%v, %v]...", v, h)
					if g[v][h] == '.' {
						// try fix the smudge!
						g[v][h] = '#'

						nv, nh := calcReflectionRemainder(g, cv, ch)
						newVal := -1
						if nv != -1 {
							newVal = (nv + 1) * 100
						} else if nh != -1 {
							newVal = nh + 1
						}
						//fmt.Printf(" (%v)\n", newVal)

						if newVal != -1 {
							// We've found it!
							sum += newVal
							found = true
							break calcGrid
						}

						g[v][h] = '.'
					} else {
						// try fix the smudge!
						g[v][h] = '.'

						nv, nh := calcReflectionRemainder(g, cv, ch)
						newVal := -1
						if nv != -1 {
							newVal = (nv + 1) * 100
						} else if nh != -1 {
							newVal = nh + 1
						}
						//fmt.Printf(" (%v)\n", newVal)

						if newVal != -1 {
							// We've found it!
							//fmt.Printf("Found new: %v\n", newVal)
							sum += newVal
							found = true
							break calcGrid
						}

						g[v][h] = '#'
					}
				}
			}

			if !found {
				panic("No new")
			}
		}
		return sum
	}
	// solve part 1 here

	p := 0
	for _, g := range grids {
		//fmt.Printf("Doing %v\n", g)
		pv, ph := calcReflectionRemainder(g, -1, -1)

		if pv != -1 {
			p += (pv + 1) * 100
		} else if ph != -1 {
			p += ph + 1
		} else {
			panic("Could not find")
		}
	}

	return p
}

func calcReflectionRemainder(grid [][]rune, exclV int, exclH int) (int, int) {
	// Try horizontal
	for v := 0; v < len(grid)-1; v++ {
		if v == exclV {
			continue
		}
		mirror := true
	m1:
		for h := 0; h < len(grid[v])-1; h++ {
			if grid[v][h] != grid[v+1][h] {
				mirror = false
				break m1
			}
		}

		if mirror {
			// Confirm a mirror
			lhs := v
			rhs := v + 1

			confirmedMirror := true

		m2:
			for lhs >= 0 && rhs < len(grid) {
				for h := 0; h < len(grid[lhs]); h++ {
					if grid[lhs][h] != grid[rhs][h] {
						confirmedMirror = false
						break m2
					}
				}
				lhs -= 1
				rhs += 1
			}

			if confirmedMirror {
				//fmt.Printf("Found at horiz %v-%v\n", v, v+1)
				return v, -1
			}
		}
	}

	// Else try vertical
	for h := 0; h < len(grid[0])-1; h++ {
		if h == exclH {
			continue
		}
		mirror := true
	m3:
		for v := 0; v < len(grid)-1; v++ {
			if grid[v][h] != grid[v][h+1] {
				mirror = false
				break m3
			}
		}

		if mirror {
			// Confirm a mirror
			uhs := h
			bhs := h + 1

			confirmedMirror := true

		m4:
			for uhs >= 0 && bhs < len(grid[0]) {
				for v := 0; v < len(grid); v++ {
					if grid[v][uhs] != grid[v][bhs] {
						confirmedMirror = false
						break m4
					}
				}
				uhs -= 1
				bhs += 1
			}

			if confirmedMirror {
				//fmt.Printf("Found at vert %v-%v\n", h, h+1)
				return -1, h
			}
		}
	}

	return -1, -1
}
