package main

import (
	"fmt"
	"math"
	"sort"
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
type Node struct {
	Y    int
	X    int
	Next *Node
	Prev *Node
}

func printNodes(nodes []*Node) {
	maxY := 0
	maxX := 0
	minY := math.MaxInt
	minX := math.MaxInt

	for _, n := range nodes {
		if maxY < n.Y {
			maxY = n.Y
		}
		if minY > n.Y {
			minY = n.Y
		}

		if maxX < n.X {
			maxX = n.X
		}
		if minX > n.X {
			minX = n.X
		}
	}

	print := false
	if print {
		gridToPrint := make([][]rune, 0)
		for y := 0; y <= maxY-minY; y++ {
			gridLine := make([]rune, 0)
			for x := 0; x <= maxX-minX; x++ {
				gridLine = append(gridLine, ' ')
			}
			gridToPrint = append(gridToPrint, gridLine)
		}

		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				for _, n := range nodes {
					//fmt.Printf("Checking node %v\n", n)
					if n.X == x && n.Y == y {
						if n.Prev.X == n.X {
							// From previous, went up/down
							if n.Prev.Y < n.Y {
								// Has gone down
								if n.Next.X > n.X {
									// And right
									gridToPrint[y-minY][x-minX] = '└'
								} else {
									// And left
									gridToPrint[y-minY][x-minX] = '┘'
								}
							} else {
								// Has gone up
								if n.Next.X > n.X {
									// And right
									gridToPrint[y-minY][x-minX] = '┌'
								} else {
									// And left
									gridToPrint[y-minY][x-minX] = '┐'
								}
							}
						} else if n.Prev.Y == n.Y {
							// From previous, went left/right
							if n.Prev.X > n.X {
								// Has gone left
								if n.Next.Y > n.Y {
									// And down
									gridToPrint[y-minY][x-minX] = '┌'
								} else {
									// And up
									gridToPrint[y-minY][x-minX] = '└'
								}
							} else {
								// Has gone right
								if n.Next.Y > n.Y {
									// And down
									gridToPrint[y-minY][x-minX] = '┐'
								} else {
									// And up
									gridToPrint[y-minY][x-minX] = '┘'
								}
							}
						} else {
							panic(fmt.Sprintf("Could not align %v with Prev: %v", n, n.Prev))
						}
					}
				}
			}
		}

		// Now connect
		for y := range gridToPrint {
			for x := range gridToPrint[y] {
				if gridToPrint[y][x] == '┌' || gridToPrint[y][x] == '└' {
					xRun := x + 1
					for gridToPrint[y][xRun] == ' ' {
						gridToPrint[y][xRun] = '─'
						xRun++
					}
				}

				if gridToPrint[y][x] == '┌' || gridToPrint[y][x] == '┐' {
					yRun := y + 1
					for gridToPrint[yRun][x] == ' ' {
						gridToPrint[yRun][x] = '│'
						yRun++
					}
				}
			}
		}

		for y := range gridToPrint {
			for x := range gridToPrint[y] {
				fmt.Printf("%v", string(gridToPrint[y][x]))
			}
			fmt.Println()
		}
	}
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

var rad90 float64 = 90 * math.Pi / 180

