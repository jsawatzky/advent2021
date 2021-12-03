package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numLines int
	ones := make([]int, 12)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		numLines++
		for i, c := range line {
			if c == '1' {
				ones[i] += 1
			}
		}
	}

	var gamma, epsilon string
	for _, c := range ones {
		if c < numLines-c {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	g, err := strconv.ParseInt(gamma, 2, 32)
	if err != nil {
		panic(err)
	}
	e, err := strconv.ParseInt(epsilon, 2, 32)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", g*e)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numLines := 1000
	lines := make([]string, 0, numLines)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
	}

	var oxygen, co2 string

	var prefix string
	tempLines := lines
	for curBit := 0; curBit < 12; curBit++ {
		var ones int
		for _, l := range tempLines {
			if l[curBit] == '1' {
				ones += 1
			}
		}

		if ones < len(tempLines)-ones {
			prefix += "0"
		} else {
			prefix += "1"
		}
		newLines := make([]string, 0, len(tempLines))
		for _, l := range tempLines {
			if strings.HasPrefix(l, prefix) {
				newLines = append(newLines, l)
			}
		}
		if len(newLines) == 1 {
			oxygen = newLines[0]
			break
		}

		tempLines = newLines
	}

	prefix = ""
	tempLines = lines
	for curBit := 0; curBit < 12; curBit++ {
		var ones int
		for _, l := range tempLines {
			if l[curBit] == '1' {
				ones += 1
			}
		}

		if ones >= len(tempLines)-ones {
			prefix += "0"
		} else {
			prefix += "1"
		}
		newLines := make([]string, 0, len(tempLines))
		for _, l := range tempLines {
			if strings.HasPrefix(l, prefix) {
				newLines = append(newLines, l)
			}
		}
		if len(newLines) == 1 {
			co2 = newLines[0]
			break
		}

		tempLines = newLines
	}

	o, err := strconv.ParseInt(oxygen, 2, 32)
	if err != nil {
		panic(err)
	}
	c, err := strconv.ParseInt(co2, 2, 32)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 2: %d\n", o*c)
}

func main() {
	partOne()
	partTwo()
}
