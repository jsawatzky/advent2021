package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func partOne() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	inputArr := strings.Split(strings.TrimSpace(string(input)), ",")
	crabs := make([]int, 0, len(inputArr))
	for _, i := range inputArr {
		c, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		crabs = append(crabs, c)
	}

	sort.Ints(crabs)
	median := crabs[len(crabs)/2]

	totalFuel := 0
	for _, c := range crabs {
		if c > median {
			totalFuel += c - median
		} else {
			totalFuel += median - c
		}
	}

	fmt.Printf("Part 1: %d\n", totalFuel)
}

func partTwo() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	inputArr := strings.Split(strings.TrimSpace(string(input)), ",")
	crabs := make([]int, 0, len(inputArr))
	for _, i := range inputArr {
		c, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		crabs = append(crabs, c)
	}

	sum := 0
	for _, c := range crabs {
		sum += c
	}
	avg := sum / len(crabs)

	totalFuel := 0
	for _, c := range crabs {
		var move int
		if c > avg {
			move = c - avg
		} else {
			move = avg - c
		}
		totalFuel += (move * (move + 1)) / 2
	}

	fmt.Printf("Part 2: %d\n", totalFuel)
}

func main() {
	partOne()
	partTwo()
}
