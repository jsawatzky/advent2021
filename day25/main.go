package main

import (
	"fmt"

	"github.com/jsawatzky/advent/helpers"
)

func readInput() [][]rune {
	lines := helpers.ReadInputLines()
	grid := [][]rune{}
	for _, l := range lines {
		r := []rune{}
		for _, c := range l {
			r = append(r, c)
		}
		grid = append(grid, r)
	}
	return grid
}

func DoMove(herd rune, grid [][]rune) ([][]rune, bool) {
	newGrid := make([][]rune, len(grid))
	for i := 0; i < len(newGrid); i++ {
		newGrid[i] = make([]rune, len(grid[i]))
		for j := 0; j < len(newGrid[i]); j++ {
			newGrid[i][j] = '.'
		}
	}

	moved := false

	for i, r := range grid {
		for j, v := range r {
			if v == '.' {
				continue
			} else if v != herd {
				newGrid[i][j] = v
				continue
			}
			switch herd {
			case '>':
				if grid[i][(j+1)%len(r)] == '.' {
					newGrid[i][(j+1)%len(r)] = v
					moved = true
				} else {
					newGrid[i][j] = v
				}
			case 'v':
				if grid[(i+1)%len(grid)][j] == '.' {
					newGrid[(i+1)%len(grid)][j] = v
					moved = true
				} else {
					newGrid[i][j] = v
				}
			}
		}
	}

	return newGrid, moved
}

func partOne() {

	grid := readInput()

	var ans int
	for i := 0; ; i++ {
		var h1Moved, h2Moved bool

		grid, h1Moved = DoMove('>', grid)
		grid, h2Moved = DoMove('v', grid)

		if !(h1Moved || h2Moved) {
			ans = i + 1
			break
		}
	}

	fmt.Printf("Part 1: %d\n", ans)
}

func partTwo() {

	fmt.Printf("Part 2: ALL DONE\n")
}

func main() {
	partOne()
	partTwo()
}
