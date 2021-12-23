package main

import (
	"fmt"
	"math"
	"sync"

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

type Room [4]rune
type Hallway [11]rune
type Input struct {
	Hallway Hallway
	Rooms   [4]Room
}

func (in Input) Solved() bool {
	for i, r := range in.Rooms {
		if r[0] == '.' || r[1] == '.' || r[2] == '.' || r[3] == '.' {
			return false
		} else if r[0] == r[1] && r[0] == r[2] && r[0] == r[3] && r[0] == ExpectedRooms[i] {
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
	fmt.Print("  #")
	for _, r := range in.Rooms {
		fmt.Printf("%c#", r[2])
	}
	fmt.Println()
	fmt.Print("  #")
	for _, r := range in.Rooms {
		fmt.Printf("%c#", r[3])
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
	input.Rooms[0][1] = 'D'
	input.Rooms[0][2] = 'D'
	input.Rooms[0][3] = r1[1]
	input.Rooms[1][0] = r0[5]
	input.Rooms[1][1] = 'C'
	input.Rooms[1][2] = 'B'
	input.Rooms[1][3] = r1[3]
	input.Rooms[2][0] = r0[7]
	input.Rooms[2][1] = 'B'
	input.Rooms[2][2] = 'A'
	input.Rooms[2][3] = r1[5]
	input.Rooms[3][0] = r0[9]
	input.Rooms[3][1] = 'A'
	input.Rooms[3][2] = 'C'
	input.Rooms[3][3] = r1[7]

	return input
}

type CacheValue struct {
	StartEnergy int
	Valid       bool
}

type Result struct {
	Energy int
	Valid  bool
}

var memo map[Input]CacheValue = make(map[Input]CacheValue)
var mu sync.Mutex

func Solve(in Input, total int, results chan<- Result) {
	defer close(results)

	mu.Lock()
	c, ok := memo[in]
	mu.Unlock()
	if ok {
		if !c.Valid {
			return
		} else if c.StartEnergy < total {
			return
		}
	}

	if in.Solved() {
		results <- Result{Energy: total, Valid: true}
		return
	}

	min := math.MaxInt32
	valid := false

	for i, r := range in.Rooms {
		if r[0] == '.' && r[1] == '.' && r[2] == '.' && r[3] == '.' {
			continue
		} else if r[0] == r[1] && r[0] == r[2] && r[0] == r[3] && r[0] == ExpectedRooms[i] {
			continue
		} else if r[0] == '.' && r[1] == ExpectedRooms[i] && r[2] == ExpectedRooms[i] && r[3] == ExpectedRooms[i] {
			continue
		} else if r[0] == '.' && r[1] == '.' && r[2] == ExpectedRooms[i] && r[3] == ExpectedRooms[i] {
			continue
		} else if r[0] == '.' && r[1] == '.' && r[2] == '.' && r[3] == ExpectedRooms[i] {
			continue
		}
		j := 0
		for r[j] == '.' {
			j++
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
			res := make(chan Result)
			go Solve(Input{Hallway: newHall, Rooms: newRooms}, total+energy, res)
			for result := range res {
				if result.Valid {
					if result.Energy < min {
						min = result.Energy
					}
					valid = true
				}

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
			res := make(chan Result)
			go Solve(Input{Hallway: newHall, Rooms: newRooms}, total+energy, res)
			for result := range res {
				if result.Valid {
					if result.Energy < min {
						min = result.Energy
					}
					valid = true
				}

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
		} else if in.Rooms[r][2] != h && in.Rooms[r][2] != '.' {
			continue
		} else if in.Rooms[r][3] != h && in.Rooms[r][3] != '.' {
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
		j := 0
		for j < 4 && in.Rooms[r][j] == '.' {
			j++
		}
		j--
		newHall := in.Hallway
		newRooms := in.Rooms
		newHall[i] = '.'
		newRooms[r][j] = h
		energy := (helpers.Abs(BarredSpaces[r]-i) + j + 1) * EnergyMap[h]
		res := make(chan Result)
		go Solve(Input{Hallway: newHall, Rooms: newRooms}, total+energy, res)
		for result := range res {
			if result.Valid {
				if result.Energy < min {
					min = result.Energy
				}
				valid = true
			}

		}
	}

	mu.Lock()
	memo[in] = CacheValue{StartEnergy: total, Valid: valid}
	mu.Unlock()

	results <- Result{Energy: min, Valid: valid}
}

func partOne() {
	fmt.Printf("Part 1: See old commit\n")
}

func partTwo() {
	input := readInput()

	ans := math.MaxInt32
	results := make(chan Result)
	go Solve(input, 0, results)

	for res := range results {
		if res.Valid {
			if res.Energy < ans {
				ans = res.Energy
			}
		}
	}

	fmt.Printf("Part 2: %d\n", ans)
}

func main() {
	partOne()
	partTwo()
}
