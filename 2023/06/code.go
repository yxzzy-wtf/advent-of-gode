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

func calcDist(timeCharged int, totalTime int) int {
	if timeCharged == 0 || totalTime-timeCharged == 0 {
		return 0
	}

	return (totalTime - timeCharged) * timeCharged
}

func canBeat(timeCharged int, data []int) bool {
	return calcDist(timeCharged, data[0]) > data[1]
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	pairs := make([][]int, 0)
	var lines []string
	if part2 {
		lines = strings.Split(strings.ReplaceAll(input, " ", ""), "\n")
	} else {
		lines = strings.Split(input, "\n")
	}

	numGet := regexp.MustCompile(`\d+`)
	timeS := numGet.FindAllString(lines[0], -1)
	distS := numGet.FindAllString(lines[1], -1)

	for i := range timeS {
		time, _ := strconv.Atoi(timeS[i])
		dist, _ := strconv.Atoi(distS[i])
		pairs = append(pairs, []int{time, dist})
	}
	// solve part 1 here

	// Binary search to find any midpoint where you can win
	marginOfError := 1
	for _, pair := range pairs {
		possibleStartRange := []int{0, pair[0]}
		start := pair[0] / 2

		for !(!canBeat(start-1, pair) && canBeat(start, pair)) {
			if canBeat(start, pair) {
				// must go lower
				possibleStartRange[1] = start
			} else {
				// Must go higher
				possibleStartRange[0] = start
			}
			start = possibleStartRange[0] + (possibleStartRange[1]-possibleStartRange[0])/2
		}

		possibleEndRange := []int{0, pair[0]}
		end := pair[0] / 2

		for !(!canBeat(end+1, pair) && canBeat(end, pair)) {
			if !canBeat(end, pair) {
				// must go lower
				possibleEndRange[1] = end
			} else {
				// Must go higher
				possibleEndRange[0] = end
			}
			end = possibleEndRange[0] + (possibleEndRange[1]-possibleEndRange[0])/2
		}

		marginOfError *= end - start + 1
	}

	return marginOfError
}
