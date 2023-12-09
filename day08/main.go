package main

import (
	"fmt"
	"strings"

	"github.com/dmytzo/adventofcode2023/tool"
)

var instructionToNodeIndex = map[rune]int{
	'L': 0,
	'R': 1,
}

func taskOne(input []string) {
	var (
		currentNode  = "AAA"
		finalNode    = "ZZZ"
		instructions = input[0]
		nodeMap      = map[string][2]string{}

		res int
	)

	for _, line := range input[2:] {
		parts := strings.Split(line, " = ")
		nodes := strings.Split(parts[1][1:len(parts[1])-1], ", ")
		nodeMap[parts[0]] = [2]string(nodes)
	}

	for {
		for _, instruction := range instructions {
			res++

			currentNode = nodeMap[currentNode][instructionToNodeIndex[instruction]]

			if currentNode == finalNode {
				fmt.Println(res)
				return
			}
		}
	}
}

func taskTwo(input []string) {
	var (
		instructions = input[0]
		nodeMap      = map[string][2]string{}

		startNodes []string
		steps      []int
	)

	for _, line := range input[2:] {
		parts := strings.Split(line, " = ")
		nodes := strings.Split(parts[1][1:len(parts[1])-1], ", ")
		nodeMap[parts[0]] = [2]string(nodes)

		if parts[0][2] == 'A' {
			startNodes = append(startNodes, parts[0])
		}
	}

	for _, node := range startNodes {
		var (
			currentNode = node

			res  int
			done bool
		)

		for !done {
			for _, instruction := range instructions {
				res++

				currentNode = nodeMap[currentNode][instructionToNodeIndex[instruction]]

				if currentNode[2] == 'Z' {
					steps = append(steps, res)
					done = true
					break
				}
			}
		}
	}

	fmt.Println(lcm(steps[0], steps[1], steps[2:]...))
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	res := a * b / gcd(a, b)

	for _, i := range integers {
		res = lcm(res, i)
	}

	return res
}

func main() {
	input := tool.InputLines("./input.txt")
	taskOne(input)
	taskTwo(input)
}
