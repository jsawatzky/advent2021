package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jsawatzky/advent/helpers"
)

func readInput() (string, map[string]string) {
	lines := helpers.ReadInputLines()
	start := lines[0]
	subs := make(map[string]string)
	lines = lines[2:]
	for _, l := range lines {
		form := strings.Split(l, " -> ")
		subs[form[0]] = form[1]
	}

	return start, subs
}

func partOne() {
	template, subs := readInput()

	for i := 0; i < 10; i++ {
		newTemplate := ""
		for j := 1; j < len(template); j++ {
			pair := template[j-1 : j+1]
			if s, ok := subs[pair]; ok {
				newTemplate += pair[:1] + s
			} else {
				newTemplate += pair[:1]
			}
		}
		template = newTemplate + template[len(template)-1:]
	}

	counts := make(map[rune]int)
	for _, c := range template {
		counts[c]++
	}

	min := math.MaxInt32
	max := 0

	for _, v := range counts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	ans := max - min

	fmt.Printf("Part 1: %d\n", ans)
}

func partTwo() {
	template, subs := readInput()

	pairCounts := make(map[string]int64)
	for j := 1; j < len(template); j++ {
		pair := template[j-1 : j+1]
		pairCounts[pair]++
	}

	for i := 0; i < 40; i++ {
		newPairCounts := make(map[string]int64)
		for p, v := range pairCounts {
			if s, ok := subs[p]; ok {
				newPairCounts[p[:1]+s] += v
				newPairCounts[s+p[1:]] += v
			} else {
				newPairCounts[p] += v
			}
		}
		pairCounts = newPairCounts
	}

	counts := make(map[rune]int64)
	for p, v := range pairCounts {
		for _, c := range p {
			counts[c] += v
		}
	}
	tRunes := []rune(template)
	counts[tRunes[0]]++
	counts[tRunes[len(tRunes)-1]]++

	var min int64 = math.MaxInt64
	var max int64

	for _, v := range counts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	ans := max/2 - min/2

	fmt.Printf("Part 2: %d\n", ans)
}

func main() {
	partOne()
	partTwo()
}
