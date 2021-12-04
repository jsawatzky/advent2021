package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func In(val int, arr []int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

type BingoCard [][]int

func IsWinner(c BingoCard, marked []int) bool {
	cols := make([]bool, len(c[0]))
	for i := range cols {
		cols[i] = true
	}

	for _, row := range c {
		rowWin := true
		for j, col := range row {
			if !In(col, marked) {
				cols[j] = false
				rowWin = false
			}
		}
		if rowWin {
			return true
		}
	}

	for _, col := range cols {
		if col {
			return true
		}
	}

	return false
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	marksStr := strings.Split(strings.TrimSpace(scanner.Text()), ",")
	marks := make([]int, 0, len(marksStr))
	for _, m := range marksStr {
		i, err := strconv.Atoi(m)
		if err != nil {
			panic(err)
		}
		marks = append(marks, i)
	}

	boards := make([]BingoCard, 0, 20)

	for scanner.Scan() {
		card := make(BingoCard, 5)
		for i := 0; i < 5; i++ {
			scanner.Scan()
			cardRowStr := strings.Fields(strings.TrimSpace(scanner.Text()))
			card[i] = make([]int, 0, len(cardRowStr))
			for _, m := range cardRowStr {
				j, err := strconv.Atoi(m)
				if err != nil {
					panic(err)
				}
				card[i] = append(card[i], j)
			}
		}
		boards = append(boards, card)
	}

	for i := range marks {
		curMarks := marks[:i+1]
		for _, b := range boards {
			if IsWinner(b, curMarks) {
				var sum int
				for _, row := range b {
					for _, val := range row {
						if !In(val, curMarks) {
							sum += val
						}
					}
				}
				fmt.Printf("Part 1: %d\n", sum*marks[i])
				return
			}
		}
	}
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	marksStr := strings.Split(strings.TrimSpace(scanner.Text()), ",")
	marks := make([]int, 0, len(marksStr))
	for _, m := range marksStr {
		i, err := strconv.Atoi(m)
		if err != nil {
			panic(err)
		}
		marks = append(marks, i)
	}

	boards := make([]BingoCard, 0, 20)

	for scanner.Scan() {
		card := make(BingoCard, 5)
		for i := 0; i < 5; i++ {
			scanner.Scan()
			cardRowStr := strings.Fields(strings.TrimSpace(scanner.Text()))
			card[i] = make([]int, 0, len(cardRowStr))
			for _, m := range cardRowStr {
				j, err := strconv.Atoi(m)
				if err != nil {
					panic(err)
				}
				card[i] = append(card[i], j)
			}
		}
		boards = append(boards, card)
	}

	var lastWin int
	for i := range marks {
		curMarks := marks[:i+1]
	checkWinner:
		for j, b := range boards {
			if IsWinner(b, curMarks) {
				var sum int
				for _, row := range b {
					for _, val := range row {
						if !In(val, curMarks) {
							sum += val
						}
					}
				}
				lastWin = sum * marks[i]
				if j+1 < len(boards) {
					boards = append(boards[:j], boards[j+1:]...)
				} else {
					boards = boards[:j]
				}
				goto checkWinner
			}
		}
	}

	fmt.Printf("Part 2: %d\n", lastWin)
}

func main() {
	partOne()
	partTwo()
}
