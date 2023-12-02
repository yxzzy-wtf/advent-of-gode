package main

import (
	"regexp"
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
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		sum := 0

		gameIdRxp := regexp.MustCompile(`Game (?P<ID>\d+):(?P<Data>.*)`)

		for _, line := range strings.Split(input, "\n") {

			matches := gameIdRxp.FindStringSubmatch(line)
			if len(matches) == 0 {
				continue
			}

			data := matches[2]

			redMax := 0
			greenMax := 0
			blueMax := 0

			for _, draw := range strings.Split(data, ";") {
				redDraw := 0
				greenDraw := 0
				blueDraw := 0

				for _, single := range strings.Split(strings.Trim(draw, " "), ",") {
					split := strings.Split(strings.Trim(single, " "), " ")

					count, _ := strconv.Atoi(strings.Trim(split[0], " "))
					colour := split[1]

					if colour == "red" {
						redDraw += count
					} else if colour == "blue" {
						blueDraw += count
					} else if colour == "green" {
						greenDraw += count
					}
				}

				if redDraw > redMax {
					redMax = redDraw
				}
				if greenDraw > greenMax {
					greenMax = greenDraw
				}
				if blueDraw > blueMax {
					blueMax = blueDraw
				}

			}

			power := redMax * greenMax * blueMax
			sum += power
		}

		return sum
	}
	// solve part 1 here

	redMax := 12
	greenMax := 13
	blueMax := 14
	sum := 0

	gameIdRxp := regexp.MustCompile(`Game (?P<ID>\d+):(?P<Data>.*)`)

	for _, line := range strings.Split(input, "\n") {

		matches := gameIdRxp.FindStringSubmatch(line)
		if len(matches) == 0 {
			continue
		}

		id, _ := strconv.Atoi(matches[1])
		data := matches[2]

		valid := true
	drawLoop:
		for _, draw := range strings.Split(data, ";") {
			redDraw := 0
			greenDraw := 0
			blueDraw := 0

			for _, single := range strings.Split(strings.Trim(draw, " "), ",") {
				split := strings.Split(strings.Trim(single, " "), " ")

				count, _ := strconv.Atoi(strings.Trim(split[0], " "))
				colour := split[1]

				if colour == "red" {
					redDraw += count
				} else if colour == "blue" {
					blueDraw += count
				} else if colour == "green" {
					greenDraw += count
				}

				if redDraw > redMax || blueDraw > blueMax || greenDraw > greenMax {
					valid = false
					break drawLoop
				}
			}
		}

		if valid {
			sum += id
		}

	}

	return sum
}
