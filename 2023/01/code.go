package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

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
		return "not implemented"
	}

	sum := 0
	for _, line := range strings.Split(input, "\n") {
		var firstChar rune
		var secondChar rune

		for i := range line {
			if unicode.IsNumber(firstChar) && unicode.IsNumber(secondChar) {
				break
			}

			if !unicode.IsNumber(firstChar) && unicode.IsNumber([]rune(line)[i]) {
				firstChar = []rune(line)[i]
			}

			if !unicode.IsNumber(secondChar) && unicode.IsNumber([]rune(line)[len(line)-i-1]) {
				secondChar = []rune(line)[len(line)-i-1]
			}
		}

		numberStr := fmt.Sprintf("%v%v", string(firstChar), string(secondChar))
		number, _ := strconv.Atoi(numberStr)

		sum += number
	}

	return sum
}
