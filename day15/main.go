package main

import (
	"container/heap"
	"fmt"
	"math"
	"strings"

	"github.com/jsawatzky/advent/helpers"
)

type GridPoint struct {
	X    int
	Y    int
	Risk int
	Dist int
	Prev *GridPoint
}

type IntHeap []*GridPoint

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i].Dist < h[j].Dist
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(*GridPoint))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h IntHeap) In(gp *GridPoint) int {
	for i, v := range h {
		if v == gp {
			return i
		}
	}
	return -1
}

func readInput() [][]*GridPoint {
	res := make([][]int, 0, 10)
	lines := helpers.ReadInputLines()
	for _, l := range lines {
		res = append(res, helpers.AtoiArr(strings.Split(l, "")))
	}
	grid := make([][]*GridPoint, 0, len(res))
	for i, r := range res {
		a := make([]*GridPoint, 0, len(r))
		for j, v := range r {
			a = append(a, &GridPoint{X: i, Y: j, Risk: v, Dist: math.MaxInt32})
		}
		grid = append(grid, a)
	}
	return grid
}

func partOne() {

	grid := readInput()
	grid[0][0].Dist = 0
	targetX := len(grid) - 1
	targetY := len(grid[targetX]) - 1

	h := &IntHeap{}
	for _, r := range grid {
		for _, v := range r {
			heap.Push(h, v)
		}
	}

	for h.Len() > 0 {
		p := heap.Pop(h).(*GridPoint)
		helpers.ForEachNeighborInt(p.X, p.Y, len(grid), len(grid[p.X]), false, func(i1, i2 int) {
			neigh := grid[i1][i2]
			index := h.In(neigh)
			if index < 0 {
				return
			}
			newDist := p.Dist + neigh.Risk
			if newDist < neigh.Dist {
				neigh.Dist = newDist
				neigh.Prev = p
			}
			heap.Fix(h, index)
		})
	}

	target := grid[targetX][targetY]

	ans := target.Dist

	fmt.Printf("Part 1: %d\n", ans)
}

func partTwo() {

	smallGrid := readInput()
	grid := make([][]*GridPoint, 0, len(smallGrid)*5)
	for i := 0; i < len(smallGrid)*5; i++ {
		grid = append(grid, make([]*GridPoint, 0, len(smallGrid[0])*5))
	}

	for i, r := range smallGrid {
		grid[i] = append(grid[i], r...)
	}

	for gi := 0; gi < 5; gi++ {
		for gj := 0; gj < 5; gj++ {
			if gi == 0 && gj == 0 {
				continue
			}

			newTile := make([][]*GridPoint, 0, len(smallGrid))
			for _, r := range smallGrid {
				row := make([]*GridPoint, 0, len(r))
				for _, v := range r {
					newRisk := v.Risk + gi + gj
					if newRisk > 9 {
						newRisk -= 9
					}
					row = append(row, &GridPoint{X: gi*len(smallGrid) + v.X, Y: gj*len(smallGrid[v.X]) + v.Y, Risk: newRisk, Dist: v.Dist, Prev: v.Prev})
				}
				newTile = append(newTile, row)
			}

			for i, r := range newTile {
				grid[gi*len(smallGrid)+i] = append(grid[gi*len(smallGrid)+i], r...)
			}
		}
	}

	grid[0][0].Dist = 0
	targetX := len(grid) - 1
	targetY := len(grid[targetX]) - 1

	h := &IntHeap{}
	for _, r := range grid {
		for _, v := range r {
			heap.Push(h, v)
		}
	}

	for h.Len() > 0 {
		p := heap.Pop(h).(*GridPoint)
		if p.Dist > 50000 {
			panic("wtf")
		}
		helpers.ForEachNeighborInt(p.X, p.Y, len(grid), len(grid[p.X]), false, func(i1, i2 int) {
			neigh := grid[i1][i2]
			index := h.In(neigh)
			if index < 0 {
				return
			}
			newDist := p.Dist + neigh.Risk
			if newDist < neigh.Dist {
				neigh.Dist = newDist
				neigh.Prev = p
			}
			heap.Fix(h, index)
		})
	}

	target := grid[targetX][targetY]

	ans := target.Dist

	fmt.Printf("Part 2: %d\n", ans)
}

func main() {
	partOne()
	partTwo()
}
