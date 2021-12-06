package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partOne() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	inputArr := strings.Split(strings.TrimSpace(string(input)), ",")
	fish := make([]int, 0, 10000)
	for _, i := range inputArr {
		f, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		fish = append(fish, f)
	}

	for day := 0; day < 80; day++ {
		l := len(fish)
		for i := 0; i < l; i++ {
			if fish[i] == 0 {
				fish[i] = 6
				fish = append(fish, 8)
			} else {
				fish[i]--
			}
		}
	}

	fmt.Printf("Part 1: %d\n", len(fish))
}

func partTwo() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	inputArr := strings.Split(strings.TrimSpace(string(input)), ",")
	fish := make([]int, 0, 10000)
	for _, i := range inputArr {
		f, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		fish = append(fish, f)
	}

	var numFish int64 = int64(len(fish))
	ready := make(map[int]int64)
	for _, f := range fish {
		ready[f]++
	}

	for day := 0; day < 256; day++ {
		r := ready[day]
		numFish += r
		ready[day+7] += r
		ready[day+9] += r
	}

	fmt.Printf("Part 2: %d\n", numFish)
}

func main() {
	partOne()
	partTwo()
}
