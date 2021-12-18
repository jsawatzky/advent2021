package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jsawatzky/advent/helpers"
)

type Number struct {
	Value  int
	Left   *Number
	Right  *Number
	Parent *Number
	Depth  int
}

func (n *Number) IsRegular() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Number) IsPair() bool {
	return n.Left != nil && n.Right != nil
}

func (n *Number) IsLeft() bool {
	if n.Parent == nil {
		return false
	}
	return n.Parent.Left == n
}

func (n *Number) Explode() {
	if n.IsRegular() {
		panic("cannot explode regular")
	}

	if !n.Left.IsRegular() && !n.Right.IsRegular() {
		panic("cannot explode non regular pair " + n.Print())
	}

	if n.IsLeft() {
		nextParent := n.Parent
		for nextParent != nil && nextParent.IsLeft() {
			nextParent = nextParent.Parent
		}
		if nextParent != nil && nextParent.Parent != nil {
			reg := nextParent.Parent.Left
			for !reg.IsRegular() {
				reg = reg.Right
			}
			reg.Value += n.Left.Value
		}

		reg := n.Parent.Right
		for !reg.IsRegular() {
			reg = reg.Left
		}
		reg.Value += n.Right.Value
	} else {
		reg := n.Parent.Left
		for !reg.IsRegular() {
			reg = reg.Right
		}
		reg.Value += n.Left.Value

		nextParent := n.Parent
		for nextParent != nil && !nextParent.IsLeft() {
			nextParent = nextParent.Parent
		}
		if nextParent != nil && nextParent.Parent != nil {
			reg := nextParent.Parent.Right
			for !reg.IsRegular() {
				reg = reg.Left
			}
			reg.Value += n.Right.Value
		}
	}

	n.Left = nil
	n.Right = nil
	n.Value = 0
}

func (n *Number) Split() {
	if !n.IsRegular() {
		panic("cannot split non regular")
	}

	n.Left = &Number{
		Value:  n.Value / 2,
		Parent: n,
		Depth:  n.Depth + 1,
	}
	n.Right = &Number{
		Value:  n.Value - n.Left.Value,
		Parent: n,
		Depth:  n.Depth + 1,
	}
}

func (n *Number) Nest() {
	n.Depth++
	if n.Left != nil {
		n.Left.Nest()
	}
	if n.Right != nil {
		n.Right.Nest()
	}
}

func (n *Number) Magnitude() int {
	if n.IsRegular() {
		return n.Value
	}
	return 3*n.Left.Magnitude() + 2*n.Right.Magnitude()
}

func (n *Number) ReduceExplode() bool {
	if n.IsPair() && n.Depth == 4 {
		n.Explode()
		return true
	}

	if n.IsPair() {
		return n.Left.ReduceExplode() || n.Right.ReduceExplode()
	}
	return false
}

func (n *Number) ReduceSplit() bool {
	if n.IsRegular() && n.Value >= 10 {
		n.Split()
		return true
	}

	if n.IsPair() {
		return n.Left.ReduceSplit() || n.Right.ReduceSplit()
	}
	return false
}

func (n *Number) Print() string {
	if n.IsRegular() {
		return strconv.FormatInt(int64(n.Value), 10)
	} else {
		return fmt.Sprintf("[%s,%s]", n.Left.Print(), n.Right.Print())
	}
}

func (n *Number) Copy() *Number {
	c := &Number{}
	data, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, c)
	if err != nil {
		panic(err)
	}
	return c
}

func ParseNumber(s string, depth int) (*Number, string) {
	if s[0] != '[' {
		panic("invalid number")
	}

	result := &Number{Depth: depth}
	s = s[1:]
	if s[0] == '[' {
		result.Left, s = ParseNumber(s, depth+1)
		s = s[1:]
	} else {
		result.Left = &Number{Value: helpers.Atoi(s[:1]), Depth: depth + 1}
		s = s[2:]
	}
	if s[0] == '[' {
		result.Right, s = ParseNumber(s, depth+1)
	} else {
		result.Right = &Number{Value: helpers.Atoi(s[:1]), Depth: depth + 1}
		s = s[1:]
	}

	result.Left.Parent = result
	result.Right.Parent = result

	if s[0] != ']' {
		panic("invalid number end")
	}

	return result, s[1:]
}

func readInput() []*Number {
	lines := helpers.ReadInputLines()
	numbers := make([]*Number, len(lines))
	for i, l := range lines {
		var left string
		numbers[i], left = ParseNumber(l, 0)
		if len(left) > 0 {
			panic("leftover number")
		}
	}
	return numbers
}

func partOne() {
	nums := readInput()

	total := nums[0]
	for _, n := range nums[1:] {
		total.Nest()
		n.Nest()
		total = &Number{Left: total, Right: n}
		total.Left.Parent = total
		total.Right.Parent = total
		for {
			if total.ReduceExplode() {
				continue
			} else if total.ReduceSplit() {
				continue
			}
			break
		}
	}

	ans := total.Magnitude()

	fmt.Printf("Part 1: %d\n", ans)
}

func partTwo() {

	nums := readInput()

	maxMag := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			nums = readInput()
			nums[i].Nest()
			nums[j].Nest()
			sum := &Number{Left: nums[i], Right: nums[j]}
			sum.Left.Parent = sum
			sum.Right.Parent = sum
			for {
				if sum.ReduceExplode() {
					continue
				} else if sum.ReduceSplit() {
					continue
				}
				break
			}
			mag := sum.Magnitude()
			if mag > maxMag {
				maxMag = mag
			}
		}
	}

	fmt.Printf("Part 2: %d\n", maxMag)
}

func main() {
	partOne()
	partTwo()
}
