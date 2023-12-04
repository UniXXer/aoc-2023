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

	fmt.Printf("solveSecond is not implemented yet!")
	fmt.Println()

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
