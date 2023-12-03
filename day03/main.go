package main

import (
	"fmt"
	"unicode"

	"github.com/dmytzo/adventofcode2023/tool"
)

func taskOne(input []string) {
	var res int

	for lineIdx, line := range input {
		for idx, c := range line {
			if !unicode.IsOneOf([]*unicode.RangeTable{unicode.Symbol, unicode.Punct}, c) || c == '.' {
				continue
			}

			for _, n := range getRelatedRawNums(idx, lineIdx, input) {
				res += tool.MustInt(n)
			}
		}
	}

	fmt.Println(res)
}

func taskTwo(input []string) {
	var res int

	for lineIdx, line := range input {
		for idx, c := range line {
			if c != '*' {
				continue
			}

			if rawNums := getRelatedRawNums(idx, lineIdx, input); len(rawNums) == 2 {
				ratio := 1
				for _, n := range rawNums {
					ratio *= tool.MustInt(n)
				}

				res += ratio
			}
		}
	}

	fmt.Println(res)
}

func getRelatedRawNums(idx, lineIdx int, input []string) []string {
	inputLen := len(input)

	lines := []string{input[lineIdx]}

	if lineIdx != 0 {
		lines = append(lines, input[lineIdx-1])
	}

	if lineIdx != inputLen-1 {
		lines = append(lines, input[lineIdx+1])
	}

	var rawNums []string

	lineLen := len(input[0])

	startIdx := idx
	if idx != 0 {
		startIdx--
	}

	endIdx := idx
	if idx != lineLen-1 {
		endIdx++
	}

	for _, line := range lines {
		for currentIdx := startIdx; currentIdx <= endIdx; currentIdx++ {
			v := rune(line[currentIdx])

			if !unicode.IsDigit(v) {
				continue
			}

			if currentIdx > startIdx && unicode.IsDigit(rune(line[currentIdx-1])) {
				continue
			}

			startNumIdx, endNumIdx := currentIdx, currentIdx

			for startNumIdx != 0 && unicode.IsDigit(rune(line[startNumIdx-1])) {
				startNumIdx--
			}

			for endNumIdx != lineLen-1 && unicode.IsDigit(rune(line[endNumIdx+1])) {
				endNumIdx++
			}

			rawNums = append(rawNums, line[startNumIdx:endNumIdx+1])
		}
	}

	return rawNums
}

func main() {
	input := tool.InputLines("./input.txt")

	taskOne(input)
	taskTwo(input)
}
