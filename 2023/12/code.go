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
	if part2 {
		return "not implemented"
	}
	// solve part 1 here

	sum := 0
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}

		spl := strings.Split(l, " ")
		combo := make([]int, 0)
		for _, c := range strings.Split(spl[1], ",") {
			ci, _ := strconv.Atoi(c)
			combo = append(combo, ci)
		}

		//fmt.Printf("\nLooking   in %v %v\n", spl[0], combo)
		poss := iteratePossibilities(spl[0], combo)

		//fmt.Printf("In %v %v: %v\n", spl[0], combo, poss)

		sum += poss
	}

	return sum
}

func iteratePossibilities(springs string, combo []int) int {
	if len(combo) == 0 {
		panic("Combo is empty")
	}

	lenMatch := combo[0]

	if len(springs) < lenMatch {
		// Cannot match, not enough space
		return 0
	}

	p := 0
	for i := 0; i < len(springs)-lenMatch+1; i++ {
		match := true
		forceMatch := false
		for m := i; m < i+lenMatch; m++ {
			if rune(springs[m]) == '.' {
				match = false
			}

			if rune(springs[m]) == '#' && i == m {
				forceMatch = true
			}
		}

		if !match {
			// Keep trying
			if forceMatch {
				break
			}
			continue
		}

		matchFrom := i + lenMatch
		if len(springs) > i+lenMatch {
			if rune(springs[i+lenMatch]) == '#' {
				// Would be contigious, does not work
				if forceMatch {
					break
				}
				continue
			}
			matchFrom += 1
		}

		// Cool! We are matched!
		if len(combo) == 1 {
			// Check, finally, that there are no possible # in the future (which would violate the requirements)
			anyFuture := false
			for j := matchFrom; j < len(springs); j++ {
				if rune(springs[j]) == '#' {
					anyFuture = true
				}
			}

			if !anyFuture {
				//fmt.Printf("MATCHED %v in %v from [%v;%v)\n", lenMatch, springs, i, i+lenMatch-1)
				p += 1
			}
		} else {
			//fmt.Printf("Matched %v in %v from [%v;%v)\n", lenMatch, springs, i, i+lenMatch-1)
			p += iteratePossibilities(springs[matchFrom:], combo[1:])
		}

		if forceMatch {
			break
		}
	}

	return p
}
