package day08

import (
	"unixxer-aoc-2023/commons"
)

var instructions string
var nodes map[string]Node = make(map[string]Node)

type Node struct {
	name  string
	left  string
	right string
}

func newNode(nodeStr string) *Node {
	var node Node

	node.name = nodeStr[0:3]
	node.left = nodeStr[7:10]
	node.right = nodeStr[12:15]

	return &node
}

func solveFirst(inputs []string) error {

	instructions = inputs[0]

	for i := 2; i < len(inputs); i++ {
		node := newNode(inputs[i])
		nodes[node.name] = *node
	}

	currentNode := nodes["AAA"]
	steps := 0
	currPos := 0

	for currentNode.name != "ZZZ" {
		var instruction string

		if currPos < len(instructions) {
			instruction = string(instructions[currPos])

			switch instruction {
			case "L":
				currentNode = nodes[currentNode.left]
			case "R":
				currentNode = nodes[currentNode.right]
			}
			currPos++
			steps++
		} else {
			currPos = 0
		}
	}

	commons.PrintInfoFormat("Arrived after %d steps at %s", steps, currentNode.name)

	return nil

}

func solveSecond(inputs []string) error {

	solveFirst(inputs)

	aCounter := 0
	zCounter := 0

	currNodes := []Node{}

	for name, node := range nodes {
		if string(name[2]) == "A" {
			//changeNodeName(aCounter, "A", name)
			currNodes = append(currNodes, node)
			aCounter++
		} else if string(name[2]) == "Z" {
			//changeNodeName(zCounter, "Z", name)
			zCounter++
		}
	}

	commons.PrintDebugFormat("Found %d A-Nodes and %d Z-Nodes", aCounter, zCounter)

	allSteps := []int{}

	for i, currentNode := range currNodes {
		steps := 0
		currPos := 0

		for currentNode.name[2] != 'Z' {
			var instruction string

			if currPos < len(instructions) {
				instruction = string(instructions[currPos])

				switch instruction {
				case "L":
					currentNode = nodes[currentNode.left]
				case "R":
					currentNode = nodes[currentNode.right]
				}
				currPos++
				steps++
			} else {
				currPos = 0
			}
		}

		commons.PrintDebugFormat("Arrived after %d steps at for index %d", steps, i)

		allSteps = append(allSteps, steps)
	}

	lcm := LCM(allSteps[0], allSteps[1], allSteps[2], allSteps[3], allSteps[4], allSteps[5])

	commons.PrintInfoFormat("Will arrive after %d steps", lcm)

	return nil

}

// THANKS TO https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
