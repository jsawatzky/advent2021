package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/jsawatzky/advent/helpers"
)

type Cave struct {
	Name        string
	Big         bool
	Connections []string
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

var AllPaths [][]string = make([][]string, 0, 10)

func FindPaths(cur string, curPath []string, caves map[string]*Cave, smallVisit bool) {
	curPath = append(curPath, cur)
	for _, c := range caves[cur].Connections {
		if c == "end" {
			AllPaths = append(AllPaths, curPath)
			continue
		}
		nextCave := caves[c]
		if nextCave.Big || !helpers.InStr(c, curPath) {
			FindPaths(c, curPath, caves, smallVisit)
		} else if !smallVisit && c != "start" {
			FindPaths(c, curPath, caves, true)
		}
	}
}

func partOne() {
	lines := helpers.ReadInputLines()

	caves := make(map[string]*Cave)
	for _, l := range lines {
		conn := strings.Split(l, "-")
		c1, ok := caves[conn[0]]
		if !ok {
			caves[conn[0]] = &Cave{
				Name: conn[0],
				Big:  IsUpper(conn[0]),
			}
			c1 = caves[conn[0]]
		}
		c2, ok := caves[conn[1]]
		if !ok {
			caves[conn[1]] = &Cave{
				Name: conn[1],
				Big:  IsUpper(conn[1]),
			}
			c2 = caves[conn[1]]
		}
		c1.Connections = append(c1.Connections, c2.Name)
		c2.Connections = append(c2.Connections, c1.Name)
	}

	FindPaths("start", []string{}, caves, true)

	ans := len(AllPaths)

	fmt.Printf("Part 1: %d\n", ans)
}

func partTwo() {
	lines := helpers.ReadInputLines()

	caves := make(map[string]*Cave)
	for _, l := range lines {
		conn := strings.Split(l, "-")
		c1, ok := caves[conn[0]]
		if !ok {
			caves[conn[0]] = &Cave{
				Name: conn[0],
				Big:  IsUpper(conn[0]),
			}
			c1 = caves[conn[0]]
		}
		c2, ok := caves[conn[1]]
		if !ok {
			caves[conn[1]] = &Cave{
				Name: conn[1],
				Big:  IsUpper(conn[1]),
			}
			c2 = caves[conn[1]]
		}
		c1.Connections = append(c1.Connections, c2.Name)
		c2.Connections = append(c2.Connections, c1.Name)
	}

	FindPaths("start", []string{}, caves, false)

	ans := len(AllPaths)

	fmt.Printf("Part 2: %d\n", ans)
}

func main() {
	partOne()
	AllPaths = make([][]string, 0, 10)
	partTwo()
}
