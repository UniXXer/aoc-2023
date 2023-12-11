package day08

import (
	"fmt"
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

	for(currentNode.name != "ZZZ") {
		var instruction string

		if(currPos < len(instructions)) {
			instruction = string(instructions[currPos])

			switch (instruction) {
				case "L" : currentNode = nodes[currentNode.left]
				case "R" : currentNode = nodes[currentNode.right]
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

	fmt.Printf("solveSecond is not implemented yet!")
	fmt.Println()

	return nil

}
