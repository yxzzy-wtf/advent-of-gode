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
func testMatch(text string, match string, i int) bool {
	if len(text)-i-len(match) < 0 {
		return false
	}

	return text[i:i+len(match)] == match
}

func getNumber(text string, i int) rune {
	if unicode.IsNumber(rune(text[i])) {
		return rune(text[i])
	} else {
		// Might be a string
		if testMatch(text, "one", i) {
			return rune("1"[0])
		}
		if testMatch(text, "two", i) {
			return rune("2"[0])
		}
		if testMatch(text, "three", i) {
			return rune("3"[0])
		}
		if testMatch(text, "four", i) {
			return rune("4"[0])
		}
		if testMatch(text, "five", i) {
			return rune("5"[0])
		}
		if testMatch(text, "six", i) {
			return rune("6"[0])
		}
		if testMatch(text, "seven", i) {
			return rune("7"[0])
		}
		if testMatch(text, "eight", i) {
			return rune("8"[0])
		}
		if testMatch(text, "nine", i) {
			return rune("9"[0])
		}
	}
	return rune("f"[0])
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {

		sum := 0
		for _, line := range strings.Split(input, "\n") {
			if line == "" {
				continue
			}
			var firstChar rune
			var secondChar rune

			for i := range line {
				if unicode.IsNumber(firstChar) && unicode.IsNumber(secondChar) {
					break
				}

				if !unicode.IsNumber(firstChar) {
					testFirst := getNumber(line, i)
					if unicode.IsNumber(testFirst) {
						firstChar = testFirst
					}
				}

				if !unicode.IsNumber(secondChar) {
					testSecond := getNumber(line, len(line)-1-i)
					if unicode.IsNumber(testSecond) {
						secondChar = testSecond
					}
				}
			}

			numberStr := fmt.Sprintf("%v%v", string(firstChar), string(secondChar))
			number, _ := strconv.Atoi(numberStr)

			sum += number
		}

		return sum
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
