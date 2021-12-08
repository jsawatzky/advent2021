package helpers

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func InputScanner() *bufio.Scanner {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(file)
}

func ReadInput() string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(input))
}

func ReadInputLines() []string {
	lines := make([]string, 0, 100)
	scanner := InputScanner()
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
