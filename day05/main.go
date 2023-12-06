package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/dmytzo/adventofcode2023/tool"
)

func taskOne(input []string) {
	transformers := getTransformers(input[2:])

	var seeds []int
	for _, v := range strings.Split(input[0][7:], " ") {
		seeds = append(seeds, tool.MustInt(v))
	}

	minTransformedSeed := seeds[0]
	for _, s := range seeds {
		transformedSeed := s
		for _, transformer := range transformers {

			for _, v := range transformer {
				if transformedSeed >= v[1] && transformedSeed < v[1]+v[2] {
					transformedSeed = v[0] + transformedSeed - v[1]
					break
				}
			}
		}
		if transformedSeed < minTransformedSeed {
			minTransformedSeed = transformedSeed
		}
	}

	fmt.Println(minTransformedSeed)
}

func taskTwo(input []string) {
	transformers := getTransformers(input[2:])
	transformersLen := len(transformers)

	var seeds []int
	for _, v := range strings.Split(input[0][7:], " ") {
		seeds = append(seeds, tool.MustInt(v))
	}

	seedsLen := len(seeds)

	location := 0

	for {
		newSeed := location
		for j := transformersLen - 1; j >= 0; j-- {
			for _, v := range transformers[j] {
				if newSeed >= v[0] && newSeed < v[0]+v[2] {
					newSeed = newSeed - v[0] + v[1]
					break
				}
			}
		}

		for i := 0; i < seedsLen; i += 2 {
			if newSeed >= seeds[i] && newSeed < seeds[i]+seeds[i+1] {
				fmt.Println(location)
				return
			}
		}

		location++
	}
}

func getTransformers(input []string) [][][3]int {
	transformers := [][][3]int{{}}

	transformersIdx := 0
	for _, line := range input {
		if line == "" {
			transformersIdx++
			transformers = append(transformers, [][3]int{})
			continue
		}

		if !unicode.IsDigit(rune(line[0])) {
			continue
		}

		parts := strings.Split(line, " ")

		transformers[transformersIdx] = append(
			transformers[transformersIdx],
			[3]int{tool.MustInt(parts[0]), tool.MustInt(parts[1]), tool.MustInt(parts[2])},
		)
	}

	return transformers
}

func main() {
	input := tool.InputLines("./input.txt")

	taskOne(input)
	taskTwo(input)
}
