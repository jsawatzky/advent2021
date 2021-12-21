package main

import (
	"fmt"

	"github.com/jsawatzky/advent/helpers"
)

const (
	P1Start = 10
	P2Start = 1
)

func partOne() {

	p1score := 0
	p2score := 0

	p1 := P1Start
	p2 := P2Start
	round := 0
	for p1score < 1000 && p2score < 1000 {
		roll := (round * 9) + 6
		if round%2 == 0 {
			p1 += roll
			for p1 > 10 {
				p1 -= 10
			}
			p1score += p1
		} else if round%2 == 1 {
			p2 += roll
			for p2 > 10 {
				p2 -= 10
			}
			p2score += p2
		}
		round++
	}

	ans := round * 3 * helpers.Min(p1score, p2score)

	fmt.Printf("Part 1: %d\n", ans)
}

var OutcomeOccurances = map[int]int64{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

type Round struct {
	Turn             int
	P1Score, P2Score int
	P1Pos, P2Pos     int
}

type Wins struct {
	P1, P2 int64
}

var memo map[Round]Wins = make(map[Round]Wins)

func Play(turn, p1score, p2score, p1pos, p2pos int) Wins {
	r := Round{Turn: turn, P1Score: p1score, P2Score: p2score, P1Pos: p1pos, P2Pos: p2pos}
	if c, ok := memo[r]; ok {
		return c
	}

	if p1score >= 21 {
		return Wins{P1: 1, P2: 0}
	} else if p2score >= 21 {
		return Wins{P1: 0, P2: 1}
	}

	var w Wins
	if turn == 1 {
		for d, occ := range OutcomeOccurances {
			p1posNew := p1pos + d
			p1posNew %= 10
			p1scoreNew := p1score + p1posNew + 1
			w2 := Play(2, p1scoreNew, p2score, p1posNew, p2pos)
			w.P1 += w2.P1 * occ
			w.P2 += w2.P2 * occ
		}
	} else {
		for d, occ := range OutcomeOccurances {
			p2posNew := p2pos + d
			p2posNew %= 10
			p2scoreNew := p2score + p2posNew + 1
			w2 := Play(1, p1score, p2scoreNew, p1pos, p2posNew)
			w.P1 += w2.P1 * occ
			w.P2 += w2.P2 * occ
		}
	}
	memo[r] = w
	return w
}

func partTwo() {

	wins := Play(1, 0, 0, P1Start-1, P2Start-1)

	var ans int64
	if wins.P1 > wins.P2 {
		ans = wins.P1
	} else {
		ans = wins.P2
	}

	fmt.Printf("Part 2: %d\n", ans)
}

func main() {
	partOne()
	partTwo()
}
