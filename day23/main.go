package main

import (
	"fmt"
	"math"

	"github.com/jsawatzky/advent/helpers"
)

var BarredSpaces = []int{2, 4, 6, 8}
var ExpectedRooms = []rune{'A', 'B', 'C', 'D'}
var RoomMap = map[rune]int{
	'A': 0,
	'B': 1,
	'C': 2,
	'D': 3,
}
var EnergyMap = map[rune]int{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
}

type Room [2]rune
type Hallway [11]rune
type Input struct {
	Hallway Hallway
	Rooms   [4]Room
}

func (in Input) Solved() bool {
	for i, r := range in.Rooms {
		if r[0] == '.' || r[1] == '.' {
			return false
		} else if r[0] == r[1] && r[0] == ExpectedRooms[i] {
			continue
		} else {
			return false
		}
	}
	return true
}

func (in Input) Print() {
	fmt.Println("#############")
	fmt.Print("#")
	for _, h := range in.Hallway {
		fmt.Printf("%c", h)
	}
	fmt.Println("#")
	fmt.Print("###")
	for _, r := range in.Rooms {
		fmt.Printf("%c#", r[0])
	}
	fmt.Println("##")
	fmt.Print("  #")
	for _, r := range in.Rooms {
		fmt.Printf("%c#", r[1])
	}
	fmt.Println()
	fmt.Println("  #########")
}

func readInput() Input {
	lines := helpers.ReadInputLines()
	var input Input
	for i := range input.Hallway {
		input.Hallway[i] = '.'
	}
	r0 := []rune(lines[2])
	r1 := []rune(lines[3])
	input.Rooms[0][0] = r0[3]
	input.Rooms[0][1] = r1[1]
	input.Rooms[1][0] = r0[5]
	input.Rooms[1][1] = r1[3]
	input.Rooms[2][0] = r0[7]
	input.Rooms[2][1] = r1[5]
	input.Rooms[3][0] = r0[9]
	input.Rooms[3][1] = r1[7]

	return input
}

type CacheValue struct {
	Energy      int
	StartEnergy int
	Valid       bool
}

var memo map[Input]CacheValue = make(map[Input]CacheValue)

func Solve(in Input, total int) (int, bool) {
	if c, ok := memo[in]; ok {
		if !c.Valid {
			return 0, false
		} else if c.StartEnergy < total {
			return 0, false
		}
	}

	if in.Solved() {
		return total, true
	}

	min := math.MaxInt32
	valid := false

	for i, r := range in.Rooms {
		if r[0] == '.' && r[1] == '.' {
			continue
		} else if r[0] == r[1] && r[0] == ExpectedRooms[i] {
			continue
		} else if r[0] == '.' && r[1] == ExpectedRooms[i] {
			continue
		}
		j := 0
		if r[j] == '.' {
			j = 1
		}

		for h := BarredSpaces[i]; h < 11; h++ {
			if helpers.InInt(h, BarredSpaces) {
				continue
			} else if in.Hallway[h] != '.' {
				break
			}
			newHall := in.Hallway
			newRooms := in.Rooms
			newHall[h] = r[j]
			newRooms[i][j] = '.'
			energy := (j + helpers.Abs(BarredSpaces[i]-h) + 1) * EnergyMap[r[j]]
			if m, v := Solve(Input{Hallway: newHall, Rooms: newRooms}, total+energy); v {
				min = helpers.Min(m, min)
				valid = true
			}
		}
		for h := BarredSpaces[i]; h >= 0; h-- {
			if helpers.InInt(h, BarredSpaces) {
				continue
			} else if in.Hallway[h] != '.' {
				break
			}
			newHall := in.Hallway
			newRooms := in.Rooms
			newHall[h] = r[j]
			newRooms[i][j] = '.'
			energy := (j + helpers.Abs(BarredSpaces[i]-h) + 1) * EnergyMap[r[j]]
			if m, v := Solve(Input{Hallway: newHall, Rooms: newRooms}, total+energy); v {
				min = helpers.Min(m, min)
				valid = true
			}
		}
	}

hallLoop:
	for i, h := range in.Hallway {
		if h == '.' {
			continue
		}
		r := RoomMap[h]
		if in.Rooms[r][0] != '.' {
			continue
		} else if in.Rooms[r][1] != h && in.Rooms[r][1] != '.' {
			continue
		}
		dir := 1
		if BarredSpaces[r] < i {
			dir = -1
		}
		for j := i + dir; j != BarredSpaces[r]; j += dir {
			if in.Hallway[j] != '.' {
				continue hallLoop
			}
		}
		if in.Rooms[r][1] == '.' {
			newHall := in.Hallway
			newRooms := in.Rooms
			newHall[i] = '.'
			newRooms[r][1] = h
			energy := (helpers.Abs(BarredSpaces[r]-i) + 2) * EnergyMap[h]
			if m, v := Solve(Input{Hallway: newHall, Rooms: newRooms}, total+energy); v {
				min = helpers.Min(m, min)
				valid = true
			}
		} else {
			newHall := in.Hallway
			newRooms := in.Rooms
			newHall[i] = '.'
			newRooms[r][0] = h
			energy := (helpers.Abs(BarredSpaces[r]-i) + 1) * EnergyMap[h]
			if m, v := Solve(Input{Hallway: newHall, Rooms: newRooms}, total+energy); v {
				min = helpers.Min(m, min)
				valid = true
			}
		}
	}

	memo[in] = CacheValue{Energy: min, StartEnergy: total, Valid: valid}

	return min, valid
}

func partOne() {
	input := readInput()

	ans, val := Solve(input, 0)
	if !val {
		panic("Input not valid")
	}

	fmt.Printf("Part 1: %d\n", ans)
}

func partTwo() {

	ans := 0

	fmt.Printf("Part 2: %d\n", ans)
}

func main() {
	partOne()
	partTwo()
}
