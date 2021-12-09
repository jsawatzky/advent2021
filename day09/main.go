package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jsawatzky/advent/helpers"
)

func partOne() {
	heights := make([][]int, 0, 100)

	lines := helpers.ReadInputLines()
	for _, l := range lines {
		heights = append(heights, helpers.AtoiArr(strings.Split(l, "")))
	}

	ans := 0

	for i, r := range heights {
		for j, v := range r {
			low := true
			if i > 0 {
				if heights[i-1][j] <= v {
					low = false
				}
			}
			if j > 0 {
				if heights[i][j-1] <= v {
					low = false
				}
			}
			if i+1 < len(heights) {
				if heights[i+1][j] <= v {
					low = false
				}
			}
			if j+1 < len(r) {
				if heights[i][j+1] <= v {
					low = false
				}
			}
			if low {
				ans += v + 1
			}
		}
	}

	fmt.Printf("Part 1: %d\n", ans)
}

type Point struct {
	X int
	Y int
}

func In(p Point, arr []Point) bool {
	for _, v := range arr {
		if p == v {
			return true
		}
	}
	return false
}

func FindBasin(start Point, heights [][]int) int {
	basin := make([]Point, 0, 100)
	front := make([]Point, 0, 10)
	front = append(front, start)

	var cur Point
	for {
		cur, front = front[len(front)-1], front[:len(front)-1]
		basin = append(basin, cur)

		if cur.X > 0 {
			p := Point{X: cur.X - 1, Y: cur.Y}
			if heights[cur.X-1][cur.Y] < 9 && !In(p, basin) && !In(p, front) {
				front = append(front, p)
			}
		}
		if cur.Y > 0 {
			p := Point{X: cur.X, Y: cur.Y - 1}
			if heights[cur.X][cur.Y-1] < 9 && !In(p, basin) && !In(p, front) {
				front = append(front, p)
			}
		}
		if cur.X+1 < len(heights) {
			p := Point{X: cur.X + 1, Y: cur.Y}
			if heights[cur.X+1][cur.Y] < 9 && !In(p, basin) && !In(p, front) {
				front = append(front, p)
			}
		}
		if cur.Y+1 < len(heights[cur.X]) {
			p := Point{X: cur.X, Y: cur.Y + 1}
			if heights[cur.X][cur.Y+1] < 9 && !In(p, basin) && !In(p, front) {
				front = append(front, p)
			}
		}

		if len(front) == 0 {
			break
		}
	}

	return len(basin)
}

func PrintBasin(start Point, basin []Point, heights [][]int) {
	for i, r := range heights {
		for j, v := range r {
			p := Point{X: i, Y: j}
			if p == start {
				fmt.Print("S")
			} else if In(p, basin) {
				fmt.Print(" ")
			} else {
				fmt.Print(v)
			}
		}
		fmt.Println()
	}
}

func partTwo() {
	heights := make([][]int, 0, 100)

	lines := helpers.ReadInputLines()
	for _, l := range lines {
		heights = append(heights, helpers.AtoiArr(strings.Split(l, "")))
	}

	sizes := make([]int, 0, 10)

	for i, r := range heights {
		for j, v := range r {
			low := true
			if i > 0 {
				if heights[i-1][j] <= v {
					low = false
				}
			}
			if j > 0 {
				if heights[i][j-1] <= v {
					low = false
				}
			}
			if i+1 < len(heights) {
				if heights[i+1][j] <= v {
					low = false
				}
			}
			if j+1 < len(r) {
				if heights[i][j+1] <= v {
					low = false
				}
			}
			if low {
				sizes = append(sizes, FindBasin(Point{X: i, Y: j}, heights))
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	ans := sizes[0] * sizes[1] * sizes[2]

	fmt.Printf("Part 2: %d\n", ans)
}

func main() {
	partOne()
	partTwo()
}
