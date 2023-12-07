package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/dmytzo/adventofcode2023/tool"
)

var cardTypes = map[rune]int{
	'A': 23,
	'K': 22,
	'Q': 21,
	'J': 20,
	'T': 19,
	'9': 18,
	'8': 17,
	'7': 16,
	'6': 15,
	'5': 14,
	'4': 13,
	'3': 12,
	'2': 11,
}

var combinationStrengthMap = map[int]int{
	5: 90,
	4: 85,
	3: 55,
	2: 25,
}

func taskOne(input []string) {
	strengthToBid := map[int]int{}

	for _, line := range input {
		var (
			parts                 = strings.Split(line, " ")
			handCards             = map[rune]int{}
			handStrength          = 0
			cardMultiplier        = 100_000_000
			combinationStrength   = 0
			combinationMultiplier = cardMultiplier * 100
		)

		for _, card := range parts[0] {
			handCards[card]++
			handStrength += cardMultiplier * cardTypes[card]
			cardMultiplier /= 100
		}

		for _, v := range handCards {
			if v == 1 {
				continue
			}

			combinationStrength += combinationStrengthMap[v]
		}

		if combinationStrength == 0 {
			combinationStrength = 10
		}

		handStrength += combinationStrength * combinationMultiplier
		strengthToBid[handStrength] = tool.MustInt(parts[1])
	}

	fmt.Println(calculateBid(strengthToBid))
}

func taskTwo(input []string) {
	joker := 'J'
	cardTypesWithJoker := maps.Clone(cardTypes)
	cardTypesWithJoker[joker] = 10

	strengthToBid := map[int]int{}

	for _, line := range input {
		var (
			parts                 = strings.Split(line, " ")
			handCards             = map[rune]int{}
			handStrength          = 0
			cardMultiplier        = 100_000_000
			combinationStrength   = 0
			combinationMultiplier = cardMultiplier * 100
		)

		for _, card := range parts[0] {
			handCards[card]++
			handStrength += cardMultiplier * cardTypesWithJoker[card]
			cardMultiplier /= 100
		}

		var (
			maxCardNumKey rune
			maxCardNum    int
		)

		if jokerNum := handCards['J']; jokerNum != 0 {
			for k, v := range handCards {
				if k == joker {
					continue
				}

				if v > maxCardNum {
					maxCardNumKey = k
					maxCardNum = v
				}
			}

			if maxCardNum != 0 {
				handCards[maxCardNumKey] += jokerNum
				handCards[joker] = 0
			}
		}

		for _, v := range handCards {
			if v == 1 {
				continue
			}

			combinationStrength += combinationStrengthMap[v]
		}

		if combinationStrength == 0 {
			combinationStrength = 10
		}

		handStrength += combinationStrength * combinationMultiplier
		strengthToBid[handStrength] = tool.MustInt(parts[1])
	}

	fmt.Println(calculateBid(strengthToBid))
}

func calculateBid(strengthToBid map[int]int) int {
	sortedStrengths := make([]int, 0, len(strengthToBid))
	for s := range strengthToBid {
		sortedStrengths = append(sortedStrengths, s)
	}

	slices.Sort(sortedStrengths)

	var res int
	for idx, k := range sortedStrengths {
		res += (strengthToBid[k] * (idx + 1))
	}

	return res
}

func main() {
	input := tool.InputLines("./input.txt")
	taskOne(input)
	taskTwo(input)
}
