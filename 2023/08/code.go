package main

import (
	"regexp"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Node struct {
	Name  string
	Left  string
	Right string
}

func getLcm(a int, b int) int {
	return (a * b) / getGcd(a, b)
}

func getGcd(a int, b int) int {
	g := 0
	for i := 1; i <= a && i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			g = i
		}
	}
	return g
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	nodeMap := make(map[string]*Node)

	lines := strings.Split(input, "\n")
	instructions := lines[0]

	nodeRegex := regexp.MustCompile(`(?P<Name>[A-Z]+) = \((?P<Left>[A-Z]+), (?P<Right>[A-Z]+)\)`)
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		matches := nodeRegex.FindStringSubmatch(line)

		name := matches[1]
		left := matches[2]
		right := matches[3]

		thisNode := Node{name, left, right}
		nodeMap[name] = &thisNode
	}

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		ghostNodes := make([]*Node, 0)
		for nodeName, node := range nodeMap {
			if rune(nodeName[2]) == 'A' {
				// Start node
				ghostNodes = append(ghostNodes, node)
			}
		}

		// Find the minimal... something?
		minSteps := make([]int, 0)
		for _, ghostNode := range ghostNodes {
			node := ghostNode
			steps := 0

			for rune(node.Name[2]) != 'Z' {
				instr := rune(instructions[steps%len(instructions)])

				if instr == 'L' {
					// Left
					node = nodeMap[node.Left]
				} else {
					// Right
					node = nodeMap[node.Right]
				}

				steps += 1
			}

			minSteps = append(minSteps, steps)
		}

		lcm := 1
		for _, st := range minSteps {
			lcm = getLcm(lcm, st)
		}

		return lcm
	}
	// solve part 1 here

	node := nodeMap["AAA"]
	totalSteps := 0
	for node.Name != "ZZZ" {
		instr := rune(instructions[totalSteps%len(instructions)])

		if instr == 'L' {
			// Left
			node = nodeMap[node.Left]
		} else {
			// Right
			node = nodeMap[node.Right]
		}

		totalSteps += 1
	}

	return totalSteps
}
