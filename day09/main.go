package main

import (
	"fmt"
	"strings"

	"github.com/dmytzo/adventofcode2023/tool"
)

func taskOne(input []string) {
	var res int
	for _, line := range input {
		var (
			currentLine = tool.MustIntSlice(strings.Fields(line))

			done bool
		)

		for !done {
			done = true

			nextLine := []int{}

			res += currentLine[len(currentLine)-1]

			for i := 0; i < len(currentLine)-1; i++ {
				newNum := currentLine[i+1] - currentLine[i]
				if newNum != 0 {
					done = false
				}

				nextLine = append(nextLine, newNum)
			}

			currentLine = nextLine
		}
	}
	fmt.Println(res)
}

func taskTwo(input []string) {
	var res int
	for _, line := range input {
		var (
			currentLine = tool.MustIntSlice(strings.Fields(line))

			calcNums []int
			done     bool
		)

		for !done {
			done = true

			nextLine := []int{}

			calcNums = append(calcNums, currentLine[0])

			for i := 0; i < len(currentLine)-1; i++ {
				newNum := currentLine[i+1] - currentLine[i]
				if newNum != 0 {
					done = false
				}

				nextLine = append(nextLine, newNum)
			}

			currentLine = nextLine
		}

		var num int
		for i := len(calcNums) - 1; i >= 0; i-- {
			num = calcNums[i] - num
		}

		res += num
	}
	fmt.Println(res)
}

func main() {
	input := tool.InputLines("./input.txt")
	taskOne(input)
	taskTwo(input)
}
