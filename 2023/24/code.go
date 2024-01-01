package main

import (
	"errors"
	"fmt"
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

type Hailstone struct {
	X int
	Y int
	Z int

	Xv int
	Yv int
	Zv int
}

func (h *Hailstone) DoPathsCrossXY(o *Hailstone, boundMin, boundMax float64) (x, y float64, err error) {

	// Find which bounds this hailstone is going to
	hyBound, hxBound := float64(h.Y), float64(h.X)
	if h.Yv < 0 {
		hyBound = boundMin
	} else if h.Yv > 0 {
		hyBound = boundMax
	}

	if h.Xv < 0 {
		hxBound = boundMin
	} else if h.Xv > 0 {
		hxBound = boundMax
	}

	// which one reached first

	/*

		19, 13, _ @ -2,  1, _

		y=19 -2c
		x=13 + c





	*/

}

func (h *Hailstone) FindIntersectionXY(o *Hailstone) (ns, x, y float64, err error) {
	if h == o {
		return -1, -1, -1, errors.New("same hailstone")
	}

	yc := (float64(h.Y) - float64(o.Y)) / (float64(o.Yv) - float64(h.Yv))
	xc := (float64(h.X) - float64(o.X)) / (float64(o.Xv) - float64(h.Xv))

	// if xInt == yInt is the same, we have a match

	fmt.Printf("%v ; %v\n", xc, yc)

	if yc == xc && yc > 0 {
		return xc, float64(h.X) + float64(h.Xv)*xc, float64(h.Y) + float64(h.Yv)*yc, nil
	}

	return -1, -1, -1, errors.New("No intersect")

	// y = ax+b        Y =
	/*
		19, 13, _ @ -2,  1, _
		18, 19, _ @ -1, -1, _

		1:
		y= 19 - 2c
		x= 13 +  c


		hY = h.Y + c*h.Yv
		oY = o.Y + c.o.Yv

		if hY=oY:

		h.Y + c*h.Yv = o.Y + c*o.Yv

		h.Y - o.Y = c*o.Yv - c*h.Yv

		h.Y - o.Y = c(o.Yv - h.Yv)
		(h.Y - o.Y) / (o.Yv - h.Yv) =

		2:
		y= 18 - c
		x= 19 - c

		--- ?

		where does y1 = y2
		19 - 2c = 18 - c
		19 - 18 - 2c = -c
		1 = -c + 2c
		1 = c
		@ 1


	*/
}

func run(part2 bool, input string) any {
	hailRxp := regexp.MustCompile(`(?P<X>\d+), +(?P<Y>\d+), +(?P<Z>\d+) @ +(?P<VX>-?\d+), +(?P<VY>-?\d+), +(?P<VZ>-?\d+)`)

	stones := make([]Hailstone, 0)

	for _, l := range strings.Split(input, "\n") {
		matches := hailRxp.FindStringSubmatch(l)
		if len(matches) == 0 {
			continue
		}

		px, _ := strconv.Atoi(matches[1])
		py, _ := strconv.Atoi(matches[2])
		pz, _ := strconv.Atoi(matches[3])
		vx, _ := strconv.Atoi(matches[4])
		vy, _ := strconv.Atoi(matches[5])
		vz, _ := strconv.Atoi(matches[6])

		stones = append(stones, Hailstone{px, py, pz, vx, vy, vz})
	}

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here

	testMin, testMax := 7, 27

	insideRange := 0
	for _, h := range stones {
		for _, o := range stones {
			ns, x, y, err := h.FindIntersectionXY(&o)

			if err == nil {
				fmt.Printf("Intersection of %v:%v at %vns (%v, %v)\n", h, o, ns, x, y)
				if x > float64(testMin) && x < float64(testMax) && y > float64(testMin) && y < float64(testMax) {
					insideRange += 1
				}
			} else {
				fmt.Printf("ERROR %v\n", err)
			}
		}
	}

	return insideRange
}
