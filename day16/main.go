package main

import (
	"fmt"
	"math"

	"github.com/jsawatzky/advent/helpers"
)

var hexToBin = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

type Packet struct {
	Version    int
	TypeID     int
	Value      int
	SubPackets []Packet
}

func parsePacket(data string) (Packet, string) {
	p := Packet{}
	p.Version = helpers.BinaryToInt(data[:3])
	p.TypeID = helpers.BinaryToInt(data[3:6])
	data = data[6:]
	switch p.TypeID {
	case 4:
		val := ""
		for data[0] == '1' {
			val += data[1:5]
			data = data[5:]
		}
		val += data[1:5]
		data = data[5:]
		p.Value = helpers.BinaryToInt(val)
	default:
		lenType := data[0]
		data = data[1:]
		if lenType == '0' {
			totalLen := helpers.BinaryToInt(data[:15])
			data = data[15:]
			packetData := data[:totalLen]
			data = data[totalLen:]
			subPackets := []Packet{}
			for len(packetData) > 0 {
				var sp Packet
				sp, packetData = parsePacket(packetData)
				subPackets = append(subPackets, sp)
			}
			p.SubPackets = subPackets
		} else {
			numPackets := helpers.BinaryToInt(data[:11])
			data = data[11:]
			subPackets := []Packet{}
			for i := 0; i < numPackets; i++ {
				var sp Packet
				sp, data = parsePacket(data)
				subPackets = append(subPackets, sp)
			}
			p.SubPackets = subPackets
		}
	}

	return p, data
}

func SumVersions(p Packet) int {
	ret := p.Version
	for _, sp := range p.SubPackets {
		ret += SumVersions(sp)
	}
	return ret
}

func CalcValue(p Packet) int {
	switch p.TypeID {
	case 0:
		total := 0
		for _, sp := range p.SubPackets {
			total += CalcValue(sp)
		}
		return total
	case 1:
		total := 1
		for _, sp := range p.SubPackets {
			total *= CalcValue(sp)
		}
		return total
	case 2:
		min := math.MaxInt64
		for _, sp := range p.SubPackets {
			val := CalcValue(sp)
			if val < min {
				min = val
			}
		}
		return min
	case 3:
		max := math.MinInt64
		for _, sp := range p.SubPackets {
			val := CalcValue(sp)
			if val > max {
				max = val
			}
		}
		return max
	case 4:
		return p.Value
	case 5:
		if len(p.SubPackets) != 2 {
			panic("too many packets")
		}
		v1 := CalcValue(p.SubPackets[0])
		v2 := CalcValue(p.SubPackets[1])
		if v1 > v2 {
			return 1
		}
		return 0
	case 6:
		if len(p.SubPackets) != 2 {
			panic("too many packets")
		}
		v1 := CalcValue(p.SubPackets[0])
		v2 := CalcValue(p.SubPackets[1])
		if v1 < v2 {
			return 1
		}
		return 0
	case 7:
		if len(p.SubPackets) != 2 {
			panic("too many packets")
		}
		v1 := CalcValue(p.SubPackets[0])
		v2 := CalcValue(p.SubPackets[1])
		if v1 == v2 {
			return 1
		}
		return 0
	default:
		panic("invalid type")
	}
}

func readInput() string {
	input := helpers.ReadInput()
	result := ""
	for _, c := range input {
		result += hexToBin[c]
	}
	return result
}

func partOne() {
	input := readInput()

	packet, left := parsePacket(input)
	if len(left) > 0 {
		for _, c := range left {
			if c != '0' {
				panic("leftover packet")
			}
		}
	}

	ans := SumVersions(packet)

	fmt.Printf("Part 1: %d\n", ans)
}

func partTwo() {
	input := readInput()

	packet, left := parsePacket(input)
	if len(left) > 0 {
		for _, c := range left {
			if c != '0' {
				panic("leftover packet")
			}
		}
	}

	ans := CalcValue(packet)

	fmt.Printf("Part 2: %d\n", ans)
}

func main() {
	partOne()
	partTwo()
}
