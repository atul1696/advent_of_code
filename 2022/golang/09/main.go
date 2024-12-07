package main

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

type position struct {
	x, y int
}

func run(scanner *bufio.Scanner, length int) int {
	rope := make([]position, length)
	visited := make(map[position]bool)
	visited[rope[length-1]] = true

	for scanner.Scan() {
		move := strings.Split(scanner.Text(), " ")
		var dir position
		switch move[0] {
		case "U":
			dir = position{0, 1}
		case "D":
			dir = position{0, -1}
		case "L":
			dir = position{-1, 0}
		case "R":
			dir = position{1, 0}
		}
		count, _ := strconv.Atoi(move[1])
		for ; count > 0; count-- {
			rope[0].x += dir.x
			rope[0].y += dir.y
			for i := 0; i < length-1; i++ {
				diffX := rope[i].x - rope[i+1].x
				diffY := rope[i].y - rope[i+1].y
				updateX := diffX > 1 || diffX < -1
				updateY := diffY > 1 || diffY < -1
				updateX = updateX || (updateY && diffX != 0)
				updateY = updateY || (updateX && diffY != 0)

				if updateX {
					if diffX > 0 {
						rope[i+1].x += 1
					} else {
						rope[i+1].x -= 1
					}
				}
				if updateY {
					if diffY > 0 {
						rope[i+1].y += 1
					} else {
						rope[i+1].y -= 1
					}
				}
			}
			visited[rope[length-1]] = true
		}
	}
	return len(visited)
}

func part1(scanner *bufio.Scanner) interface{} {
	return run(scanner, 2)
}

func part2(scanner *bufio.Scanner) interface{} {
	return run(scanner, 10)
}
