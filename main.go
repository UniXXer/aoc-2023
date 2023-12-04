package main

import (
	_ "aoc2023/day00"
	_ "aoc2023/day01"
	_ "aoc2023/day02"
	_ "aoc2023/day03"
	_ "aoc2023/day04"
	
	"errors"
	"os"
	"unixxer-aoc-2023/commons"
)

func main() {
	if err := solveDay(os.Args); err != nil {
		commons.PrintErr(err.Error())
		os.Exit(1)
	}
}

func solveDay(args []string) error {
	argsLen := len(args)
	if argsLen < 2 {
		return errors.New("Please enter a day to solve.")
	}

	i := args[1]

	exampleMode := (argsLen > 2 && args[2] == "example")

	errFirst := commons.AllDays[i].SolveFirst(exampleMode)

	if errFirst != nil {
		return errFirst
	}

	errSecond := commons.AllDays[i].SolveSecond(exampleMode)

	if errSecond != nil {
		return errSecond
	}

	return nil
}