func rotate(nodes []*Node) {

	for _, n := range nodes {
		// rotate all points by 90 degrees
		newX := int(math.Round(float64(n.X)*math.Cos(rad90) - float64(n.Y)*math.Sin(rad90)))
		newY := int(math.Round(float64(n.X)*math.Sin(rad90) + float64(n.Y)*math.Cos(rad90)))

		n.X, n.Y = newX, newY
	}
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		// More complex

		// shrink wrap?

		return "not implemented"
	}
	// solve part 1 here

	// Shrinkwrap approach

	nodes := make([]*Node, 0)
	x := 0
	y := 0
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}

		spl := strings.Split(l, " ")

		order := rune(spl[0][0])
		dist, _ := strconv.Atoi(spl[1])
		//colour := spl[2]

		node := &Node{Y: y, X: x}
		nodes = append(nodes, node)

		if order == 'R' {
			x += dist
		} else if order == 'L' {
			x -= dist
		} else if order == 'U' {
			y -= dist
		} else if order == 'D' {
			y += dist
		} else {
			panic("bad order")
		}
	}

	// Connect nodes (currently in order

	for n := range nodes {
		if n == 0 {
			nodes[n].Next = nodes[1]
			nodes[n].Prev = nodes[len(nodes)-1]
		} else if n == len(nodes)-1 {
			nodes[n].Next = nodes[0]
			nodes[n].Prev = nodes[n-1]
		} else {
			nodes[n].Next = nodes[n+1]
			nodes[n].Prev = nodes[n-1]
		}
	}

	firstNode := nodes[0]
	lastNode := nodes[len(nodes)-1]

	firstNode.Prev = lastNode
	lastNode.Next = firstNode

	// Begin downward collapse
	printNodes(nodes)
	totalArea := 0

	for len(nodes) > 0 {
		remove := make([]*Node, 0)

		// try to find an "unbreakable, top-layer node"

		/*

			imagine:

			............
			.##########.
			.##########.
			.##########.
			.####..####.
			.####..####.
			.####..####.
			.####.......
			............

		*/

		// Find the topmost node (i.e. lowest Y)
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].Y < nodes[j].Y
		})

		n1 := nodes[0]
		var n2, n3, n1adj, n2adj *Node

		if n1.Next.Y == n1.Y {
			n2 = n1.Next
			n1adj = n1.Prev
			n2adj = n2.Next
		} else {
			n2 = n1.Prev
			n1adj = n1.Next
			n2adj = n2.Prev
		}

		// Confirm that this is "unbreakable", else rotate and continue
		if n1adj.Y < n2adj.Y {
			n3 = n1adj
		} else {
			n3 = n2adj
		}
		// If there are any other nodes inside this box, reject

		unbreakable := true

		maxY := n3.Y
		var minX, maxX int
		if n1.X < n2.X {
			minX = n1.X
			maxX = n2.X
		} else {
			minX = n2.X
			maxX = n1.X
		}

		for _, n := range nodes {
			if n == n1 || n == n2 || n == n3 {
				continue
			}

			if n.X > minX && n.X < maxX && n.Y < maxY {
				unbreakable = false
			}
		}

		if !unbreakable {
			// try to find any other unbreakables
			exclusionX := [][]int{{n1.X, n2.X}}

		excludeX:
			for _, n := range nodes {
				n1 = n

				if n1.Next.Y == n1.Y {
					n2 = n1.Next
					n1adj = n1.Prev
					n2adj = n2.Next
				} else {
					n2 = n1.Prev
					n1adj = n1.Next
					n2adj = n2.Prev
				}

				if n1adj.Y < n1.Y || n2adj.Y < n1.Y {
					continue excludeX
				}

				for _, ex := range exclusionX {
					if ex[0] > ex[1] {
						if n1.X <= ex[0] && n1.X >= ex[1] || n2.X <= ex[0] && n2.X >= ex[1] {
							continue excludeX
						}
					} else {
						if n1.X >= ex[0] && n1.X <= ex[1] || n2.X >= ex[0] && n2.X <= ex[1] {
							continue excludeX
						}
					}
				}

				// Confirm that this is "unbreakable", else rotate and continue
				if n1adj.Y < n2adj.Y {
					n3 = n1adj
				} else {
					n3 = n2adj
				}

				unbreakable = true
				maxY := n3.Y
				var minX, maxX int
				if n1.X < n2.X {
					minX = n1.X
					maxX = n2.X
				} else {
					minX = n2.X
					maxX = n1.X
				}
				for _, n := range nodes {
					if n == n1 || n == n2 || n == n3 {
						continue
					}

					if n.X > minX && n.X < maxX && n.Y < maxY {
						unbreakable = false
					}
				}

				if !unbreakable {
					exclusionX = append(exclusionX, []int{n1.X, n2.X})
				} else {
					break excludeX
				}
			}

			if !unbreakable {
				rotate(nodes)
				printNodes(nodes)
				continue
			}
		}

		// We now have our "top" bar (n1-n2)
		// Find the minimum to shift it by
		// Now for the clever bit
		xLen := abs(n1.X-n2.X) + 1
		fmt.Printf("Bar len: %v (%v %v %v) \n", xLen, n1, n2, n3)

		var addY int
		if n1adj.Y < n2adj.Y {
			addY = n1adj.Y
		} else {
			addY = n2adj.Y
		}
		addY -= n1.Y
		fmt.Printf("Shift down by: %v\n", addY)

		areaCollapse := xLen * addY
		totalArea += areaCollapse

		// And collapse the nodes
		n1.Y += addY
		n2.Y += addY

		if n1.Y == n1adj.Y && n1.X == n1adj.X {
			// merge into
			remove = append(remove, n1)
			if n1.Next == n1adj {
				n1adj.Prev = n1.Prev
				n1.Prev.Next = n1adj
			} else if n1.Prev == n1adj {
				n1adj.Next = n1.Next
				n1.Next.Prev = n1adj
			} else {
				panic("Matched 1 but could not reassign")
			}
		}

		if n2.Y == n2adj.Y && n2.X == n2adj.X {
			remove = append(remove, n2)
			if n2.Next == n2adj {
				n2adj.Prev = n2.Prev
				n2.Prev.Next = n2adj
			} else if n2.Prev == n2adj {
				n2adj.Next = n2.Next
				n2.Next.Prev = n2adj
			} else {
				panic("Matched 2 but could not reassign")
			}
		}

		// finally, find any sharp corners (3 nodes in a row with the same Y)
		if n1adj.Next.Y == n1adj.Y && n1adj.Prev.Y == n1adj.Y {
			if (n1adj.X > n1adj.Next.X && n1adj.X > n1adj.Prev.X) || (n1adj.X < n1adj.Next.X && n1adj.X < n1adj.Prev.X) {
				diffNextPrev := abs(n1adj.Next.X - n1adj.Prev.X)
				diffBoth := abs(n1adj.X-n1adj.Next.X) + abs(n1adj.X-n1adj.Prev.X)
				areaCount := (diffBoth - diffNextPrev) / 2
				fmt.Printf("Offcut area: %v\n", areaCount)
				totalArea += areaCount

			}
			remove = append(remove, n1adj)
			n1adj.Next.Prev = n1adj.Prev
			n1adj.Prev.Next = n1adj.Next
		}

		if n2adj.Next.Y == n2adj.Y && n2adj.Prev.Y == n2adj.Y {
			if (n2adj.X > n2adj.Next.X && n2adj.X > n2adj.Prev.X) || (n2adj.X < n2adj.Next.X && n2adj.X < n2adj.Prev.X) {
				diffNextPrev := abs(n2adj.Next.X - n2adj.Prev.X)
				diffBoth := abs(n2adj.X-n2adj.Next.X) + abs(n2adj.X-n2adj.Prev.X)
				areaCount := (diffBoth - diffNextPrev) / 2
				fmt.Printf("Offcut area: %v\n", areaCount)
				totalArea += areaCount

			}
			remove = append(remove, n2adj)
			n2adj.Next.Prev = n2adj.Prev
			n2adj.Prev.Next = n2adj.Next
		}

		fmt.Printf("%v\n", totalArea)

		remainingNodes := make([]*Node, 0)
		for _, n := range nodes {
			keep := true
			for _, r := range remove {
				if r == n {
					keep = false
				}
			}

			if keep {
				remainingNodes = append(remainingNodes, n)
			}
		}
		nodes = remainingNodes

		//if totalArea == 48081 {
		//}
		//
		printNodes(nodes)
		//time.Sleep(500 * time.Millisecond)
	}

	// Expect 62

	return totalArea + 1
}
