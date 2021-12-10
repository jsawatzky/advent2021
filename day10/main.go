package main

import (
	"fmt"
	"sort"

	"github.com/jsawatzky/advent/helpers"
)

func partOne() {
	lines := helpers.ReadInputLines()

	ans := 0

	for _, l := range lines {
		stack := make([]rune, 0, 10)
	loop:
		for _, c := range l {
			var match rune
			switch c {
			case '(':
				fallthrough
			case '[':
				fallthrough
			case '{':
				fallthrough
			case '<':
				stack = append(stack, c)
				continue
			}
			match, stack = stack[len(stack)-1], stack[:len(stack)-1]
			switch c {
			case ')':
				if match != '(' {
					ans += 3
					break loop
				}
			case ']':
				if match != '[' {
					ans += 57
					break loop
				}
			case '}':
				if match != '{' {
					ans += 1197
					break loop
				}
			case '>':
				if match != '<' {
					ans += 25137
					break loop
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", ans)
}

func partTwo() {
	lines := helpers.ReadInputLines()

	scores := make([]int, 0, 10)

outerLoop:
	for _, l := range lines {
		stack := make([]rune, 0, 10)
		for _, c := range l {
			var match rune
			switch c {
			case '(':
				fallthrough
			case '[':
				fallthrough
			case '{':
				fallthrough
			case '<':
				stack = append(stack, c)
				continue
			}
			if len(stack) == 0 {
				continue
			}
			match, stack = stack[len(stack)-1], stack[:len(stack)-1]
			switch c {
			case ')':
				if match != '(' {
					continue outerLoop
				}
			case ']':
				if match != '[' {
					continue outerLoop
				}
			case '}':
				if match != '{' {
					continue outerLoop
				}
			case '>':
				if match != '<' {
					continue outerLoop
				}
			}
		}

		var score int
		for len(stack) > 0 {
			var next rune
			next, stack = stack[len(stack)-1], stack[:len(stack)-1]
			score *= 5
			switch next {
			case '(':
				score += 1
			case '[':
				score += 2
			case '{':
				score += 3
			case '<':
				score += 4
			}
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)

	fmt.Printf("Part 2: %d\n", scores[len(scores)/2])
}

func main() {
	partOne()
	partTwo()
}
