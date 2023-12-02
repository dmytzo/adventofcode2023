package tool

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("os open: %s", err.Error())
	}

	fileContent, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("io read all: %s", err.Error())
	}

	return fileContent
}

func InputLines(path string) []string {
	return strings.Split(string(ReadFile(path)), "\n")
}

func MustInt(i string) int {
	num, err := strconv.Atoi(i)
	if err != nil {
		log.Fatalf("atoi: %s", err.Error())
	}

	return num
}
