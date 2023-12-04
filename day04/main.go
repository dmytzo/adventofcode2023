package main

import (
	"fmt"
	"strings"

	"github.com/dmytzo/adventofcode2023/tool"
)

func taskOne(input []string) {
	var res int

	for _, line := range input {
		var points int

		parts := strings.Split(strings.Split(line, ":")[1], "|")

		winningNumsMap := make(map[string]bool)
		for _, num := range strings.Fields(parts[0]) {
			winningNumsMap[num] = true
		}

		for _, num := range strings.Fields(parts[1]) {
			if !winningNumsMap[num] {
				continue
			}

			if points == 0 {
				points = 1
				continue
			}

			points *= 2
		}

		res += points
	}

	fmt.Println(res)
}

func taskTwo(input []string) {
	var res int

	cards := make(map[int]int)

	for _, line := range input {
		parts := strings.Split(line, ":")
		cardNum := tool.MustInt(strings.Fields(parts[0])[1])
		numParts := strings.Split(parts[1], "|")

		winningNumsMap := make(map[int]bool)
		for _, rawNum := range strings.Fields(numParts[0]) {
			winningNumsMap[tool.MustInt(rawNum)] = true
		}

		cards[cardNum]++
		res++

		nextCard := cardNum + 1
		for _, rawNum := range strings.Fields(numParts[1]) {
			num := tool.MustInt(rawNum)
			if !winningNumsMap[num] {
				continue
			}

			for i := 0; i < cards[cardNum]; i++ {
				cards[nextCard]++
				res++
			}

			nextCard++
		}
	}

	fmt.Println(res)
}

func main() {
	input := tool.InputLines("./input.txt")

	taskOne(input)
	taskTwo(input)
}
