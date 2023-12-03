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

			if isAdjacent(inputs, i, number[0], number[1]) {
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

	fmt.Printf("solveSecond is not implemented yet!")
	fmt.Println()

	return nil

}

func isAdjacent(inputs []string, y int, startX int, endX int) bool {

	//above
	if y > 0 {
		line := inputs[y-1]
		for x := startX - 1; x <= endX; x++ {
			if x > -1 && x < len(line) {
				if hasSymbol(string(line[x])) {
					return true
				}
			}
		}
	}

	// left
	if startX > 0 {
		if hasSymbol(string(inputs[y][startX-1])) {
			return true
		}
	}

	// right
	if endX < len(inputs[y])-1 {
		if hasSymbol(string(inputs[y][endX])) {
			return true
		}
	}

	//below
	if y < len(inputs)-1 {
		line := []rune(inputs[y+1])
		for x := startX - 1; x <= endX; x++ {
			if x > -1 && x < len(line) {
				if hasSymbol(string(line[x])) {
					return true
				}
			}
		}
	}

	return false
}

func hasSymbol(s string) bool {

	_, err := strconv.Atoi(s)

	return err != nil && s != "."

}
