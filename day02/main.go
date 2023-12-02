package main

import (
	"fmt"
	"regexp"

	"github.com/dmytzo/adventofcode2023/tool"
)

var (
	cubes = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	recordPattern = regexp.MustCompile("(\\d+) (green|red|blue)")
	idPattern     = regexp.MustCompile("Game (\\d+)")
)

func taskOne(input []string) {
	var res int

	for _, line := range input {
		var skipLine bool

		for _, record := range recordPattern.FindAllStringSubmatch(line, -1) {
			if tool.MustInt(record[1]) > cubes[record[2]] {
				skipLine = true
				break
			}
		}

		if !skipLine {
			res += tool.MustInt(idPattern.FindStringSubmatch(line)[1])
		}
	}

	fmt.Println(res)
}

func taskTwo(input []string) {
	var res int

	for _, line := range input {
		maxNums := make(map[string]int)
		for _, record := range recordPattern.FindAllStringSubmatch(line, -1) {
			num := tool.MustInt(record[1])

			if num > maxNums[record[2]] {
				maxNums[record[2]] = num
			}
		}

		recordPower := 1
		for _, v := range maxNums {
			recordPower *= v
		}

		res += recordPower
	}

	fmt.Println(res)
}

func main() {
	input := tool.InputLines("./input.txt")

	taskOne(input)
	taskTwo(input)
}
