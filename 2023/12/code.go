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
		springs := make([]rune, 0)
		for _, r := range strings.Split(spl[0], "") {
			springs = append(springs, rune(r[0]))
		}
		p := iteratePossibilities(spl[0]+".", combo, 0)

		sum += p
	}

	return sum
}

func iteratePossibilities(springs string, combo []int, depth int) int {
	if len(combo) == 0 {
		panic("Can't try find possibilities with empty combo list")
	}
	sprlen := combo[0] + 1

	print(fmt.Sprintf("%v:", springs), depth)
	if len(springs) < sprlen {
		// Can't fit, bad match
		print(fmt.Sprintf("Cannot fit combo len %v in remaining springs %v", sprlen, springs), depth)
		return 0
	}

	p := 0
	for i := 0; i <= len(springs)-sprlen; i++ {
		match := true
		forcedMatch := true
		checking := springs[i : i+sprlen]
		print(fmt.Sprintf("Testing %v in %v", checking, springs), depth)

		for c, r := range checking {
			if rune(r) == '#' {
				if c == sprlen-1 {
					print("Failed match: ends with damaged (would be non-contiguous)", depth)
					match = false
				}
			} else if rune(r) == '.' {
				if c != sprlen-1 {
					print("Failed match: contained undamaged", depth)
					match = false
				}
			} else {
				if c != sprlen-1 {
					print(fmt.Sprintf("NOT forced match at %v", c), depth)
					forcedMatch = false
				}
			}
		}

		if match {
			print(fmt.Sprintf("Found %v+1 match (forced=%v) in %v at %v-%v (%v) (remaining: %v)", combo[0], forcedMatch, springs, i, i+sprlen-1, springs[i:i+sprlen], combo[1:]), depth)

			if len(combo) > 1 {
				p += iteratePossibilities(springs[i+sprlen:], combo[1:], depth+1)
			} else {
				p += 1
			}

			if forcedMatch {
				print("Forced match, not trying any other combos for this", depth)
				break
			}
		}
	}

	return p
}

func print(pr string, depth int) {
	print := false
	if print {
		for i := 0; i < depth; i++ {
			fmt.Printf("> ")
		}
		fmt.Println(pr)
	}
}
