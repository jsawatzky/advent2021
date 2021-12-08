package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jsawatzky/advent/helpers"
)

func partOne() {
	lines := helpers.ReadInputLines()

	ans := 0
	for _, l := range lines {
		parts := strings.Split(l, " | ")
		out := strings.Fields(parts[1])
		for _, o := range out {
			switch len(o) {
			case 2:
				fallthrough
			case 4:
				fallthrough
			case 3:
				fallthrough
			case 7:
				ans++
			default:
			}
		}
	}

	fmt.Printf("Part 1: %d\n", ans)
}

func In(val rune, s string) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}
	return false
}

func partTwo() {
	lines := helpers.ReadInputLines()

	ans := 0
	for _, l := range lines {
		parts := strings.Split(l, " | ")
		in := strings.Fields(parts[0])
		sort.Slice(in, func(i, j int) bool {
			return len(in[i]) < len(in[j])
		})
		out := strings.Fields(parts[1])

		occ := make(map[rune]int)
		for _, s := range in {
			for _, c := range s {
				occ[c]++
			}
		}

		var A, B, C, D, E rune
		for k, v := range occ {
			if v == 6 {
				B = k
			} else if v == 4 {
				E = k
			}
		}
		for _, c := range in[1] {
			if !In(c, in[0]) {
				A = c
				break
			}
		}
		for _, c := range in[2] {
			if !In(c, in[0]) && c != B {
				D = c
				break
			}
		}
		for k, v := range occ {
			if v == 8 && k != A {
				C = k
			}
		}

		res := ""
		for _, o := range out {
			switch len(o) {
			case 2:
				res += "1"
			case 3:
				res += "7"
			case 4:
				res += "4"
			case 5:
				if In(E, o) {
					res += "2"
				} else if In(C, o) {
					res += "3"
				} else {
					res += "5"
				}
			case 6:
				if In(D, o) {
					if In(C, o) {
						res += "9"
					} else {
						res += "6"
					}
				} else {
					res += "0"
				}
			case 7:
				res += "8"
			}
		}
		ans += helpers.Atoi(res)
	}

	fmt.Printf("Part 2: %d\n", ans)
}

func main() {
	partOne()
	partTwo()
}
