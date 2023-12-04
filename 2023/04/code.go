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
	gameIdRxp := regexp.MustCompile(`Card +(?P<ID>\d+):  ?(?P<Winning>.*) \|  ?(?P<Drawn>.*)`)

	if part2 {
		lines := strings.Split(strings.Trim(input, "\n"), "\n")

		points := make([]int, len(lines))
		for i, line := range lines {

			matches := gameIdRxp.FindStringSubmatch(line)
			if len(matches) == 0 {
				continue
			}

			winning := strings.Split(matches[2], " ")
			drawn := strings.Split(matches[3], " ")

			points[i] = 0
			for _, win := range winning {
				if win == "" {
					continue
				}

				for _, draw := range drawn {
					if win == draw {
						points[i] += 1
					}
				}
			}
		}

		copies := make([]int, len(points))
		for i := range points {
			copies[i] = 1
		}
		for i, p := range points {
			for j := i + 1; int(j) <= int(i)+p; j++ {
				if j < len(copies) {
					copies[j] += copies[i]
				}
			}
		}

		totalCards := int(0)
		for _, count := range copies {
			totalCards += count
		}

		return totalCards
	}
	// solve part 1 here

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
