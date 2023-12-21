package main

import (
	"fmt"
	"math"
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

type Module struct {
	Name   string
	Prefix rune

	OutputsStr []string
	Outputs    []*Module

	// For %
	On bool

	// For &
	ConjMap map[*Module]bool
}

type Pulse struct {
	From        *Module
	Destination *Module
	High        bool
}

func (m *Module) LowPulse(from *Module) []Pulse {
	if m.Name == "rx" {
		panic("WE GOT HERE")
	}

	next := make([]Pulse, 0)

	if m.Prefix == '%' {
		m.On = !m.On
		for _, n := range m.Outputs {
			next = append(next, Pulse{m, n, m.On})
		}
	} else if m.Prefix == '&' {
		m.ConjMap[from] = false

		for _, n := range m.Outputs {
			next = append(next, Pulse{m, n, true})
		}
	} else if m.Prefix == 'b' {
		for _, n := range m.Outputs {
			next = append(next, Pulse{m, n, false})
		}
	} else if m.Prefix == '.' {
		// Output, do nothing
	} else {
		panic("bad prefix")
	}

	return next
}

func (m *Module) HiPulse(from *Module) []Pulse {
	next := make([]Pulse, 0)

	if m.Prefix == '%' {
		// Ignore
	} else if m.Prefix == '&' {
		m.ConjMap[from] = true

		anyLow := false
		for _, val := range m.ConjMap {
			if !val {
				anyLow = true
			}
		}

		for _, n := range m.Outputs {
			next = append(next, Pulse{m, n, anyLow})
		}
	} else if m.Prefix == 'b' {
		for _, n := range m.Outputs {
			next = append(next, Pulse{m, n, true})
		}
	} else if m.Prefix == '.' {
		// Output, do nothing
	} else {
		panic("bad prefix")
	}

	return next

}

func (p *Pulse) Pulse(tally *Tally) ([]Pulse, bool) {
	//fmt.Printf("Pulsing %v from %v to %v [L:%v H:%v] \n", p.High, p.From, p.Destination, tally.Lo, tally.Hi)
	if p.High {
		tally.Hi += 1
		return p.Destination.HiPulse(p.From), false
	}

	if p.Destination.Name == "rx" {
		fmt.Printf("Rx! Terminating!")
		return nil, true
	}

	tally.Lo += 1
	return p.Destination.LowPulse(p.From), false
}

type Tally struct {
	Hi int
	Lo int
}

func run(part2 bool, input string) any {
	var broadcaster Module

	mods := make(map[string]*Module)

	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}

		spl := strings.Split(l, " -> ")

		dests := strings.Split(spl[1], ", ")

		if spl[0] == "broadcaster" {
			broadcaster = Module{
				"broadcaster",
				'b',
				dests,
				make([]*Module, 0),
				false,
				make(map[*Module]bool),
			}

			mods["broadcaster"] = &broadcaster
		} else {
			mod := Module{
				spl[0][1:],
				rune(spl[0][0]),
				dests,
				make([]*Module, 0),
				false,
				make(map[*Module]bool),
			}
			mods[mod.Name] = &mod
		}
	}

	// Now, configure
	toProcess := []*Module{&broadcaster}
	processed := make(map[string]bool)
	for len(toProcess) > 0 {
		m := toProcess[0]
		toProcess = toProcess[1:]

		if _, ok := processed[m.Name]; ok {
			continue
		}
		processed[m.Name] = true

		//fmt.Printf("Preprocessing %v\n", m.Name)

		for _, d := range m.OutputsStr {
			dest, ok := mods[d]
			if !ok {
				dest = &Module{
					d,
					'.',
					make([]string, 0),
					make([]*Module, 0),
					false,
					make(map[*Module]bool),
				}
			}
			m.Outputs = append(m.Outputs, dest)
			dest.ConjMap[m] = false
		}
		toProcess = append(toProcess, m.Outputs...)
	}

	// Print mappings:
	modPrint := []*Module{&broadcaster}
	processed = make(map[string]bool)
	for len(modPrint) > 0 {
		mp := modPrint[0]
		modPrint = modPrint[1:]
		if _, ok := processed[mp.Name]; ok {
			continue
		}
		processed[mp.Name] = true

		modPrint = append(modPrint, mp.Outputs...)

		//fmt.Printf("%v (%v): %v (%v) [%v]\n", mp.Name, mp.Prefix, mp.OutputsStr, mp.Outputs, &mp)
	}

	tally := Tally{0, 0}

	limit := 1000
	if part2 {
		limit = math.MaxInt
	}

	for i := 0; i < limit; i++ {
		if i%1000000 == 0 {
			fmt.Printf("Sent %v pulses\n", i)
		}
		pulsesToSend := []Pulse{{nil, &broadcaster, false}}
		for len(pulsesToSend) > 0 {
			pulse := pulsesToSend[0]
			pulsesToSend = pulsesToSend[1:]

			newPulses, rxLow := pulse.Pulse(&tally)

			if rxLow && part2 {
				return tally.Lo
			}

			//fmt.Printf("New pulses: %v\n", newPulses)

			pulsesToSend = append(pulsesToSend, newPulses...)
		}
	}

	fmt.Printf("Lo: %v ; Hi: %v\n", tally.Lo, tally.Hi)

	return tally.Lo * tally.Hi
}
