package main

import (
	"fmt"
	"strings"

	"github.com/jsawatzky/advent/helpers"
)

type Position struct {
	X int
	Y int
	Z int
}

func (p Position) Add(p2 Position) Position {
	return Position{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
		Z: p.Z + p2.Z,
	}
}

func (p Position) Sub(p2 Position) Position {
	return Position{
		X: p.X - p2.X,
		Y: p.Y - p2.Y,
		Z: p.Z - p2.Z,
	}
}

type Orientation struct {
	CoordMap [3]rune
	CoordDir [3]int
}

func (o Orientation) Transform(p Position) Position {
	x, y, z := p.X, p.Y, p.Z
	var newX, newY, newZ int

	switch o.CoordMap[0] {
	case 'x':
		newX = x
	case 'y':
		newY = x
	case 'z':
		newZ = x
	}
	switch o.CoordMap[1] {
	case 'x':
		newX = y
	case 'y':
		newY = y
	case 'z':
		newZ = y
	}
	switch o.CoordMap[2] {
	case 'x':
		newX = z
	case 'y':
		newY = z
	case 'z':
		newZ = z
	}

	newX *= o.CoordDir[0]
	newY *= o.CoordDir[1]
	newZ *= o.CoordDir[2]

	return Position{X: newX, Y: newY, Z: newZ}

}

var Orientations []Orientation = []Orientation{}

func init() {
	axis := []rune{'x', 'y', 'z'}
	dirs := []int{-1, 1}
	for _, a1 := range axis {
		for _, a2 := range axis {
			if a1 == a2 {
				continue
			}
			for _, a3 := range axis {
				if a1 == a3 || a2 == a3 {
					continue
				}
				for _, d1 := range dirs {
					for _, d2 := range dirs {
						for _, d3 := range dirs {
							Orientations = append(Orientations, Orientation{CoordMap: [3]rune{a1, a2, a3}, CoordDir: [3]int{d1, d2, d3}})
						}
					}
				}
			}
		}
	}
}

type Scanner struct {
	Mapped  bool
	Beacons []Beacon
}

type Beacon struct {
	Pos Position
}

func readInput() []Scanner {
	lines := helpers.ReadInputLines()
	scanners := []Scanner{}
	var curScanner Scanner
	for _, l := range lines {
		if strings.HasPrefix(l, "---") {
			continue
		}
		if len(l) == 0 {
			scanners = append(scanners, curScanner)
			curScanner = Scanner{}
			continue
		}

		b := helpers.AtoiArr(strings.Split(l, ","))
		curScanner.Beacons = append(curScanner.Beacons, Beacon{Pos: Position{
			X: b[0],
			Y: b[1],
			Z: b[2],
		}})
	}
	scanners = append(scanners, curScanner)

	return scanners
}

func In(b Beacon, beacons []Beacon) bool {
	for _, v := range beacons {
		if b == v {
			return true
		}
	}
	return false
}

var postitions []Position = []Position{}

func partOne() {

	scanners := readInput()

	mainScanner := scanners[0]
	scanners = scanners[1:]

	beacons := mainScanner.Beacons
	numMapped := 0

	for numMapped < len(scanners) {
		fmt.Println(numMapped)
	inner:
		for i, s := range scanners {
			if s.Mapped {
				continue
			}
			maxCount := 0
			for _, b1 := range beacons {
				for _, b2 := range s.Beacons {
					for _, o := range Orientations {
						newP := o.Transform(b2.Pos)
						d := b1.Pos.Sub(newP)
						count := 1
						for _, b3 := range beacons {
							if b1 == b3 {
								continue
							}
							for _, b4 := range s.Beacons {
								if b2 == b4 {
									continue
								}
								if b3.Pos.Sub(o.Transform(b4.Pos)) == d {
									count++
								}
							}
						}
						if count >= 12 {
							postitions = append(postitions, d)
							for _, newB := range s.Beacons {
								newB.Pos = o.Transform(newB.Pos)
								newB.Pos = newB.Pos.Add(d)
								if !In(newB, beacons) {
									beacons = append(beacons, newB)
								}
							}
							numMapped++
							scanners[i].Mapped = true
							continue inner
						} else if count > maxCount {
							maxCount = count
						}
					}
				}
			}
		}
	}

	ans := len(beacons)

	fmt.Printf("Part 1: %v\n", ans)
}

func partTwo() {

	maxDist := 0
	for _, p1 := range postitions {
		for _, p2 := range postitions {
			if p1 == p2 {
				continue
			}
			d := helpers.Abs(p1.X-p2.X) + helpers.Abs(p1.Y-p2.Y) + helpers.Abs(p1.Z-p2.Z)
			if d > maxDist {
				maxDist = d
			}
		}
	}

	ans := maxDist

	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	partOne()
	partTwo()
}
