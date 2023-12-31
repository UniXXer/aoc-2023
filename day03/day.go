package day03

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"unixxer-aoc-2023/commons"
)

func solveFirst(inputs []string) error {
	r, _ := regexp.Compile(`\d+`)
	sumOfParts := 0

	for i, line := range inputs {
		numbers := r.FindAllStringIndex(line, -1)

		for _, number := range numbers {

			if len(number) > 2 {
				return errors.New("Wrong Number found!!!")
			}

			adjacent, _, _ := isAdjacent(inputs, i, number[0], number[1])

			if adjacent {
				part, err := strconv.Atoi(string(line[number[0]:number[1]]))

				if err != nil {
					return err
				}

				sumOfParts += part
			}
		}
	}

	commons.PrintInfoFormat("Sum of all parts is %d", sumOfParts)

	return nil

}

func solveSecond(inputs []string) error {

	r, _ := regexp.Compile(`\d+`)
	sumOfGears := 0

	possibleGears := make(map[string][]int)

	for i, line := range inputs {
		numbers := r.FindAllStringIndex(line, -1)

		for _, number := range numbers {

			if len(number) > 2 {
				return errors.New("Wrong Number found!!!")
			}

			adjacent, x, y := isAdjacent(inputs, i, number[0], number[1])

			if adjacent {
				if string(inputs[y][x]) == "*" {
					part, err := strconv.Atoi(string(line[number[0]:number[1]]))

					if err != nil {
						return err
					}

					gearKoordString := fmt.Sprintf("x%dy%d", x, y)
					
					possibleGears[gearKoordString] = append(possibleGears[gearKoordString], part)
				}
			}
		}
	}

	for possibleGear := range possibleGears {
		parts := possibleGears[possibleGear]

		if len(parts) == 2 {
			sumOfGears += parts[0] * parts[1]
		}
	}

	commons.PrintInfoFormat("Sum of all gears is %d", sumOfGears)

	return nil

}

func isAdjacent(inputs []string, y int, startX int, endX int) (bool, int, int) {

	//above
	if y > 0 {
		line := inputs[y-1]
		for x := startX - 1; x <= endX; x++ {
			if x > -1 && x < len(line) {
				if hasSymbol(string(line[x])) {
					return true, x, y - 1
				}
			}
		}
	}

	// left
	if startX > 0 {
		if hasSymbol(string(inputs[y][startX-1])) {
			return true, startX - 1, y
		}
	}

	// right
	if endX < len(inputs[y])-1 {
		if hasSymbol(string(inputs[y][endX])) {
			return true, endX, y
		}
	}

	//below
	if y < len(inputs)-1 {
		line := []rune(inputs[y+1])
		for x := startX - 1; x <= endX; x++ {
			if x > -1 && x < len(line) {
				if hasSymbol(string(line[x])) {
					return true, x, y + 1
				}
			}
		}
	}

	return false, -1, -1
}

func hasSymbol(s string) bool {

	_, err := strconv.Atoi(s)

	return err != nil && s != "."

}
