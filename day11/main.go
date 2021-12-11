package main

import (
	"fmt"
	"strings"

	"github.com/jsawatzky/advent/helpers"
)

func Flash(i, j int, octos [][]int) int {
	if octos[i][j] < 0 {
		return 0
	}
	octos[i][j] = -9999
	flashes := 1
	if i > 0 && j > 0 {
		octos[i-1][j-1]++
		if octos[i-1][j-1] > 9 {
			flashes += Flash(i-1, j-1, octos)
		}
	}
	if i > 0 {
		octos[i-1][j]++
		if octos[i-1][j] > 9 {
			flashes += Flash(i-1, j, octos)
		}
	}
	if i > 0 && j+1 < len(octos[i]) {
		octos[i-1][j+1]++
		if octos[i-1][j+1] > 9 {
			flashes += Flash(i-1, j+1, octos)
		}
	}
	if j+1 < len(octos[i]) {
		octos[i][j+1]++
		if octos[i][j+1] > 9 {
			flashes += Flash(i, j+1, octos)
		}
	}
	if i+1 < len(octos) && j+1 < len(octos[i]) {
		octos[i+1][j+1]++
		if octos[i+1][j+1] > 9 {
			flashes += Flash(i+1, j+1, octos)
		}
	}
	if i+1 < len(octos) {
		octos[i+1][j]++
		if octos[i+1][j] > 9 {
			flashes += Flash(i+1, j, octos)
		}
	}
	if i+1 < len(octos) && j > 0 {
		octos[i+1][j-1]++
		if octos[i+1][j-1] > 9 {
			flashes += Flash(i+1, j-1, octos)
		}
	}
	if j > 0 {
		octos[i][j-1]++
		if octos[i][j-1] > 9 {
			flashes += Flash(i, j-1, octos)
		}
	}
	return flashes
}

func partOne() {
	lines := helpers.ReadInputLines()
	octos := make([][]int, 0, 10)
	for _, l := range lines {
		octos = append(octos, helpers.AtoiArr(strings.Split(l, "")))
	}

	ans := 0

	for i := 0; i < 100; i++ {
		for i, r := range octos {
			for j := range r {
				octos[i][j]++
			}
		}
		for i, r := range octos {
			for j, o := range r {
				if o > 9 {
					ans += Flash(i, j, octos)
				}
			}
		}
		for i, r := range octos {
			for j, o := range r {
				if o < 0 {
					octos[i][j] = 0
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", ans)
}

func partTwo() {
	lines := helpers.ReadInputLines()
	octos := make([][]int, 0, 10)
	for _, l := range lines {
		octos = append(octos, helpers.AtoiArr(strings.Split(l, "")))
	}

	step := 1
	for {
		for i, r := range octos {
			for j := range r {
				octos[i][j]++
			}
		}
		for i, r := range octos {
			for j, o := range r {
				if o > 9 {
					Flash(i, j, octos)
				}
			}
		}
		all := true
		for i, r := range octos {
			for j, o := range r {
				if o < 0 {
					octos[i][j] = 0
				} else {
					all = false
				}
			}
		}
		if all {
			break
		}
		step++
	}

	fmt.Printf("Part 2: %d\n", step)
}

func main() {
	partOne()
	partTwo()
}
