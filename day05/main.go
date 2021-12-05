package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Line struct {
	Start Point
	End   Point
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	points := make([][]int, 1000)
	for i := range points {
		points[i] = make([]int, 1000)
	}
	lines := make([]Line, 0, 500)

	for scanner.Scan() {
		points := strings.Split(scanner.Text(), " -> ")
		if len(points) != 2 {
			panic("Too many points")
		}
		start := strings.Split(points[0], ",")
		if len(start) != 2 {
			panic("Too many nums")
		}
		startP := Point{}
		var err error
		startP.X, err = strconv.Atoi(start[0])
		if err != nil {
			panic(err)
		}
		startP.Y, err = strconv.Atoi(start[1])
		if err != nil {
			panic(err)
		}
		end := strings.Split(points[1], ",")
		if len(start) != 2 {
			panic("Too many nums")
		}
		endP := Point{}
		endP.X, err = strconv.Atoi(end[0])
		if err != nil {
			panic(err)
		}
		endP.Y, err = strconv.Atoi(end[1])
		if err != nil {
			panic(err)
		}

		if startP.X != endP.X && startP.Y != endP.Y {
			continue
		}

		lines = append(lines, Line{Start: startP, End: endP})
	}

	for _, l := range lines {
		if l.Start.X == l.End.X {
			row := l.Start.X
			colStart := l.Start.Y
			colEnd := l.End.Y
			if colStart > colEnd {
				colStart, colEnd = colEnd, colStart
			}
			for col := colStart; col <= colEnd; col++ {
				points[row][col]++
			}
		} else {
			col := l.Start.Y
			rowStart := l.Start.X
			rowEnd := l.End.X
			if rowStart > rowEnd {
				rowStart, rowEnd = rowEnd, rowStart
			}
			for row := rowStart; row <= rowEnd; row++ {
				points[row][col]++
			}
		}
	}

	var numOverlap int
	for _, r := range points {
		for _, val := range r {
			if val > 1 {
				numOverlap++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", numOverlap)

}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	points := make([][]int, 1000)
	for i := range points {
		points[i] = make([]int, 1000)
	}
	lines := make([]Line, 0, 500)

	for scanner.Scan() {
		points := strings.Split(scanner.Text(), " -> ")
		if len(points) != 2 {
			panic("Too many points")
		}
		start := strings.Split(points[0], ",")
		if len(start) != 2 {
			panic("Too many nums")
		}
		startP := Point{}
		var err error
		startP.X, err = strconv.Atoi(start[0])
		if err != nil {
			panic(err)
		}
		startP.Y, err = strconv.Atoi(start[1])
		if err != nil {
			panic(err)
		}
		end := strings.Split(points[1], ",")
		if len(start) != 2 {
			panic("Too many nums")
		}
		endP := Point{}
		endP.X, err = strconv.Atoi(end[0])
		if err != nil {
			panic(err)
		}
		endP.Y, err = strconv.Atoi(end[1])
		if err != nil {
			panic(err)
		}

		lines = append(lines, Line{Start: startP, End: endP})
	}

	for _, l := range lines {
		if l.Start.X == l.End.X {
			row := l.Start.X
			colStart := l.Start.Y
			colEnd := l.End.Y
			if colStart > colEnd {
				colStart, colEnd = colEnd, colStart
			}
			for col := colStart; col <= colEnd; col++ {
				points[row][col]++
			}
		} else if l.Start.Y == l.End.Y {
			col := l.Start.Y
			rowStart := l.Start.X
			rowEnd := l.End.X
			if rowStart > rowEnd {
				rowStart, rowEnd = rowEnd, rowStart
			}
			for row := rowStart; row <= rowEnd; row++ {
				points[row][col]++
			}
		} else {
			rowStart := l.Start.X
			rowEnd := l.End.X
			colStart := l.Start.Y
			colEnd := l.End.Y
			if rowStart > rowEnd {
				rowStart, rowEnd = rowEnd, rowStart
				colStart, colEnd = colEnd, colStart
			}
			colMag := 1
			if colStart > colEnd {
				colMag = -1
			}
			lineLen := rowEnd - rowStart
			for i := 0; i <= lineLen; i++ {
				points[rowStart+i][colStart+(i*colMag)]++
			}
		}
	}

	var numOverlap int
	for _, r := range points {
		for _, val := range r {
			if val > 1 {
				numOverlap++
			}
		}
	}

	fmt.Printf("Part 2: %d\n", numOverlap)
}

func main() {
	partOne()
	partTwo()
}
