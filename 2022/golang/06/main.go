package main

import (
	"bufio"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

func getStart(signal string, windowSize int) int {
	rune_map := make(map[rune]int)
	start := 0
	for i, r := range signal {
		rune_map[r] += 1
		if i >= windowSize {
			prev := rune(signal[i-windowSize])
			rune_map[prev] -= 1
			if rune_map[prev] == 0 {
				delete(rune_map, prev)
			}
		}
		if len(rune_map) == windowSize {
			start = i + 1
			break
		}
	}
	return start
}

func part1(scanner *bufio.Scanner) interface{} {
	scanner.Scan()
	signal := scanner.Text()
	return getStart(signal, 4)
}

func part2(scanner *bufio.Scanner) interface{} {
	scanner.Scan()
	signal := scanner.Text()
	return getStart(signal, 14)
}
