package main

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func getNumberAndBlank(schema [][]rune, x int, y int) int {
	if x < 0 || x >= len(schema) || y < 0 || y >= len(schema[x]) {
		// Out of bounds
		return 0
	}

	if !unicode.IsNumber(schema[x][y]) {
		// Not a number to start
		return 0
	}

	// Moves along y to find the first character...
	firstFound := false
	for !firstFound {
		if y == 0 || !unicode.IsNumber(schema[x][y-1]) {
			firstFound = true
		} else {
			y -= 1
		}
	}

	// then assemble
	numStr := ""
	for y < len(schema[x]) && unicode.IsNumber(schema[x][y]) {
		numStr += string(schema[x][y])
		schema[x][y] = '.'
		y += 1
	}

	num, _ := strconv.Atoi(numStr)

	return num
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// Assemble schematic
	schematic := make([][]rune, 0)
	for _, line := range strings.Split(input, "\n") {
		if line != "" {
			schematic = append(schematic, []rune(line))
		}
	}

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here

	// Go through array, if we find a non-number non-period, find all adjacent numbers and then turn those to periods
	engineSum := 0
	for x, _ := range schematic {
		for y, _ := range schematic[x] {
			if unicode.IsNumber(schematic[x][y]) || schematic[x][y] == '.' {
				continue
			}

			// Else, it's a symbol! Go around and, if there are any numbers, assemble them
			for xi := x - 1; xi < x+2; xi++ {
				for yi := y - 1; yi < y+2; yi++ {
					engineSum += getNumberAndBlank(schematic, xi, yi)
				}
			}

		}
	}

	return engineSum
}
