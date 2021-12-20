package main

import (
	"fmt"

	"github.com/jsawatzky/advent/helpers"
)

type Point struct {
	X int
	Y int
}

func In(p Point, points []Point) bool {
	for _, v := range points {
		if p == v {
			return true
		}
	}
	return false
}

func readInput() ([]rune, []Point, int, int) {
	lines := helpers.ReadInputLines()
	alg := []rune(lines[0])
	lines = lines[2:]
	points := []Point{}
	for i, l := range lines {
		for j, c := range l {
			if c == '#' {
				points = append(points, Point{X: j, Y: i})
			}
		}
	}
	return alg, points, len(lines[0]), len(lines)
}

func Print(p []Point, mx, my, Mx, My int) {
	for j := my - 3; j <= My+3; j++ {
		for i := mx - 3; i <= Mx+3; i++ {
			if In(Point{X: i, Y: j}, p) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func partOne() {

	alg, points, w, h := readInput()
	background := 0

	minX := 0
	minY := 0
	maxX := w
	maxY := h

	for iter := 0; iter < 2; iter++ {
		newPoints := []Point{}
		for i := minX - 1; i <= maxX+1; i++ {
			for j := minY - 1; j <= maxY+1; j++ {
				bin := ""
				for y := j - 1; y <= j+1; y++ {
					for x := i - 1; x <= i+1; x++ {
						if x < minX || y < minY || x > maxX || y > maxY {
							bin += fmt.Sprint(background)
						} else {
							if In(Point{X: x, Y: y}, points) {
								bin += "1"
							} else {
								bin += "0"
							}
						}
					}
				}
				idx := helpers.BinaryToInt(bin)
				if alg[idx] == '#' {
					newPoints = append(newPoints, Point{X: i, Y: j})
				}
			}
		}
		minX--
		minY--
		maxX++
		maxY++
		if background == 0 && alg[0] == '#' {
			background = 1
		} else if background == 1 && alg[511] == '.' {
			background = 0
		}
		points = newPoints
	}

	ans := len(points)

	fmt.Printf("Part 1: %d\n", ans)
}

func partTwo() {

	alg, points, w, h := readInput()
	background := 0

	minX := 0
	minY := 0
	maxX := w
	maxY := h

	for iter := 0; iter < 50; iter++ {
		newPoints := []Point{}
		for i := minX - 1; i <= maxX+1; i++ {
			for j := minY - 1; j <= maxY+1; j++ {
				bin := ""
				for y := j - 1; y <= j+1; y++ {
					for x := i - 1; x <= i+1; x++ {
						if x < minX || y < minY || x > maxX || y > maxY {
							bin += fmt.Sprint(background)
						} else {
							if In(Point{X: x, Y: y}, points) {
								bin += "1"
							} else {
								bin += "0"
							}
						}
					}
				}
				idx := helpers.BinaryToInt(bin)
				if alg[idx] == '#' {
					newPoints = append(newPoints, Point{X: i, Y: j})
				}
			}
		}
		minX--
		minY--
		maxX++
		maxY++
		if background == 0 && alg[0] == '#' {
			background = 1
		} else if background == 1 && alg[511] == '.' {
			background = 0
		}
		points = newPoints
	}

	ans := len(points)

	fmt.Printf("Part 2: %d\n", ans)
}

func main() {
	partOne()
	partTwo()
}
