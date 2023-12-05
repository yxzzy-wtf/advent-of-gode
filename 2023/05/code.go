package main

import (
	"math"
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

func mapRange(ranges [][]int, t int) int {
	for _, r := range ranges {
		if t >= r[0] && t <= r[1] {
			// In range, map
			diff := t - r[0]
			return r[2] + diff
		}
	}
	return t
}

func buildOnRange(ranges [][]int, line string) [][]int {
	lineSplit := strings.Split(line, " ")

	rangeStart, _ := strconv.Atoi(lineSplit[1])
	rangeLength, _ := strconv.Atoi(lineSplit[2])
	mapStart, _ := strconv.Atoi(lineSplit[0])

	return append(ranges, []int{rangeStart, rangeStart + rangeLength - 1, mapStart})
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		lines := strings.Split(input, "\n")

		seedPairs := regexp.MustCompile(`\d+ \d+`).FindAllString(lines[0], -1)[1:]
		seeds := make([][]int, 0)
		for _, s := range seedPairs {
			if s == "" {
				continue
			}
			spl := strings.Split(s, " ")
			start, _ := strconv.Atoi(spl[0])
			end, _ := strconv.Atoi(spl[1])

			seeds = append(seeds, []int{start, end})
		}

		mapList := make([][][]int, 0)
		for l := 2; l < len(lines); l++ {
			line := lines[l]

			if line == "" {
				continue
			}

			if line[len(line)-1:] == ":" {
				// Start a new mapping chain
				m := make([][]int, 0)
				for i := l + 1; i < len(lines) && lines[i] != ""; i++ {
					m = buildOnRange(m, lines[i])
					l += 1
				}

				mapList = append(mapList, m)
			}
		}

		// Now go through each seed
		lowest := math.MaxInt
		for _, sp := range seeds {
			for si := sp[0]; si < sp[0]+sp[1]; si++ {
				//fmt.Printf("%v\n", sp[0]+sp[1])

				sm := si
				for _, r := range mapList {
					mappedTo := mapRange(r, sm)
					sm = mappedTo
				}

				if lowest > sm {
					lowest = sm
				}
			}

		}

		return lowest
	}
	// solve part 1 here

	lines := strings.Split(input, "\n")

	seeds := strings.Split(regexp.MustCompile(`seeds: (?P<seeds>.*)`).FindStringSubmatch(lines[0])[1], " ")

	mapList := make([][][]int, 0)
	for l := 2; l < len(lines); l++ {
		line := lines[l]

		if line == "" {
			continue
		}

		if line[len(line)-1:] == ":" {
			// Start a new mapping chain
			m := make([][]int, 0)
			for i := l + 1; i < len(lines) && lines[i] != ""; i++ {
				m = buildOnRange(m, lines[i])
				l += 1
			}

			mapList = append(mapList, m)
		}
	}

	// Now go through each seed
	lowest := math.MaxInt
	for _, s := range seeds {
		si, _ := strconv.Atoi(s)

		for _, r := range mapList {
			mappedTo := mapRange(r, si)
			si = mappedTo
		}

		if lowest > si {
			lowest = si
		}
	}

	return lowest
}
