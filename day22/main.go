package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jsawatzky/advent/helpers"
)

type Step struct {
	State                  bool
	X1, X2, Y1, Y2, Z1, Z2 int
}

func (s Step) InP1Range() bool {
	return helpers.Abs(s.X1) <= 50 &&
		helpers.Abs(s.X2) <= 50 &&
		helpers.Abs(s.Y1) <= 50 &&
		helpers.Abs(s.Y2) <= 50 &&
		helpers.Abs(s.Z1) <= 50 &&
		helpers.Abs(s.Z2) <= 50
}

func (s Step) Size() int {
	return (s.X2 - s.X1) * (s.Y2 - s.Y1) * (s.Z2 - s.Z1)
}

func (s Step) Intersection(s2 Step) Step {
	return Step{
		X1: helpers.Max(s.X1, s2.X1),
		X2: helpers.Min(s.X2, s2.X2),
		Y1: helpers.Max(s.Y1, s2.Y1),
		Y2: helpers.Min(s.Y2, s2.Y2),
		Z1: helpers.Max(s.Z1, s2.Z1),
		Z2: helpers.Min(s.Z2, s2.Z2),
	}
}

func (s Step) Intersects(s2 Step) bool {
	i := s.Intersection(s2)
	return i.X1 < i.X2 && i.Y1 < i.Y2 && i.Z1 < i.Z2
}

func readInput() []Step {
	lines := helpers.ReadInputLines()
	steps := []Step{}
	for _, l := range lines {
		s := Step{}
		ins := strings.Split(l, " ")
		s.State = ins[0] == "on"
		cube := strings.Split(ins[1], ",")
		coord := helpers.AtoiArr(strings.Split(cube[0][2:], ".."))
		s.X1 = helpers.Min(coord[0], coord[1])
		s.X2 = helpers.Max(coord[0], coord[1]) + 1
		coord = helpers.AtoiArr(strings.Split(cube[1][2:], ".."))
		s.Y1 = helpers.Min(coord[0], coord[1])
		s.Y2 = helpers.Max(coord[0], coord[1]) + 1
		coord = helpers.AtoiArr(strings.Split(cube[2][2:], ".."))
		s.Z1 = helpers.Min(coord[0], coord[1])
		s.Z2 = helpers.Max(coord[0], coord[1]) + 1
		steps = append(steps, s)
	}
	return steps
}

func partOne() {

	steps := readInput()
	grid := [101][101][101]bool{}

	for _, s := range steps {
		if s.InP1Range() {
			for x := s.X1; x < s.X2; x++ {
				for y := s.Y1; y < s.Y2; y++ {
					for z := s.Z1; z < s.Z2; z++ {
						grid[x+50][y+50][z+50] = s.State
					}
				}
			}
		}
	}

	ans := 0

	for _, s := range grid {
		for _, r := range s {
			for _, v := range r {
				if v {
					ans++
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", ans)
}

func partTwo() {
	steps := readInput()

	cubes := make(map[Step]struct{})

	for _, s := range steps {
	loop:
		for c := range cubes {
			if s.Intersects(c) {
				delete(cubes, c)

				xs := []int{s.X1, s.X2, c.X1, c.X2}
				sort.Ints(xs)
				ys := []int{s.Y1, s.Y2, c.Y1, c.Y2}
				sort.Ints(ys)
				zs := []int{s.Z1, s.Z2, c.Z1, c.Z2}
				sort.Ints(zs)

				for i := 0; i < len(xs)-1; i++ {
					for j := 0; j < len(ys)-1; j++ {
						for k := 0; k < len(zs)-1; k++ {
							newCube := Step{
								X1: xs[i],
								X2: xs[i+1],
								Y1: ys[j],
								Y2: ys[j+1],
								Z1: zs[k],
								Z2: zs[k+1],
							}
							if c.Intersects(newCube) && !s.Intersects(newCube) {
								cubes[newCube] = struct{}{}
							}
						}
					}
				}
				goto loop
			}
		}
		if s.State {
			cubes[s] = struct{}{}
		}

	}

	var ans uint64

	for c := range cubes {
		ans += uint64(c.Size())
	}

	fmt.Printf("Part 2: %d\n", ans)
}

func main() {
	partOne()
	partTwo()
}
