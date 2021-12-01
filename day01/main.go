package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	prevLevel, err := strconv.ParseInt(scanner.Text(), 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	for scanner.Scan() {
		lvl, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		if lvl > prevLevel {
			i++
		}
		prevLevel = lvl
	}

	fmt.Printf("Part 1: %d\n", i)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	vals := make(chan int64, 3)

	var curSum int64

	scanner := bufio.NewScanner(file)
	for i := 0; i < 3; i++ {
		scanner.Scan()
		lvl, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		vals <- lvl
		curSum += lvl
	}

	var i int

	for scanner.Scan() {
		prevSum := curSum
		prevLvl := <-vals
		lvl, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		vals <- lvl
		curSum += lvl - prevLvl
		if curSum > prevSum {
			i++
		}
	}

	fmt.Printf("Part 2: %d\n", i)
}

func main() {
	partOne()
	partTwo()
}
