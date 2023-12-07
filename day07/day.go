package day07

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
	"unixxer-aoc-2023/commons"
)

type Hand struct {
	original 	string
	bid 		int
	sort 		string
}

var cardMapping = map[string]string{
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

func LineToHand(line string) *Hand {
	var hand Hand
	splits := strings.Split(line, " ")
	hand.original = strings.TrimSpace(splits[0])
	hand.bid, _ = strconv.Atoi(strings.TrimSpace(splits[1]))
	hand.sort = calcSortHash(&hand)
	return &hand
}

func calcSortHash(hand *Hand) string {
	charCounter := make(map[string]int)
	for i := 0; i < len(hand.original); i++ {
		var c = string(hand.original[i])
		charCounter[c]++
		hand.sort += cardMapping[string(c)]
	}

	hand.sort = determineHandType(charCounter) + hand.sort

	return hand.sort
}

func determineHandType(chars map[string]int) string {
	values := make([]int, 0, len(chars))
	for k := range chars {
		values = append(values, chars[k])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	switch values[0] {
		case 5: return "A" //Five of a kind
		case 4: return "B" //Four of a kind
		case 3: {
			switch values[1] {
				case 2 : return "C" // Full house
				default : return "D" // Three of a kind
			}
		}
		case 2: {
			switch values[1] {
				case 2 : return "E" // Two pair
				default : return "F" // One pair 
			}
		}
		default : return "G" // High card
	}
}

var AllHands []Hand = []Hand{}

func solveFirst(inputs []string) error {

	for _, line := range inputs {

		AllHands = append(AllHands, *LineToHand(line))
	}

	slices.SortFunc(AllHands, func (a, b Hand) int {
		return strings.Compare(a.sort, b.sort) * -1
	})

	totalWinnings := 0

	for rank, hand := range AllHands {
		totalWinnings += (rank+1) * hand.bid
	}

	commons.PrintInfoFormat("Total Winnings are %d", totalWinnings)

	return nil

}

func solveSecond(inputs []string) error {

	fmt.Printf("solveSecond is not implemented yet!")
	fmt.Println()

	return nil

}
