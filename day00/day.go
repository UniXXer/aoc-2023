package day00

import (
	"fmt"
	"strconv"
	"unixxer-aoc-2023/commons"
)

func solveFirst(inputs []string) error {
	var maximum, curr int = 0, 0

	for _, s := range inputs {
		if s != "" {
			val, err := strconv.Atoi(s)

			if err != nil {
				return err
			}

			curr = curr + val
		} else {
			curr = 0
		}

		maximum = max(curr, maximum)

	}

	commons.PrintInfo(fmt.Sprintf("The maximum is %d", maximum))

	return nil
}

func solveSecond(inputs []string) error {

	fmt.Printf("solveSecond is not implemented yet!")

	return nil

}
