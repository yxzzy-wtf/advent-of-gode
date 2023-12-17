package main

import (
	"fmt"
	"math"
	"sort"
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

type Path struct {
	CurrentV         int
	CurrentH         int
	CurrentDirection rune
	CurrentRemaining int
	HeatLoss         int
	Steps            []Step
}

type Step struct {
	V int
	H int
}

type WeightedCell struct {
	V        int
	H        int
	HeatLoss int

	N []int
	S []int
	E []int
	W []int
}

// What if each cell

func removeDuplicates(paths []Path) []Path {
	set := make(map[string]bool)

	result := make([]Path, 0)

	for _, p := range paths {
		str := fmt.Sprintf("%v;%v DIR:%v(%v)", p.CurrentV, p.CurrentH, p.CurrentDirection, p.CurrentRemaining)
		if _, ok := set[str]; !ok {
			result = append(result, p)
			set[str] = true
		}
	}

	return result
}

func run(part2 bool, input string) any {
	grid := make([][]WeightedCell, 0)
	for v, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}

		gridLine := make([]WeightedCell, 0)
		for h, c := range strings.Split(l, "") {
			m, _ := strconv.Atoi(c)
			wc := WeightedCell{v, h, m,
				[]int{math.MaxInt, math.MaxInt, math.MaxInt},
				[]int{math.MaxInt, math.MaxInt, math.MaxInt},
				[]int{math.MaxInt, math.MaxInt, math.MaxInt},
				[]int{math.MaxInt, math.MaxInt, math.MaxInt}}
			gridLine = append(gridLine, wc)
		}
		grid = append(grid, gridLine)
	}

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here

	grid[0][0].N = []int{grid[0][0].HeatLoss, grid[0][0].HeatLoss, grid[0][0].HeatLoss}
	grid[0][0].S = []int{grid[0][0].HeatLoss, grid[0][0].HeatLoss, grid[0][0].HeatLoss}
	grid[0][0].E = []int{grid[0][0].HeatLoss, grid[0][0].HeatLoss, grid[0][0].HeatLoss}
	grid[0][0].W = []int{grid[0][0].HeatLoss, grid[0][0].HeatLoss, grid[0][0].HeatLoss}

	paths := []Path{
		{0, 0, 'E', 3, grid[0][0].HeatLoss, []Step{{0, 0}}},
		{0, 0, 'S', 3, grid[0][0].HeatLoss, []Step{{0, 0}}},
	}

	var p Path
	for len(paths) > 0 {
		p, paths = paths[0], paths[1:]

		if p.CurrentV == len(grid)-1 && p.CurrentH == len(grid[0])-1 {
			break
		}

		// Check L and R
		if p.CurrentDirection == 'N' || p.CurrentDirection == 'S' {
			// E & W
			eH := p.CurrentH + 1
			if eH < len(grid[p.CurrentV]) {
				eC := grid[p.CurrentV][eH]
				if eC.HeatLoss+p.HeatLoss < eC.E[2] {
					eC.E[2] = eC.HeatLoss + p.HeatLoss

					eClone := make([]Step, len(p.Steps))
					copy(eClone, p.Steps)
					eClone = append(eClone, Step{p.CurrentV, eH})

					ePath := Path{p.CurrentV, eH, 'E', 2, eC.E[2], eClone}
					paths = append(paths, ePath)
				}
			}

			wH := p.CurrentH - 1
			if wH >= 0 {
				wC := grid[p.CurrentV][wH]
				if wC.HeatLoss+p.HeatLoss < wC.W[2] {
					wC.W[2] = wC.HeatLoss + p.HeatLoss

					wClone := make([]Step, len(p.Steps))
					copy(wClone, p.Steps)
					wClone = append(wClone, Step{p.CurrentV, wH})

					wPath := Path{p.CurrentV, wH, 'W', 2, wC.W[2], wClone}
					paths = append(paths, wPath)
				}
			}
		} else if p.CurrentDirection == 'E' || p.CurrentDirection == 'W' {
			nV := p.CurrentV - 1
			if nV >= 0 {
				nC := grid[nV][p.CurrentH]
				if nC.HeatLoss+p.HeatLoss < nC.N[2] {
					nC.N[2] = nC.HeatLoss + p.HeatLoss

					nClone := make([]Step, len(p.Steps))
					copy(nClone, p.Steps)
					nClone = append(nClone, Step{nV, p.CurrentH})

					nPath := Path{nV, p.CurrentH, 'N', 2, nC.N[2], nClone}
					paths = append(paths, nPath)
				}
			}

			sV := p.CurrentV + 1
			if sV < len(grid) {
				sC := grid[sV][p.CurrentH]
				if sC.HeatLoss+p.HeatLoss < sC.S[2] {
					sC.S[2] = sC.HeatLoss + p.HeatLoss

					sClone := make([]Step, len(p.Steps))
					copy(sClone, p.Steps)
					sClone = append(sClone, Step{sV, p.CurrentH})

					sPath := Path{sV, p.CurrentH, 'S', 2, sC.S[2], sClone}
					paths = append(paths, sPath)
				}
			}
		} else {
			panic("bad direction")
		}

		if p.CurrentRemaining > 0 {
			// Take one step forward
			p.CurrentRemaining -= 1

			if p.CurrentDirection == 'N' {
				nV := p.CurrentV - 1
				if nV >= 0 {
					nC := grid[nV][p.CurrentH]

					if nC.HeatLoss+p.HeatLoss < nC.N[p.CurrentRemaining] {
						// It's less on this path! Continue!
						p.HeatLoss += nC.HeatLoss
						nC.N[p.CurrentRemaining] = p.HeatLoss

						p.CurrentV = nV
						p.Steps = append(p.Steps, Step{nV, p.CurrentH})
						paths = append(paths, p)
					}
				}
			} else if p.CurrentDirection == 'S' {
				sV := p.CurrentV + 1
				if sV < len(grid) {
					sC := grid[sV][p.CurrentH]

					if sC.HeatLoss+p.HeatLoss < sC.S[p.CurrentRemaining] {
						// It's less on this path! Continue!
						p.HeatLoss += sC.HeatLoss
						sC.S[p.CurrentRemaining] = p.HeatLoss

						p.CurrentV = sV
						p.Steps = append(p.Steps, Step{sV, p.CurrentH})
						paths = append(paths, p)
					}
				}
			} else if p.CurrentDirection == 'E' {
				eH := p.CurrentH + 1
				if eH < len(grid[p.CurrentV]) {
					eC := grid[p.CurrentV][eH]

					if eC.HeatLoss+p.HeatLoss < eC.E[p.CurrentRemaining] {
						// It's less on this path! Continue!
						p.HeatLoss += eC.HeatLoss
						eC.E[p.CurrentRemaining] = p.HeatLoss

						p.CurrentH = eH
						p.Steps = append(p.Steps, Step{p.CurrentV, eH})
						paths = append(paths, p)
					}
				}
			} else if p.CurrentDirection == 'W' {
				wH := p.CurrentH - 1
				if wH >= 0 {
					wC := grid[p.CurrentV][wH]

					if wC.HeatLoss+p.HeatLoss < wC.W[p.CurrentRemaining] {
						// It's less on this path! Continue!
						p.HeatLoss += wC.HeatLoss
						wC.W[p.CurrentRemaining] = p.HeatLoss

						p.CurrentH = wH
						p.Steps = append(p.Steps, Step{p.CurrentV, wH})
						paths = append(paths, p)
					}
				}
			} else {
				panic("bad dir")
			}
		}

		sort.Slice(paths, func(i, j int) bool {
			return paths[i].HeatLoss < paths[j].HeatLoss
		})
	}

	fmt.Println("Printing solve:")

	for v := range grid {
		for h := range grid[v] {
			printed := false
		printHash:
			for _, s := range p.Steps {
				if s.H == h && s.V == v {
					printed = true
					fmt.Printf("#")
					break printHash
				}
			}

			if !printed {
				fmt.Printf("%v", grid[v][h].HeatLoss)
			}
		}
		fmt.Println()
	}

	// Are we not counting the heat loss at position [0][0]?

	return p.HeatLoss - grid[0][0].HeatLoss
}
