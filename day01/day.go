package day01

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unixxer-aoc-2023/commons"
)

type digit struct {
	name  string
	value string
}

var allDigits = []digit{
	{name: "one", value: "1"},
	{name: "two", value: "2"},
	{name: "three", value: "3"},
	{name: "four", value: "4"},
	{name: "five", value: "5"},
	{name: "six", value: "6"},
	{name: "seven", value: "7"},
	{name: "eight", value: "8"},
	{name: "nine", value: "9"},
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
	var d digit
	for i := 0; i < len(inputs); i++ {
		iterStr := inputs[i]

		d = findNext(iterStr)

		if d.value == "" {
			continue
		}

		iterStr = strings.Replace(iterStr, d.name, d.value, 1)

		for d.name != "" {
			d = findNext(iterStr)
			iterStr = strings.Replace(iterStr, d.name, d.value, 1)
		}

		//commons.PrintDebug(inputs[i] + " => " + iterStr)

		inputs[i] = iterStr
	}
}

func findNext(input string) digit {
	var digitFound digit
	lastPos := math.MaxInt64

	for _, d := range allDigits {
		curPos := strings.Index(input, d.name)
		if curPos > -1 && curPos < lastPos {
			lastPos = curPos
			digitFound = d
		}
	}

	return digitFound
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
