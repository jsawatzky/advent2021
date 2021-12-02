package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var hor, depth int64

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) != 2 {
			log.Fatal("Incorrect line")
		}
		amt, err := strconv.ParseInt(line[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		switch line[0] {
		case "up":
			depth -= amt
		case "down":
			depth += amt
		case "forward":
			hor += amt
		default:
			log.Fatal("Invailid order")
		}
	}

	fmt.Printf("Part 1: %d\n", hor*depth)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var hor, depth, aim int64

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) != 2 {
			log.Fatal("Incorrect line")
		}
		amt, err := strconv.ParseInt(line[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		switch line[0] {
		case "up":
			aim -= amt
		case "down":
			aim += amt
		case "forward":
			hor += amt
			depth += aim * amt
		default:
			log.Fatal("Invailid order")
		}
	}

	fmt.Printf("Part 2: %d\n", hor*depth)
}

func main() {
	partOne()
	partTwo()
}
