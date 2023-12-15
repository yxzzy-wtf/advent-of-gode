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
type Lense struct {
	Label string
	Focus int
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		boxes := make([][]Lense, 256)
		for _, l := range strings.Split(input, "\n") {
			if l == "" {
				continue
			}

			for _, s := range strings.Split(l, ",") {
				if strings.Contains(s, "-") {
					// Dash logic
					spl := strings.Split(s, "-")
					label := spl[0]
					box := hash(label)

				remove:
					for i := range boxes[box] {
						if boxes[box][i].Label == label {
							boxes[box] = append(boxes[box][:i], boxes[box][i+1:]...)
							break remove
						}
					}

				} else if strings.Contains(s, "=") {
					// Eq logic
					spl := strings.Split(s, "=")
					label := spl[0]
					lens, _ := strconv.Atoi(spl[1])

					box := hash(label)

					insert := Lense{label, lens}
					replaced := false

				replace:
					for i := range boxes[box] {
						if boxes[box][i].Label == label {
							boxes[box][i] = insert
							replaced = true
							break replace
						}
					}

					if !replaced {
						boxes[box] = append(boxes[box], insert)
					}

				} else {
					panic(fmt.Sprintf("bad control '%v'", s))
				}
			}
		}

		// Calculate
		sum := 0
		for b := range boxes {
			for l := range boxes[b] {
				val := (1 + b) * (1 + l) * boxes[b][l].Focus
				sum += val
			}
		}

		return sum
	}
	// solve part 1 here

	hsum := 0
	for _, l := range strings.Split(input, "\n") {
		for _, s := range strings.Split(l, ",") {
			hsum += hash(s)
		}
	}

	return hsum
}

func hash(str string) int {
	h := 0
	for _, s := range str {
		ascii := int(rune(s))
		h += ascii
		h *= 17
		h = h % 256
	}
	return h
}
