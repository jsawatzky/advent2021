package main

import (
	"fmt"
	"strings"

	"github.com/jsawatzky/advent/helpers"
)

type Ins struct {
	Ins   string
	Dest  string
	Other string
}

type Block struct {
	ZDiv int
	XAdd int
	YAdd int
}

func readInput() []Block {
	lines := helpers.ReadInputLines()
	blocks := make([]Block, 14)
	for i := 0; i < 14; i++ {
		blocks[i].ZDiv = helpers.Atoi(strings.Split(lines[4], " ")[2])
		blocks[i].XAdd = helpers.Atoi(strings.Split(lines[5], " ")[2])
		blocks[i].YAdd = helpers.Atoi(strings.Split(lines[15], " ")[2])
		lines = lines[18:]
	}
	return blocks
}

type State struct {
	Digit, Z int
}

type Result struct {
	Result string
	Valid  bool
}

var memo map[State]Result

func DFSMax(s State, blocks []Block) Result {
	if s.Digit == 14 {
		if s.Z == 0 {
			return Result{Valid: true}
		}
		return Result{Valid: false}
	}

	if c, ok := memo[s]; ok {
		return c
	}

	res := Result{}
	for i := 9; i > 0; i-- {
		var w, x, y, z int
		w = i
		z = s.Z
		x += z
		x %= 26
		z /= blocks[s.Digit].ZDiv
		x += blocks[s.Digit].XAdd
		if x == w {
			x = 1
		} else {
			x = 0
		}
		if x == 0 {
			x = 1
		} else {
			x = 0
		}
		y += 25
		y *= x
		y++
		z *= y
		y *= 0
		y += w
		y += blocks[s.Digit].YAdd
		y *= x
		z += y
		r := DFSMax(State{Digit: s.Digit + 1, Z: z}, blocks)
		if r.Valid {
			res = Result{Valid: true, Result: fmt.Sprint(i) + r.Result}
			break
		}
	}

	memo[s] = res
	return res
}

func DFSMin(s State, blocks []Block) Result {
	if s.Digit == 14 {
		if s.Z == 0 {
			return Result{Valid: true}
		}
		return Result{Valid: false}
	}

	if c, ok := memo[s]; ok {
		return c
	}

	res := Result{}
	for i := 1; i <= 9; i++ {
		var w, x, y, z int
		w = i
		z = s.Z
		x += z
		x %= 26
		z /= blocks[s.Digit].ZDiv
		x += blocks[s.Digit].XAdd
		if x == w {
			x = 1
		} else {
			x = 0
		}
		if x == 0 {
			x = 1
		} else {
			x = 0
		}
		y += 25
		y *= x
		y++
		z *= y
		y *= 0
		y += w
		y += blocks[s.Digit].YAdd
		y *= x
		z += y
		r := DFSMax(State{Digit: s.Digit + 1, Z: z}, blocks)
		if r.Valid {
			res = Result{Valid: true, Result: fmt.Sprint(i) + r.Result}
			break
		}
	}

	memo[s] = res
	return res
}

func partOne() {

	memo = make(map[State]Result)
	blocks := readInput()

	res := DFSMax(State{}, blocks)
	if !res.Valid {
		panic("not valid")
	}

	fmt.Printf("Part 1: %s\n", res.Result)
}

func partTwo() {

	memo = make(map[State]Result)
	blocks := readInput()

	res := DFSMin(State{}, blocks)
	if !res.Valid {
		panic("not valid")
	}

	fmt.Printf("Part 2: %s\n", res.Result)
}

func main() {
	partOne()
	partTwo()
}
