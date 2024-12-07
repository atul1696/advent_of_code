package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

func part1(scanner *bufio.Scanner) interface{} {
	totalStrength := 0
	prev, curr := 1, 1
	cycles := 0
	isNoop := false
	for scanner.Scan() {
		switch cycles {
		case 20, 60, 100, 140, 180, 220:
			{
				if isNoop {
					totalStrength += curr * cycles
				} else {
					totalStrength += prev * cycles
				}
			}
		case 21, 61, 101, 141, 181, 221:
			{
				if !isNoop {
					totalStrength += prev * (cycles - 1)
				}
			}
		}
		op := scanner.Text()
		isNoop = op == "noop"
		cycles += 1
		if !isNoop {
			cycles += 1
			val, _ := strconv.Atoi(op[5:])
			prev = curr
			curr += val
		} else {
			prev = curr
		}
	}
	return totalStrength
}

func part2(scanner *bufio.Scanner) interface{} {
	width, height := 40, 6
	crt := make([][]string, height)
	for i := range crt {
		crt[i] = make([]string, width)
		for j := range crt[i] {
			crt[i][j] = "."
		}
	}
	x := 1
	nextUpdate := 0
	nextVal := x
	for cycle := 0; cycle < width*height; cycle++ {
		if nextUpdate == cycle {
			x = nextVal
			scanner.Scan()
			op := scanner.Text()
			isNoop := op == "noop"
			if !isNoop {
				nextUpdate = cycle + 2
				val, _ := strconv.Atoi(op[5:])
				nextVal = x + val
			} else {
				nextUpdate = cycle + 1
				nextVal = x
			}
		}
		row, col := cycle/40, cycle%40
		diff := x - col
		if diff >= -1 && diff <= 1 {
			crt[row][col] = "#"
		}
	}
	for _, row := range crt {
		fmt.Println(strings.Join(row, ""))
	}
	return 0
}
