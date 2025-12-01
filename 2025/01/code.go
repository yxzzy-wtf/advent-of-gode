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

		dial_val := 50
		zcount := 0
		for _, line := range strings.Split(input, "\n") {
			if line == "" {
				continue
			}

			dir, cstr := line[0], line[1:]
			c, _ := strconv.Atoi(cstr)

			full_rotations := c / 100
			zcount += full_rotations

			c -= full_rotations * 100

			new_val := dial_val
			if dir == 'R' {
				new_val += c
				if new_val > 99 {
					new_val -= 100
					zcount++
				}
			} else {
				new_val -= c
				if new_val < 0 {
					new_val += 100
					if dial_val != 0 {
						zcount++
					}
				} else if new_val == 0 {
					zcount++
				}
			}

			dial_val = new_val
		}

		return zcount
	}
	// solve part 1 here

	dial_val := 50
	zcount := 0
	for _, line := range strings.Split(input, "\n") {
		//fmt.Printf("%v\n", line)
		if line == "" {
			continue
		}

		dir, cstr := line[0], line[1:]
		c, _ := strconv.Atoi(cstr)
		c = c % 100

		if dir == 'R' {
			dial_val += c
		} else {
			dial_val += (100 - c)
		}

		if dial_val >= 100 {
			dial_val = dial_val % 100
		}

		if dial_val == 0 {
			zcount++
		}
		//fmt.Printf("Pos: %v (%v)\n", dial_val, zcount)
	}

	return zcount
}
