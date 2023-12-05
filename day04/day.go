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
		allCards = append(allCards, 1)
		winningNumbers, myNumbers := parseLine(line)

		allWinningNumbers := "-" + strings.Join(winningNumbers, "-") + "-"

		for i := range myNumbers {
			searchStr := "-" + myNumbers[i] + "-"
			idx := strings.Index(allWinningNumbers, searchStr)

			if idx > -1 {
				cardPoints[j]++
			}
		}
	}

	for i := 0; i < len(allCards); i++ {
		currPoints := cardPoints[i]

		fmt.Println("allCards: ", allCards)

		commons.PrintDebugFormat("Card %d has %d points.", i, currPoints)

		for j := i+1; j <= i+currPoints; j++ {
			if(j >= len(allCards)) {
				continue
			}

			allCards[j] += allCards[i]
		}
	} 

	cardCounter := 0

	for i := 0; i < len(allCards); i++ {
		cardCounter += allCards[i]
	}

	commons.PrintInfoFormat("Count of all card is %d", cardCounter)

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
