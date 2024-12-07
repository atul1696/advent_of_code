package main

import (
	"bufio"
	"strings"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

func score1(p1, p2 string) int {
	score := 0
	shape := 1
	if p2 == "X" {
		shape = 1
		if p1 == "A" {
			score = 3
		} else if p1 == "C" {
			score = 6
		}
	} else if p2 == "Y" {
		shape = 2
		if p1 == "B" {
			score = 3
		} else if p1 == "A" {
			score = 6
		}
	} else {
		shape = 3
		if p1 == "C" {
			score = 3
		} else if p1 == "B" {
			score = 6
		}
	}
	return score + shape
}

func part1(scanner *bufio.Scanner) interface{} {
	total := 0
	for scanner.Scan() {
		choice := strings.Split(scanner.Text(), " ")
		total += score1(choice[0], choice[1])
	}
	return total
}

func score2(p1, outcome string) int {
	score := 0
	shape := 1
	if outcome == "X" {
		if p1 == "A" {
			shape = 3
		} else if p1 == "C" {
			shape = 2
		}
	} else if outcome == "Y" {
		score = 3
		if p1 == "B" {
			shape = 2
		} else if p1 == "C" {
			shape = 3
		}
	} else {
		score = 6
		if p1 == "A" {
			shape = 2
		} else if p1 == "B" {
			shape = 3
		}
	}
	return score + shape
}

func part2(scanner *bufio.Scanner) interface{} {
	total := 0
	for scanner.Scan() {
		choice := strings.Split(scanner.Text(), " ")
		total += score2(choice[0], choice[1])
	}
	return total
}
