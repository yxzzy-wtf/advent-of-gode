package main

import (
	"math"
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

	universe := make([][]rune, 0)
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}
		line := make([]rune, 0)
		empty := true
		for r := range strings.Split(l, "") {
			if rune(l[r]) != '.' {
				empty = false
			}
			line = append(line, rune(l[r]))
		}

		if empty {
			for e := range line {
				line[e] = '↕'
			}
		}
		universe = append(universe, line)
	}

	// Now expand inner cols
	for h := 0; h < len(universe[0]); h++ {
		empty := true
		for v := range universe {
			if universe[v][h] != '.' && universe[v][h] != '↕' {
				empty = false
			}
		}

		if empty {
			for v := range universe {
				if universe[v][h] == '↕' {
					universe[v][h] = '⥁'
				} else {
					universe[v][h] = '↔'
				}
			}
		}
	}

	// when you're ready to do part 2, remove this "not implemented" block
	factor := 1
	if part2 {
		factor = 1000000
	} else {
		// solve part 1 here
		factor = 2
	}

	coords := make([][]int, 0)

	vr := 0
	for v := range universe {
		if universe[v][0] == '↕' || universe[v][0] == '⥁' {
			// expand else normal v
			vr += factor
		} else {
			vr += 1
		}

		hr := 0
		for h := range universe {
			if universe[0][h] == '↔' || universe[0][h] == '⥁' {
				hr += factor
			} else {
				hr += 1
			}

			if universe[v][h] == '#' {
				coords = append(coords, []int{vr, hr})
			}
		}
	}

	sumHyp := 0
	compared := 0
	for c1 := 0; c1 < len(coords)-1; c1++ {
		for c2 := c1 + 1; c2 < len(coords); c2++ {
			// add rise and run?
			lenv := int(math.Abs(float64(coords[c1][0] - coords[c2][0])))
			lenh := int(math.Abs(float64(coords[c1][1] - coords[c2][1])))

			// convert to int?
			sumHyp += lenv + lenh

			compared++
		}
	}

	return sumHyp
}
