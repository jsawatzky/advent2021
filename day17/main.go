package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jsawatzky/advent/helpers"
)

type Target struct {
	MinX int
	MaxX int
	MinY int
	MaxY int
}

func readInput() Target {
	input := helpers.ReadInput()
	input = strings.TrimPrefix(input, "target area: x=")
	area := strings.Split(input, ", ")
	xRange := helpers.AtoiArr(strings.Split(area[0], ".."))
	yRange := helpers.AtoiArr(strings.Split(strings.TrimPrefix(area[1], "y="), ".."))
	return Target{MinX: xRange[0], MaxX: xRange[1], MinY: yRange[0], MaxY: yRange[1]}
}

func partOne() {
	target := readInput()

	maxVert := (target.MinY * -1) - 1

	ans := maxVert * (maxVert + 1) / 2

	fmt.Printf("Part 1: %d\n", ans)
}

func partTwo() {
	target := readInput()

	maxVert := (target.MinY * -1) - 1
	minHorz := int(math.Ceil(0.5 * (math.Sqrt(float64(8*target.MinX+1)) - 1)))

	ans := 0

	for xVel := minHorz; xVel <= target.MaxX; xVel++ {
	outer:
		for yVel := target.MinY; yVel <= maxVert; yVel++ {
			xPos := 0
			xVelTmp := xVel
			yPos := 0
			yVelTmp := yVel
			for xPos <= target.MaxX && yPos >= target.MinY {
				if xPos >= target.MinX && yPos <= target.MaxY {
					ans++
					continue outer
				}
				xPos += xVelTmp
				yPos += yVelTmp
				if xVelTmp > 0 {
					xVelTmp--
				}
				yVelTmp--
			}
		}
	}

	fmt.Printf("Part 2: %d\n", ans)
}

func main() {
	partOne()
	partTwo()
}
