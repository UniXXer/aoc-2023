package day02

import (
	"strconv"
	"strings"
	"unixxer-aoc-2023/commons"
)

func solveFirst(inputs []string) error {

	max := make(map[string]int)

	max["red"] = 12
	max["green"] = 13
	max["blue"] = 14

	result := 0

	for i, line := range inputs {
		normalLine := line[strings.Index(line, ": ")+2 : len(line)]

		gamePossible, err := isGamePossible(strings.TrimSpace(normalLine), max)

		if err != nil {
			return err
		}

		if gamePossible {
			result += (i + 1)
		}
	}

	commons.PrintInfoFormat("Sum of all possible games is %d", result)

	return nil

}

func solveSecond(inputs []string) error {

	result := 0

	for _, line := range inputs {
		normalLine := line[strings.Index(line, ": ")+2 : len(line)]

		power, err := powerOfGame(strings.TrimSpace(normalLine))

		if err != nil {
			return err
		}

		result += power
	}

	commons.PrintInfoFormat("Sum of all powers is %d", result)

	return nil
}

func isGamePossible(line string, max map[string]int) (bool, error) {

	games := strings.Split(line, ";")

	for _, round := range games {
		colors := strings.Split(strings.TrimSpace(round), ",")
		for _, color := range colors {
			colorParts := strings.Split(strings.TrimSpace(color), " ")
			counter, err := strconv.Atoi(strings.TrimSpace(colorParts[0]))

			if err != nil {
				return false, err
			}

			colorType := strings.TrimSpace(colorParts[1])

			if max[colorType] < counter {
				return false, nil
			}

		}
	}

	return true, nil
}

func powerOfGame(line string) (int, error) {

	maxPerColor := make(map[string]int)

	maxPerColor["red"] = 0
	maxPerColor["green"] = 0
	maxPerColor["blue"] = 0

	games := strings.Split(line, ";")

	for _, round := range games {
		colors := strings.Split(strings.TrimSpace(round), ",")
		for _, color := range colors {
			colorParts := strings.Split(strings.TrimSpace(color), " ")
			counter, err := strconv.Atoi(strings.TrimSpace(colorParts[0]))
			colorName := strings.TrimSpace(colorParts[1])

			if err != nil {
				return 0, err
			}

			if maxPerColor[colorName] < counter {
				maxPerColor[colorName] = counter
			}
		}
	}

	powerOfSet := maxPerColor["red"] * maxPerColor["green"] * maxPerColor["blue"]

	return powerOfSet, nil
}
