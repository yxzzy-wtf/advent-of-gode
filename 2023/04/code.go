package main

import (
	"regexp"
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
	if part2 {
		return "not implemented"
	}
	// solve part 1 here

	gameIdRxp := regexp.MustCompile(`Card +(?P<ID>\d+):  ?(?P<Winning>.*) \|  ?(?P<Drawn>.*)`)

	sum := 0
	for _, line := range strings.Split(input, "\n") {

		matches := gameIdRxp.FindStringSubmatch(line)
		if len(matches) == 0 {
			continue
		}

		winning := strings.Split(matches[2], " ")
		drawn := strings.Split(matches[3], " ")

		points := 0
		for _, win := range winning {
			if win == "" {
				continue
			}

			for _, draw := range drawn {
				if win == draw {
					if points == 0 {
						points = 1
					} else {
						points = points * 2
					}
				}
			}
		}
		sum += points
	}

	return sum
}
