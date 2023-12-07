package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Hand struct {
	Cards string
	Bid   int
}

func getCardPower(card rune) int {
	if unicode.IsNumber(card) {
		i, _ := strconv.Atoi(string(card))
		return i
	}

	if card == 'T' {
		return 10
	}

	if card == 'J' {
		return 11
	}

	if card == 'Q' {
		return 12
	}

	if card == 'K' {
		return 13
	}

	if card == 'A' {
		return 14
	}

	panic(fmt.Sprintf("Bad card %v", card))
}

func getCardPowerJokerRules(card rune) int {
	if unicode.IsNumber(card) {
		i, _ := strconv.Atoi(string(card))
		return i
	}

	if card == 'T' {
		return 10
	}

	if card == 'J' {
		return 0
	}

	if card == 'Q' {
		return 12
	}

	if card == 'K' {
		return 13
	}

	if card == 'A' {
		return 14
	}

	panic(fmt.Sprintf("Bad card %v", card))
}

func getHandTypeRanking(hand Hand) int {
	cardUniq := make(map[string]int)

	for _, c := range strings.Split(hand.Cards, "") {
		if _, ok := cardUniq[c]; !ok {
			cardUniq[c] = 0
		}
		cardUniq[c] += 1
	}

	if len(cardUniq) == 1 {
		// Full house, top priority
		return 7
	}

	if len(cardUniq) == 2 {
		// Either four of a kind or full house
		for _, x := range cardUniq {
			if x == 4 || x == 1 {
				// Four of a kind
				return 6
			} else {
				// Must be full house
				return 5
			}
		}
	}

	if len(cardUniq) == 3 {
		// Either three of a kind or two pair
		for _, x := range cardUniq {
			if x == 3 {
				// Three of a kind
				return 4
			} else if x == 2 {
				// Two pair
				return 3
			}
		}
	}

	if len(cardUniq) == 4 {
		// Must be one pair
		return 2
	}

	// Must be high card
	return 1
}
func getHandTypeRankingJokerRules(hand Hand) int {
	cardUniq := make(map[string]int)

	jokerCount := 0
	for _, c := range strings.Split(hand.Cards, "") {
		if _, ok := cardUniq[c]; !ok {
			cardUniq[c] = 0
		}
		cardUniq[c] += 1

		if c == "J" {
			jokerCount += 1
		}
	}

	if len(cardUniq) == 1 {
		// Full house, top priority
		return 7
	}

	if len(cardUniq) == 2 {
		// Either four of a kind or full house
		for _, x := range cardUniq {
			if x == 4 || x == 1 {
				// Four of a kind
				if jokerCount == 4 || jokerCount == 1 {
					// Can be converted to 5oac
					return 7
				}
				return 6
			} else {
				// Must be full house
				if jokerCount == 3 || jokerCount == 2 {
					// Can be converted to 5oac
					return 7
				}
				return 5
			}
		}
	}

	if len(cardUniq) == 3 {
		// Either three of a kind or two pair
		for _, x := range cardUniq {
			if x == 3 {
				// Three of a kind
				if jokerCount == 1 || jokerCount == 3 {
					// Can be converted to 4oac
					return 6
				}
				return 4
			} else if x == 2 {
				// Two pair
				if jokerCount == 2 {
					// Can be converted to 4oac
					return 6
				} else if jokerCount == 1 {
					// Can be converted to full house
					return 5
				}
				return 3
			}
		}
	}

	if len(cardUniq) == 4 {
		// Must be one pair
		if jokerCount == 2 || jokerCount == 1 {
			// Can be converted to 3oac
			return 4
		}
		return 2
	}

	// Must be high card
	if jokerCount == 1 {
		// Converted to one pair
		return 2
	}
	return 1
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	hands := make([]Hand, 0)

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		lspl := strings.Split(line, " ")
		bid, _ := strconv.Atoi(lspl[1])

		hands = append(hands, Hand{lspl[0], bid})
	}

	if part2 {
		sort.SliceStable(hands, func(i, j int) bool {
			// First compare rank
			hRank1 := getHandTypeRankingJokerRules(hands[i])
			hRank2 := getHandTypeRankingJokerRules(hands[j])

			if hRank1 != hRank2 {
				return hRank1 < hRank2
			}

			// Otherwise, compare cards
			for k := 0; k < 5; k++ {
				cRank1 := getCardPowerJokerRules(rune(hands[i].Cards[k]))
				cRank2 := getCardPowerJokerRules(rune(hands[j].Cards[k]))

				if cRank1 != cRank2 {
					return cRank1 < cRank2
				}
			}

			panic(fmt.Sprintf("Could not sort %v and %v", hands[i].Cards, hands[j].Cards))
		})

	} else {
		sort.SliceStable(hands, func(i, j int) bool {
			// First compare rank
			hRank1 := getHandTypeRanking(hands[i])
			hRank2 := getHandTypeRanking(hands[j])

			if hRank1 != hRank2 {
				return hRank1 < hRank2
			}

			// Otherwise, compare cards
			for k := 0; k < 5; k++ {
				cRank1 := getCardPower(rune(hands[i].Cards[k]))
				cRank2 := getCardPower(rune(hands[j].Cards[k]))

				if cRank1 != cRank2 {
					return cRank1 < cRank2
				}
			}

			panic(fmt.Sprintf("Could not sort %v and %v", hands[i].Cards, hands[j].Cards))
		})
	}
	// solve part 1 here

	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * hand.Bid
	}

	return winnings
}
