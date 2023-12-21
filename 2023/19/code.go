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
type Part struct {
	X int
	M int
	A int
	S int
}

type PartAcceptor struct {
	Wfl  []Workflow
	Xmax int
	Xmin int
	Mmax int
	Mmin int
	Amax int
	Amin int
	Smax int
	Smin int
}

type Workflow struct {
	Comparing   string
	CompareType string
	Value       int
	Result      string
	Method      func(p Part) string
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	workflows := make(map[string][]Workflow)
	parts := make([]Part, 0)

	inParts := false
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			inParts = true
			continue
		}

		if !inParts {
			cutoff := strings.Index(l, "{")
			name := l[0:cutoff]
			wkflow := l[cutoff+1 : len(l)-1]

			functions := make([]Workflow, 0)
			for _, wkF := range strings.Split(wkflow, ",") {

				if strings.Contains(wkF, ":") {
					// Mapping function
					spl := strings.Split(wkF, ":")

					comparing := spl[0][0:2]
					compared, _ := strconv.Atoi(spl[0][2:])
					res := spl[1]

					if comparing == "x>" {
						functions = append(functions, Workflow{"x", ">", compared, res, func(p Part) string {
							if p.X > compared {
								return res
							}
							return ""
						}})
					} else if comparing == "x<" {
						functions = append(functions, Workflow{"x", "<", compared, res, func(p Part) string {
							if p.X < compared {
								return res
							}
							return ""
						}})
					} else if comparing == "m>" {
						functions = append(functions, Workflow{"m", ">", compared, res, func(p Part) string {
							if p.M > compared {
								return res
							}
							return ""
						}})
					} else if comparing == "m<" {
						functions = append(functions, Workflow{"m", "<", compared, res, func(p Part) string {
							if p.M < compared {
								return res
							}
							return ""
						}})
					} else if comparing == "a>" {
						functions = append(functions, Workflow{"a", ">", compared, res, func(p Part) string {
							if p.A > compared {
								return res
							}
							return ""
						}})
					} else if comparing == "a<" {
						functions = append(functions, Workflow{"a", "<", compared, res, func(p Part) string {
							if p.A < compared {
								return res
							}
							return ""
						}})
					} else if comparing == "s>" {
						functions = append(functions, Workflow{"s", ">", compared, res, func(p Part) string {
							if p.S > compared {
								return res
							}
							return ""
						}})
					} else if comparing == "s<" {
						functions = append(functions, Workflow{"s", "<", compared, res, func(p Part) string {
							if p.S < compared {
								return res
							}
							return ""
						}})
					} else {
						panic("!??")
					}

				} else {
					// Just a return function
					functions = append(functions, Workflow{"", "", 0, wkF, func(p Part) string { return wkF }})
				}
			}

			workflows[name] = functions
		} else {
			l = strings.Replace(strings.Replace(l, "{", "", -1), "}", "", -1)

			spl := strings.Split(l, ",")

			x, _ := strconv.Atoi(spl[0][2:])
			m, _ := strconv.Atoi(spl[1][2:])
			a, _ := strconv.Atoi(spl[2][2:])
			s, _ := strconv.Atoi(spl[3][2:])

			parts = append(parts, Part{x, m, a, s})
		}
	}

	if part2 {
		pa := PartAcceptor{workflows["in"], 4000, 1, 4000, 1, 4000, 1, 4000, 1}

		pas := []PartAcceptor{pa}

		combos := int64(0)
		for len(pas) > 0 {
			//fmt.Printf("%v\n", pas)
			p := pas[0]
			pas = pas[1:]

			for _, wf := range p.Wfl {
				if wf.CompareType != "" {
					//fmt.Printf("Splitting comparator\n")
					// Comparator! Will split
					if wf.CompareType == ">" {
						if wf.Comparing == "x" && p.Xmax > wf.Value {
							if p.Xmin > wf.Value {
								if wf.Result == "A" {
									fmt.Printf("Accepted: %v\n", p)
									combos += int64(p.Xmax-p.Xmin+1) * int64(p.Mmax-p.Mmin+1) * int64(p.Amax-p.Amin+1) * int64(p.Smax-p.Smin+1)
								} else if wf.Result != "R" {
									newP := PartAcceptor{workflows[wf.Result], p.Xmax, p.Xmin, p.Mmax, p.Mmin, p.Amax, p.Amin, p.Smax, p.Smin}
									pas = append(pas, newP)
								}
							} else { // Split
								if wf.Result == "A" {
									fmt.Printf("Accepted (split): %v\n", p)
									combos += int64(p.Xmax-(wf.Value+1)+1) * int64(p.Mmax-p.Mmin+1) * int64(p.Amax-p.Amin+1) * int64(p.Smax-p.Smin+1)
								} else if wf.Result != "R" {
									newP := PartAcceptor{workflows[wf.Result], p.Xmax, p.Xmin, p.Mmax, p.Mmin, p.Amax, p.Amin, p.Smax, p.Smin}
									pas = append(pas, newP)
								}

								p.Xmax = wf.Value
							}
						} else if wf.Comparing == "m" && p.Mmax > wf.Value {
							if p.Mmin > wf.Value {
								if wf.Result == "A" {
									fmt.Printf("Accepted: %v\n", p)
									combos += int64(p.Xmax-p.Xmin+1) * int64(p.Mmax-p.Mmin+1) * int64(p.Amax-p.Amin+1) * int64(p.Smax-p.Smin+1)
								} else if wf.Result != "R" {
									newP := PartAcceptor{workflows[wf.Result], p.Xmax, p.Xmin, p.Mmax, p.Mmin, p.Amax, p.Amin, p.Smax, p.Smin}
									pas = append(pas, newP)
								}
							} else { // Split
								if wf.Result == "A" {
									fmt.Printf("Accepted: %v\n", p)
									combos += int64(p.Xmax-p.Xmin+1) * int64(p.Mmax-(wf.Value+1)+1) * int64(p.Amax-p.Amin+1) * int64(p.Smax-p.Smin+1)
								} else if wf.Result != "R" {
									newP := PartAcceptor{workflows[wf.Result], p.Xmax, p.Xmin, p.Mmax, p.Mmin, p.Amax, p.Amin, p.Smax, p.Smin}
									pas = append(pas, newP)
								}

								p.Mmax = wf.Value
							}
						} else if wf.Comparing == "a" && p.Amax > wf.Value {
							if p.Amin > wf.Value {
								if wf.Result == "A" {
									fmt.Printf("Accepted: %v\n", p)
									combos += int64(p.Xmax-p.Xmin+1) * int64(p.Mmax-p.Mmin+1) * int64(p.Amax-p.Amin+1) * int64(p.Smax-p.Smin+1)
								} else if wf.Result != "R" {
									newP := PartAcceptor{workflows[wf.Result], p.Xmax, p.Xmin, p.Mmax, p.Mmin, p.Amax, p.Amin, p.Smax, p.Smin}
									pas = append(pas, newP)
								}
							} else { // Split
								if wf.Result == "A" {
									fmt.Printf("Accepted: %v\n", p)
									combos += int64(p.Xmax-p.Xmin+1) * int64(p.Mmax-p.Mmin+1) * int64(p.Amax-(wf.Value+1)+1) * int64(p.Smax-p.Smin+1)
								} else if wf.Result != "R" {
									newP := PartAcceptor{workflows[wf.Result], p.Xmax, p.Xmin, p.Mmax, p.Mmin, p.Amax, p.Amin, p.Smax, p.Smin}
									pas = append(pas, newP)
								}

								p.Amax = wf.Value
							}
						} else if wf.Comparing == "s" && p.Smax > wf.Value {
							if p.Smin > wf.Value {
								if wf.Result == "A" {
									fmt.Printf("Accepted: %v\n", p)
									combos += int64(p.Xmax-p.Xmin+1) * int64(p.Mmax-p.Mmin+1) * int64(p.Amax-p.Amin+1) * int64(p.Smax-p.Smin+1)
								} else if wf.Result != "R" {
									newP := PartAcceptor{workflows[wf.Result], p.Xmax, p.Xmin, p.Mmax, p.Mmin, p.Amax, p.Amin, p.Smax, p.Smin}
									pas = append(pas, newP)
								}
							} else { // Split
								if wf.Result == "A" {
									fmt.Printf("Accepted: %v\n", p)
									combos += int64(p.Xmax-p.Xmin+1) * int64(p.Mmax-p.Mmin+1) * int64(p.Amax-p.Amin+1) * int64(p.Smax-p.Smin+1)
								} else if wf.Result != "R" {
									newP := PartAcceptor{workflows[wf.Result], p.Xmax, p.Xmin, p.Mmax, p.Mmin, p.Amax, p.Amin, p.Smax, wf.Value + 1}
									pas = append(pas, newP)
								}

								p.Smax = wf.Value
							}
						}
					} else if wf.CompareType == "<" {
						if wf.Comparing == "x" && p.Xmin < wf.Value {
							if p.Xmax < wf.Value {
								if wf.Result == "A" {
									fmt.Printf("Accepted: %v\n", p)
									combos += int64(p.Xmax-p.Xmin+1) * int64(p.Mmax-p.Mmin+1) * int64(p.Amax-p.Amin+1) * int64(p.Smax-p.Smin+1)
								} else if wf.Result != "R" {
									newP := PartAcceptor{workflows[wf.Result], p.Xmax, p.Xmin, p.Mmax, p.Mmin, p.Amax, p.Amin, p.Smax, p.Smin}
									pas = append(pas, newP)
								}
							} else { // Split
								if wf.Result == "A" {
									fmt.Printf("Accepted (split): %v\n", p)
									combos += int64((wf.Value-1)-p.Xmin+1) * int64(p.Mmax-p.Mmin+1) * int64(p.Amax-p.Amin+1) * int64(p.Smax-p.Smin+1)
								} else if wf.Result != "R" {
									newP := PartAcceptor{workflows[wf.Result], p.Xmax, p.Xmin, p.Mmax, p.Mmin, p.Amax, p.Amin, p.Smax, p.Smin}
									pas = append(pas, newP)
								}

								p.Xmin = wf.Value
							}
						} else if wf.Comparing == "m" && p.Mmin < wf.Value {
							if p.Mmax < wf.Value {
								if wf.Result == "A" {
									fmt.Printf("Accepted: %v\n", p)
									combos += int64(p.Xmax-p.Xmin+1) * int64(p.Mmax-p.Mmin+1) * int64(p.Amax-p.Amin+1) * int64(p.Smax-p.Smin+1)
								} else if wf.Result != "R" {
									newP := PartAcceptor{workflows[wf.Result], p.Xmax, p.Xmin, p.Mmax, p.Mmin, p.Amax, p.Amin, p.Smax, p.Smin}
									pas = append(pas, newP)
								}
							} else { // Split
								if wf.Result == "A" {
									fmt.Printf("Accepted: %v\n", p)
									combos += int64(p.Xmax-p.Xmin+1) * int64((wf.Value-1)-p.Mmin+1) * int64(p.Amax-p.Amin+1) * int64(p.Smax-p.Smin+1)
								} else if wf.Result != "R" {
									newP := PartAcceptor{workflows[wf.Result], p.Xmax, p.Xmin, p.Mmax, p.Mmin, p.Amax, p.Amin, p.Smax, p.Smin}
									pas = append(pas, newP)
								}

								p.Mmin = wf.Value
							}
						} else if wf.Comparing == "a" && p.Amin < wf.Value {
							if p.Amax < wf.Value {
								if wf.Result == "A" {
									fmt.Printf("Accepted: %v\n", p)
									combos += int64(p.Xmax-p.Xmin+1) * int64(p.Mmax-p.Mmin+1) * int64(p.Amax-p.Amin+1) * int64(p.Smax-p.Smin+1)
								} else if wf.Result != "R" {
									newP := PartAcceptor{workflows[wf.Result], p.Xmax, p.Xmin, p.Mmax, p.Mmin, p.Amax, p.Amin, p.Smax, p.Smin}
									pas = append(pas, newP)
								}
							} else { // Split
								if wf.Result == "A" {
									fmt.Printf("Accepted: %v\n", p)
									combos += int64(p.Xmax-p.Xmin+1) * int64(p.Mmax-p.Mmin+1) * int64((wf.Value-1)-p.Amin+1) * int64(p.Smax-p.Smin+1)
								} else if wf.Result != "R" {
									newP := PartAcceptor{workflows[wf.Result], p.Xmax, p.Xmin, p.Mmax, p.Mmin, p.Amax, p.Amin, p.Smax, p.Smin}
									pas = append(pas, newP)
								}

								p.Amin = wf.Value
							}
						} else if wf.Comparing == "s" && p.Smin < wf.Value {
							if p.Smax < wf.Value {
								if wf.Result == "A" {
									fmt.Printf("Accepted: %v\n", p)
									combos += int64(p.Xmax-p.Xmin+1) * int64(p.Mmax-p.Mmin+1) * int64(p.Amax-p.Amin+1) * int64(p.Smax-p.Smin+1)
								} else if wf.Result != "R" {
									newP := PartAcceptor{workflows[wf.Result], p.Xmax, p.Xmin, p.Mmax, p.Mmin, p.Amax, p.Amin, p.Smax, p.Smin}
									pas = append(pas, newP)
								}
							} else { // Split
								if wf.Result == "A" {
									fmt.Printf("Accepted: %v\n", p)
									combos += int64(p.Xmax-p.Xmin+1) * int64(p.Mmax-p.Mmin+1) * int64(p.Amax-p.Amin+1) * int64((wf.Value-1)-p.Smin+1)
								} else if wf.Result != "R" {
									newP := PartAcceptor{workflows[wf.Result], p.Xmax, p.Xmin, p.Mmax, p.Mmin, p.Amax, p.Amin, p.Smax, p.Smin}
									pas = append(pas, newP)
								}

								p.Smin = wf.Value
							}
						}
					} else {
						panic("!?!?")
					}

				} else if wf.Result == "A" {
					// Auto accept
					fmt.Printf("Accepted: %v\n", p)
					combos += int64(p.Xmax-p.Xmin+1) * int64(p.Mmax-p.Mmin+1) * int64(p.Amax-p.Amin+1) * int64(p.Smax-p.Smin+1)
				} else if wf.Result == "R" {
					// Do nothing, rejected
					fmt.Printf("Rejected: %v\n", p)
				} else {
					//fmt.Printf("Moved to %v (%v)\n", wf.Result, p)
					p.Wfl = workflows[wf.Result]
					pas = append(pas, p)
				}
			}
		}

		// Expect:  167409079868000
		// Getting: 397530093888000

		return combos

	}

	ratingSum := 0
	for _, p := range parts {
		wk := workflows["in"]

		for wk != nil {
		workflowOne:
			for _, fnc := range wk {
				res := fnc.Method(p)

				if res != "" {
					if res == "A" {
						ratingSum += (p.X + p.M + p.A + p.S)
						wk = nil
					} else if res == "R" {
						wk = nil
					} else {
						// it's a mapping
						wk = workflows[res]
					}
					break workflowOne
				}
			}
		}
	}

	// solve part 1 here
	return ratingSum
}
