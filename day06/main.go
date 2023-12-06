package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/dmytzo/adventofcode2023/tool"
)

func taskOne(input []string) {
	var res int64 = 1

	timeInput := tool.MustIntSlice(strings.Fields(strings.Split(input[0], ":")[1]))
	distanceInput := tool.MustIntSlice(strings.Fields(strings.Split(input[1], ":")[1]))

	for i := 0; i < len(timeInput); i++ {
		x1, y1 := solveEquation(float64(timeInput[i]), float64(distanceInput[i]))

		res *= (x1 - y1)
	}

	fmt.Println(res)
}

func taskTwo(input []string) {
	timeInput := tool.MustInt(strings.ReplaceAll(strings.Split(input[0], ":")[1], " ", ""))
	distanceInput := tool.MustInt(strings.ReplaceAll(strings.Split(input[1], ":")[1], " ", ""))

	x1, y1 := solveEquation(float64(timeInput), float64(distanceInput))

	fmt.Println(x1 - y1)
}

func solveEquation(b, c float64) (int64, int64) {
	// ax^2 - bx + c = 0

	d := -b*-b - 4*c

	x1 := (b + math.Sqrt(d)) / 2
	y1 := b - x1

	x1Rounded := int64(x1)
	y1Rounded := int64(y1)

	if x1 == float64(x1Rounded) {
		x1Rounded--
	}

	return x1Rounded, y1Rounded
}

func main() {
	input := tool.InputLines("./input.txt")
	taskOne(input)
	taskTwo(input)
}
