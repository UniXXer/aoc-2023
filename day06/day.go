package day06

import (
	"fmt"
	"strconv"
	"strings"
	"unixxer-aoc-2023/commons"
)

func solveFirst(inputs []string) error {

	times := parseLine(inputs[0])
	dists := parseLine(inputs[1])

	winMargin := 1

	for race := 0; race < len(times); race++ {
		wins := 0
		for buttonTime := 1; buttonTime <= times[race]; buttonTime++ {
			myDistance := (times[race] - buttonTime) * buttonTime

			if myDistance > dists[race] {
				wins++
			}
		}

		winMargin = winMargin * wins
	}

	commons.PrintInfoFormat("Total win margin is %d", winMargin)

	return nil

}

func solveSecond(inputs []string) error {

	fmt.Printf("solveSecond is not implemented yet!")
	fmt.Println()

	return nil

}

func parseLine(line string) []int {
	splits := strings.Split(line, " ")
	numbers := []int{}

	for _, split := range splits {
		s, err := strconv.Atoi(split)
		if err == nil {
			numbers = append(numbers, s)
		}
	}

	return numbers
}
