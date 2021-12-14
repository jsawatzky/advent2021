package main

import (
	"fmt"
	"strings"

	"github.com/jsawatzky/advent/helpers"
)

type Fold struct {
	Axis string
	Line int
}

type Point struct {
	X int
	Y int
}

func GetInput() ([][]bool, []Fold) {
	points := make([]Point, 0, 100)
	var maxX, maxY int

	lines := helpers.ReadInputLines()
	var i int
	var l string
	for i, l = range lines {
		if len(l) == 0 {
			break
		}
		cord := strings.Split(l, ",")
		x, y := helpers.Atoi(cord[0]), helpers.Atoi(cord[1])
		points = append(points, Point{X: x, Y: y})
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}
	lines = lines[i+1:]

	paper := make([][]bool, 0, maxY+1)
	for i := 0; i < maxY+1; i++ {
		paper = append(paper, make([]bool, maxX+1))
	}

	for _, p := range points {
		paper[p.Y][p.X] = true
	}

	folds := make([]Fold, 0, 30)

	for _, l := range lines {
		l = strings.TrimPrefix(l, "fold along ")
		fold := strings.Split(l, "=")
		folds = append(folds, Fold{Axis: fold[0], Line: helpers.Atoi(fold[1])})
	}

	return paper, folds
}

func DoFold(paper [][]bool, fold Fold) [][]bool {
	if fold.Axis == "y" {
		for i := 1; i+fold.Line < len(paper) && fold.Line-i >= 0; i++ {
			r1 := paper[fold.Line+i]
			r2 := paper[fold.Line-i]
			for j := 0; j < len(r1); j++ {
				r1[j] = r1[j] || r2[j]
			}
			paper[fold.Line-i] = r1
		}
		paper = paper[:fold.Line]
	} else {
		for i, r := range paper {
			for j := 1; j+fold.Line < len(r) && fold.Line-j >= 0; j++ {
				r[fold.Line-j] = r[fold.Line-j] || r[fold.Line+j]
			}
			paper[i] = r[:fold.Line]
		}
	}

	return paper
}

func partOne() {

	paper, folds := GetInput()

	paper = DoFold(paper, folds[0])

	ans := 0

	for _, r := range paper {
		for _, v := range r {
			if v {
				ans++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", ans)
}

func PrintPaper(paper [][]bool) {
	for _, r := range paper {
		for _, v := range r {
			if v {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func partTwo() {

	paper, folds := GetInput()

	for _, f := range folds {
		paper = DoFold(paper, f)
	}

	fmt.Println("Part 2:")
	PrintPaper(paper)
}

func main() {
	partOne()
	partTwo()
}
