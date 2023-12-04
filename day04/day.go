package day04

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"unixxer-aoc-2023/commons"
)

func solveFirst(inputs []string) error {

	var sumOfCardPoints float64 = 0

	for _, line := range inputs {
		winningNumbers, myNumbers := parseLine(line)

		allWinningNumbers := "-" + strings.Join(winningNumbers, "-") + "-"

		hitCounter := -1

		for i := range myNumbers {
			searchStr := "-" + myNumbers[i] + "-"
			idx := strings.Index(allWinningNumbers, searchStr)

			if idx > -1 {
				hitCounter++
			}
		}

		if hitCounter > -1 {
			cardPoints := math.Pow(2, float64(hitCounter))
			sumOfCardPoints += cardPoints
		}
	}

	commons.PrintInfoFormat("Sum of all card points is %.0f", sumOfCardPoints)

	return nil

}

func solveSecond(inputs []string) error {

	cardPoints := make(map[int]int)

	allCards := []int{}

	for j, line := range inputs {
		allCards = append(allCards, j+1)
		winningNumbers, myNumbers := parseLine(line)

		allWinningNumbers := "-" + strings.Join(winningNumbers, "-") + "-"

		for i := range myNumbers {
			searchStr := "-" + myNumbers[i] + "-"
			idx := strings.Index(allWinningNumbers, searchStr)

			if idx > -1 {
				cardPoints[j+1]++
			}
		}
	}

	pos := 0

	for pos < len(allCards) {
		currPoints := cardPoints[allCards[pos]]

		var cardsToAdd []int

		for k := 1; k <= currPoints; k++ {
			cardsToAdd = append(cardsToAdd, allCards[pos+k])
		}

		commons.PrintDebug(fmt.Sprintf("The card %d at %d has %d points. Therefore adding %s", []any{allCards[pos], pos, currPoints, cardsToAdd}...))

		for l := 0; l < len(cardsToAdd); l++ {
			addCard := cardsToAdd[l]
			allCards = insertIntoArray(allCards, addCard, pos+l)
		}

		fmt.Println(allCards)

		pos++

		if pos == 30 {
			break
		}
	}

	fmt.Println(allCards)
	fmt.Println(cardPoints)

	return nil

}

func parseLine(line string) ([]string, []string) {
	r, _ := regexp.Compile(`\d+`)

	winningNumbersStr := string(line[strings.Index(line, ":"):strings.Index(line, "|")])
	myNumbersStr := string(line[strings.Index(line, "|"):])

	winningNumbers := r.FindAllString(winningNumbersStr, -1)
	myNumbers := r.FindAllString(myNumbersStr, -1)

	return winningNumbers, myNumbers
}

func insertIntoArray(array []int, value int, idx int) []int {
	array = append(array, 0)
	copy(array[idx+1:], array[idx:])
	array[idx] = value
	return array
}
