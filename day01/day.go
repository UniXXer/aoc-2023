package day01

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unixxer-aoc-2023/commons"
)

type digit struct {
	name  string
	value string
}

var allDigitsExtended = []digit{
	{name: "one", value: "one1one"},
	{name: "two", value: "two2two"},
	{name: "three", value: "three3three"},
	{name: "four", value: "four4four"},
	{name: "five", value: "five5five"},
	{name: "six", value: "six6six"},
	{name: "seven", value: "seven7seven"},
	{name: "eight", value: "eight8eight"},
	{name: "nine", value: "nine9nine"},
}

func solveFirst(inputs []string) error {

	resultCounter, err := calculateCalibration(inputs)

	if err != nil {
		return err
	}

	commons.PrintInfoFormat("Sum of calibration values is %d", resultCounter)

	return nil
}

func solveSecond(inputs []string) error {

	normalizeInputs(inputs)

	resultCounter, err := calculateCalibration(inputs)

	if err != nil {
		return err
	}

	commons.PrintInfoFormat("Sum of calibration values is %d", resultCounter)

	return nil
}

func normalizeInputs(inputs []string) {
	for i := 0; i < len(inputs); i++ {
		iterStr := inputs[i]

		for _, d := range allDigitsExtended {
			iterStr = strings.Replace(iterStr, d.name, d.value, -1)
		}

		inputs[i] = iterStr
	}
}

func calculateCalibration(inputs []string) (resultCounter int, e error) {
	resultCounter = 0

	for i, s := range inputs {

		firstIdx, firstNumber := findFirstNumber(s)
		lastIdx, lastNumber := findLastNumber(s)

		if firstIdx == -1 || lastIdx == -1 {
			e = errors.New(fmt.Sprintf("Either the firstIdx or the lastIdx is -1 in line %d! [%d, %s] [%d, %s]",
				i, firstIdx, firstNumber, lastIdx, lastNumber))
		}

		lineCounter, err := strconv.Atoi(firstNumber + lastNumber)

		if err != nil {
			e = err
		}

		resultCounter += lineCounter
	}

	return
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
