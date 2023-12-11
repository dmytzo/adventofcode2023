package main

import (
	"fmt"

	"github.com/dmytzo/adventofcode2023/tool"
)

func taskOne(input []string) {
	solveWithExtender(input, 1)
}

func taskTwo(input []string) {
	solveWithExtender(input, 999_999)
}

func solveWithExtender(input []string, extender int) {
	var coords []tool.Coord

	notEmptyCols := map[int]bool{}
	notEmptyRows := map[int]bool{}

	for lineIdx, line := range input {
		emptyRow := true
		for colIdx, col := range line {
			if col != '.' {
				coords = append(coords, tool.Coord{Row: lineIdx, Col: colIdx})

				notEmptyCols[colIdx] = true
				emptyRow = false
			}
		}

		if !emptyRow {
			notEmptyRows[lineIdx] = true
		}
	}

	var res int
	pathMap := map[tool.Coord]map[tool.Coord]bool{}

	for _, c := range coords {
		for rowIdx := c.Row; rowIdx < len(input); rowIdx++ {
			for colIdx, col := range input[rowIdx] {
				if rowIdx == c.Row && colIdx < c.Col {
					continue
				}

				currentCoord := tool.Coord{Row: rowIdx, Col: colIdx}

				if col == '.' || pathMap[c][currentCoord] {
					continue
				}

				res += currentCoord.Row - c.Row
				for i := c.Row; i <= currentCoord.Row; i++ {
					if !notEmptyRows[i] {
						res += extender
					}
				}

				colIdxFrom, colIdxTo := c.Col, currentCoord.Col
				if colIdxFrom > colIdxTo {
					colIdxFrom, colIdxTo = colIdxTo, colIdxFrom
				}

				res += colIdxTo - colIdxFrom
				for i := colIdxFrom; i <= colIdxTo; i++ {
					if !notEmptyCols[i] {
						res += extender
					}
				}

				if pathMap[c] == nil {
					pathMap[c] = map[tool.Coord]bool{}
				}

				pathMap[c][currentCoord] = true
			}
		}
	}

	fmt.Println(res)
}

func main() {
	input := tool.InputLines("./input.txt")

	taskOne(input)
	taskTwo(input)
}
