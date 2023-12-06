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

	winMargin := calculateWinMargin(times, dists)

	commons.PrintInfoFormat("Total win margin is %d", winMargin)

	return nil

}

func solveSecond(inputs []string) error {

	times := parseLine(inputs[0])
	dists := parseLine(inputs[1])

	time, _ := strconv.Atoi(strings.Trim(strings.Replace(fmt.Sprint(times), " ", "", -1), "[]"))
	dist, _ := strconv.Atoi(strings.Trim(strings.Replace(fmt.Sprint(dists), " ", "", -1), "[]"))

	winMargin := calculateWinMargin([]int{time}, []int{dist})

	commons.PrintInfoFormat("Total win margin is %d", winMargin)

	return nil

}

func calculateWinMargin(times, dists []int) int {
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

	return winMargin
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
