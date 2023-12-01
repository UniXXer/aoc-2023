package day01

import (
	"errors"
	"fmt"
	"strconv"
	"unixxer-aoc-2023/commons"
)

func solveFirst(inputs []string) error {

	resultCounter := 0

	for i, s := range inputs {

		firstIdx, firstNumber := findFirstNumber(s)
		lastIdx, lastNumber := findLastNumber(s)

		if firstIdx == -1 || lastIdx == -1 {
			return errors.New(fmt.Sprintf("Either the firstIdx or the lastIdx is -1 in line %d! [%d, %s] [%d, %s]",
				i, firstIdx, firstNumber, lastIdx, lastNumber))
		}

		lineCounter, err := strconv.Atoi(firstNumber + lastNumber)

		if err != nil {
			return err
		}
		commons.PrintDebugFormat("Calibration value for line %d is %d", i, lineCounter)

		resultCounter += lineCounter
	}

	commons.PrintInfoFormat("Sum of calibration values is %d", resultCounter)

	return nil
}

func solveSecond(inputs []string) error {

	fmt.Printf("solveSecond is not implemented yet!")
	fmt.Println()

	return nil
}

func findFirstNumber(input string) (idx int, number string) {
	slice := []rune(input)
	idx = -1

	for i, s := range slice {
		if _, err := strconv.Atoi(string(s)); err == nil {
			idx = i
			number = string(s)
			break
		}
	}

	return
}

func findLastNumber(input string) (idx int, number string) {
	slice := []rune(input)
	idx = -1

	for i := len(slice) - 1; i >= 0; i-- {
		if _, err := strconv.Atoi(string(slice[i])); err == nil {
			return i, string(slice[i])
		}
	}

	return -1, ""
}
