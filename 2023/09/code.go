package main

import (
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
	data := make([][]int, 0)

	for _, line := range strings.Split(input, "\n") {
		dataLine := make([]int, 0)
		for _, s := range strings.Split(line, " ") {
			d, _ := strconv.Atoi(s)
			dataLine = append(dataLine, d)
		}
		data = append(data, dataLine)
	}

	if part2 {
		return "not implemented"
	}
	// solve part 1 here

	sumNexts := 0
	for _, d := range data {
		sumNexts += getNext(d)
	}

	return sumNexts
}

func getNext(data []int) int {
	diffTriangle := make([][]int, 0)
	diffTriangle = append(diffTriangle, data)

	for !allZero(diffTriangle[len(diffTriangle)-1]) {
		diffTriangle = append(diffTriangle, getDiffs(diffTriangle[len(diffTriangle)-1]))
		//fmt.Printf("%v\n", diffTriangle[len(diffTriangle)-1])
	}

	add := 0
	for i := len(diffTriangle) - 2; i >= 0; i -= 1 {
		add += diffTriangle[i][len(diffTriangle[i])-1]
	}

	return add
}

func getDiffs(data []int) []int {
	nextData := make([]int, 0)
	for i := 0; i < len(data)-1; i += 1 {
		nextData = append(nextData, data[i+1]-data[i])
	}
	return nextData
}

func allZero(data []int) bool {
	for _, d := range data {
		if d != 0 {
			return false
		}
	}
	return true
}
