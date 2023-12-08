package day07

import (
	"slices"
	"sort"
	"strconv"
	"strings"
	"unixxer-aoc-2023/commons"
)

type Hand struct {
	original string
	bid      int
	sort     string
}

var withJokers bool = false

var cardMapping_Part1 = map[string]string{
	"A": "A",
	"K": "B",
	"Q": "C",
	"J": "D",
	"T": "E",
	"9": "F",
	"8": "G",
	"7": "H",
	"6": "I",
	"5": "J",
	"4": "K",
	"3": "L",
	"2": "M",
}

var cardMapping_Part2 = map[string]string{
	"A": "A",
	"K": "B",
	"Q": "C",
	"T": "D",
	"9": "E",
	"8": "F",
	"7": "G",
	"6": "H",
	"5": "I",
	"4": "J",
	"3": "K",
	"2": "L",
	"J": "M",
}

func LineToHand(line string, cardMapper map[string]string) *Hand {
	var hand Hand
	splits := strings.Split(line, " ")
	hand.original = strings.TrimSpace(splits[0])
	hand.bid, _ = strconv.Atoi(strings.TrimSpace(splits[1]))
	hand.sort = calcSortHash(&hand, cardMapper)
	return &hand
}

func calcSortHash(hand *Hand, cardMapper map[string]string) string {
	charCounter := make(map[string]int)

	for i := 0; i < len(hand.original); i++ {
		var c = string(hand.original[i])
		charCounter[c]++
		hand.sort += cardMapper[string(c)]
	}

	hand.sort = determineHandType(charCounter) + hand.sort

	return hand.sort
}

func determineHandType(chars map[string]int) string {
	values := make([]int, 0, len(chars))
	jokers := 0
	for k := range chars {
		if withJokers && k == "J" {
			jokers = chars[k]
		} else {
			values = append(values, chars[k])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	if !withJokers {
		return typeOfHand(values)
	} else {
		jokersLeft := jokers

		if jokers >= 4 {
			return "A"
		}
		
		for i := jokersLeft; i > 0; i-- {
			for j := len(values); j > 0; j-- {
				if values[j-1] < 5 {
					values[j-1]++
				}
			}
		}
	

		/*
		if jokers == 3 {
			if values[0] == 2 {
				return "A"
			}
			
			return "B"
		}

		if jokers == 2 {
			if values[0] == 3 {
				return "A"
			}

			if values[0] == 2 {
				return "B"
			}
		}

		if jokers == 1 {
			if values[0] == 4 {
				return "A"
			}

			if values[0] == 3 {
				return "B"
			}

			if values[0] == 2 && values[1] == 2 {
				return "C"
			}

			if values[0] == 2 {
				return "D"
			}

			if values[0] == 1 {
				return "F"
			}
		}*/

		return typeOfHand(values)
	}
}

func typeOfHand(values []int) string {
	switch values[0] {
	case 5:
		return "A" //Five of a kind
	case 4:
		return "B" //Four of a kind
	case 3:
		{
			switch values[1] {
			case 2:
				return "C" // Full house
			default:
				return "D" // Three of a kind
			}
		}
	case 2:
		{
			switch values[1] {
			case 2:
				return "E" // Two pair
			default:
				return "F" // One pair
			}
		}
	default:
		return "G" // High card
	}
}

var AllHands []Hand

func solveFirst(inputs []string) error {

	withJokers = false
	AllHands = []Hand{}

	for _, line := range inputs {
		AllHands = append(AllHands, *LineToHand(line, cardMapping_Part1))
	}

	slices.SortFunc(AllHands, func(a, b Hand) int {
		return strings.Compare(a.sort, b.sort) * -1
	})

	totalWinnings := 0

	for rank, hand := range AllHands {
		totalWinnings += (rank + 1) * hand.bid
	}

	commons.PrintInfoFormat("Total Winnings are %d", totalWinnings)

	return nil
}

func solveSecond(inputs []string) error {

	withJokers = true
	AllHands := []Hand{}

	for _, line := range inputs {
		AllHands = append(AllHands, *LineToHand(line, cardMapping_Part2))
	}

	slices.SortFunc(AllHands, func(a, b Hand) int {
		return strings.Compare(a.sort, b.sort) * -1
	})

	totalWinnings := 0

	for rank, hand := range AllHands {
		totalWinnings += (rank + 1) * hand.bid
	}

	commons.PrintInfoFormat("Total Winnings are %d", totalWinnings)

	return nil
}
