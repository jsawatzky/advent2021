package helpers

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

type CloseFunc func()

func InputScanner() (*bufio.Scanner, CloseFunc) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(file), func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}
}

func ReadInput() string {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(input))
}

func ReadInputLines() []string {
	lines := make([]string, 0, 100)
	scanner, close := InputScanner()
	defer close()
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	return lines
}

func ReadCsvInput() []string {
	return strings.Split(ReadInput(), ",")
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func AtoiArr(arr []string) []int {
	ret := make([]int, 0, len(arr))
	for _, s := range arr {
		ret = append(ret, Atoi(s))
	}
	return ret
}
