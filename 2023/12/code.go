package main

import (
	"fmt"
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
var cache = make(map[string]int)

func run(part2 bool, input string) any {
	sum := 0
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		for _, l := range strings.Split(input, "\n") {
			if l == "" {
				continue
			}

			spl := strings.Split(l, " ")
			combo := make([]int, 0)
			comboBase := make([]int, 0)
			for _, c := range strings.Split(spl[1], ",") {
				ci, _ := strconv.Atoi(c)
				combo = append(combo, ci)
				comboBase = append(comboBase, ci)
			}

			springs := spl[0]

			// unfold
			for i := 0; i < 4; i++ {
				springs += "?" + spl[0]

				comboBase = append(comboBase, combo...)
			}
			poss := iteratePossibilities(springs, comboBase)

			sum += poss
		}

		return sum
	} else {
		// Part 1
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
			poss := iteratePossibilities(spl[0], combo)
			sum += poss
		}
	}
	// solve part 1 here

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
				p += 1
			}
		} else {
			matchKey := fmt.Sprintf("%v : %v", springs[matchFrom:], combo[1:])

			if precount, ok := cache[matchKey]; ok {
				p += precount
			} else {
				poss := iteratePossibilities(springs[matchFrom:], combo[1:])
				cache[matchKey] = poss
				p += poss
			}
		}

		if forceMatch {
			break
		}
	}

	return p
}
