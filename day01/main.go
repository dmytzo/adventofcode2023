package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/dmytzo/adventofcode2023/tool"
)

var letterDidgitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func taskOne(input []string) {
	var res int

	for _, v := range input {
		var firstNumRaw, lastNumRaw string
		for _, c := range v {
			if unicode.IsDigit(c) {
				if firstNumRaw == "" {
					firstNumRaw = string(c)
				}

				lastNumRaw = string(c)
			}
		}

		num, err := strconv.Atoi(firstNumRaw + lastNumRaw)
		if err != nil {
			log.Fatalf("atoi: %s", err.Error())
		}

		res += num
	}

	fmt.Println(res)
}

func taskTwo(input []string) {
	var res int

	for _, line := range input {
		var firstNumRaw, lastNumRaw string

		firstNumRawIdx := len(line)
		lastNumRawIdx := 0

		for lettersNum, digitNum := range letterDidgitMap {
			idx := strings.Index(line, lettersNum)
			if idx != -1 {
				if idx <= firstNumRawIdx {
					firstNumRawIdx = idx
					firstNumRaw = digitNum
				}

				idx = strings.LastIndex(line, lettersNum)
				if idx >= lastNumRawIdx {
					lastNumRawIdx = idx
					lastNumRaw = digitNum
				}
			}

			idx = strings.Index(line, digitNum)
			if idx != -1 {
				if idx <= firstNumRawIdx {
					firstNumRawIdx = idx
					firstNumRaw = digitNum
				}

				idx = strings.LastIndex(line, digitNum)
				if idx >= lastNumRawIdx {
					lastNumRawIdx = idx
					lastNumRaw = digitNum
				}
			}
		}

		res += tool.MustInt(firstNumRaw + lastNumRaw)
	}

	fmt.Println(res)
}

func main() {
	input := tool.InputLines("./input.txt")

	taskOne(input)
	taskTwo(input)
}
