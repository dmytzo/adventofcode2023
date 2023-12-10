package main

import (
	"fmt"

	"github.com/dmytzo/adventofcode2023/tool"
)

const (
	startPoint = 'S'

	verticalPipe   = '|'
	horizontalPipe = '-'
	upRightPipe    = 'L'
	upLeftPipe     = 'J'
	downRightPipe  = 'F'
	downLeftPipe   = '7'
)

var (
	possibleUpPipes = map[rune]bool{
		verticalPipe:  true,
		downLeftPipe:  true,
		downRightPipe: true,
	}
	possibleDownPipes = map[rune]bool{
		verticalPipe: true,
		upLeftPipe:   true,
		upRightPipe:  true,
	}
	possibleLeftPipes = map[rune]bool{
		horizontalPipe: true,
		upRightPipe:    true,
		downRightPipe:  true,
	}
	possibleRightPipes = map[rune]bool{
		horizontalPipe: true,
		upLeftPipe:     true,
		downLeftPipe:   true,
	}
)

func taskOne(input []string) {
	var (
		res int

		startCoords    = getStartPoint(input)
		startPointPipe = defineStartPoint(startCoords, input)
	)

	input[startCoords.Row] = input[startCoords.Row][:startCoords.Col] + startPointPipe + input[startCoords.Row][startCoords.Col+1:]

	coords := startCoords
	prevCoords := startCoords

	for {
		newCoords := move(coords, prevCoords, input)
		prevCoords = coords
		coords = newCoords

		res++

		if coords == startCoords {
			break
		}
	}

	fmt.Println(res / 2)
}

func taskTwo(input []string) {
	var (
		res int

		startCoords = getStartPoint(input)
	)

	input[startCoords.Row] = input[startCoords.Row][:startCoords.Col] + string(verticalPipe) + input[startCoords.Row][startCoords.Col+1:]

	coords := startCoords
	prevCoords := startCoords
	pipePoints := map[tool.Coord]bool{}
	for {
		newCoords := move(coords, prevCoords, input)
		prevCoords = coords
		coords = newCoords

		pipePoints[coords] = true
		if coords == startCoords {
			break
		}
	}

	for rowIdx, line := range input {
		var inside bool

		lastPoint := '-'

		for colIdx, c := range line {
			switch {
			case !pipePoints[tool.Coord{Row: rowIdx, Col: colIdx}]:
				if inside {
					res++
				}

				continue
			case c == horizontalPipe:
				continue
			case c == verticalPipe, c == downRightPipe, c == upRightPipe:
				inside = !inside
				lastPoint = c
			case c == upLeftPipe && lastPoint == upRightPipe:
				inside = !inside
				lastPoint = horizontalPipe
			case c == downLeftPipe && lastPoint == downRightPipe:
				inside = !inside
				lastPoint = horizontalPipe
			}
		}
	}

	fmt.Println(res)
}

func getStartPoint(input []string) tool.Coord {
	for lineIdx, line := range input {
		for pointIdx, point := range line {
			if point == startPoint {
				return tool.Coord{Col: pointIdx, Row: lineIdx}
			}
		}
	}

	return tool.Coord{}
}

func defineStartPoint(coords tool.Coord, input []string) string {
	var leftPipeExist, rightPipeExist, upPipeExist, downPipeExist bool

	if coords.Row != len(input)-1 {
		downPipeExist = possibleDownPipes[rune(input[coords.Row+1][coords.Col])]
	}

	if coords.Row != 0 {
		upPipeExist = possibleUpPipes[rune(input[coords.Row-1][coords.Col])]
	}

	if coords.Col != 0 {
		leftPipeExist = possibleLeftPipes[rune(input[coords.Row][coords.Col-1])]
	}

	if coords.Col != len(input[0])-1 {
		rightPipeExist = possibleRightPipes[rune(input[coords.Row][coords.Col+1])]
	}

	var symbol rune

	switch {
	case upPipeExist && downPipeExist:
		symbol = verticalPipe
	case upPipeExist && rightPipeExist:
		symbol = upRightPipe
	case upPipeExist && leftPipeExist:
		symbol = upLeftPipe
	case downPipeExist && rightPipeExist:
		symbol = downRightPipe
	case downPipeExist && leftPipeExist:
		symbol = downLeftPipe
	case leftPipeExist && rightPipeExist:
		symbol = horizontalPipe
	}

	return string(symbol)

}

func move(coords, prevCoords tool.Coord, input []string) tool.Coord {
	v := input[coords.Row][coords.Col]
	switch v {
	case verticalPipe:
		if prevCoords.Row > coords.Row {
			return tool.Coord{Col: coords.Col, Row: coords.Row - 1}
		}

		return tool.Coord{Col: coords.Col, Row: coords.Row + 1}
	case horizontalPipe:
		if prevCoords.Col > coords.Col {
			return tool.Coord{Col: coords.Col - 1, Row: coords.Row}
		}

		return tool.Coord{Col: coords.Col + 1, Row: coords.Row}
	case upRightPipe:
		if prevCoords.Row == coords.Row {
			return tool.Coord{Col: coords.Col, Row: coords.Row - 1}
		}

		return tool.Coord{Col: coords.Col + 1, Row: coords.Row}
	case upLeftPipe:
		if prevCoords.Row == coords.Row {
			return tool.Coord{Col: coords.Col, Row: coords.Row - 1}
		}

		return tool.Coord{Col: coords.Col - 1, Row: coords.Row}
	case downRightPipe:
		if prevCoords.Row == coords.Row {
			return tool.Coord{Col: coords.Col, Row: coords.Row + 1}
		}

		return tool.Coord{Col: coords.Col + 1, Row: coords.Row}
	case downLeftPipe:
		if prevCoords.Row == coords.Row {
			return tool.Coord{Col: coords.Col, Row: coords.Row + 1}
		}

		return tool.Coord{Col: coords.Col - 1, Row: coords.Row}
	}

	return tool.Coord{}
}

func main() {
	input := tool.InputLines("./input.txt")

	input1 := make([]string, len(input))
	input2 := make([]string, len(input))

	copy(input1, input)
	copy(input2, input)

	taskOne(input1)
	taskTwo(input2)
}
