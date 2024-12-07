package main

import (
	"bufio"
	"math"
	"strconv"
	"strings"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

type Coordinate [3]int

func parseLine(line string) Coordinate {
	d := strings.Split(line, ",")
	x, _ := strconv.Atoi(d[0])
	y, _ := strconv.Atoi(d[1])
	z, _ := strconv.Atoi(d[2])
	return Coordinate{x, y, z}
}

func part1(scanner *bufio.Scanner) interface{} {
	lava := make(map[Coordinate]bool)
	for scanner.Scan() {
		lava[parseLine(scanner.Text())] = true
	}
	surfaceArea := 0
	delta := []int{-1, 1}
	for l := range lava {
		area := 6
		for i := 0; i < 3; i++ {
			for _, d := range delta {
				l[i] += d
				if lava[l] {
					area -= 1
				}
				l[i] -= d
			}
		}
		surfaceArea += area
	}
	return surfaceArea
}

func part2(scanner *bufio.Scanner) interface{} {
	lava := make(map[Coordinate]bool)
	minBound, maxBound := Coordinate{math.MaxInt, math.MaxInt, math.MaxInt}, Coordinate{}
	for scanner.Scan() {
		c := parseLine(scanner.Text())
		lava[c] = true
		for i := 0; i < 3; i++ {
			if c[i] < minBound[i] {
				minBound[i] = c[i]
			} else if c[i] > maxBound[i] {
				maxBound[i] = c[i]
			}
		}
	}

	for i := 0; i < 3; i++ {
		minBound[i] -= 1
		maxBound[i] += 1
	}

	outside := make(map[Coordinate]bool)
	queue := []Coordinate{minBound}
	outside[minBound] = true

	delta := []int{-1, 1}
	for len(queue) > 0 {
		for i := len(queue); i > 0; i-- {
			c := queue[0]
			queue = queue[1:]
			for i := 0; i < 3; i++ {
				for _, d := range delta {
					c[i] += d
					if c[i] >= minBound[i] && c[i] <= maxBound[i] && !outside[c] && !lava[c] {
						outside[c] = true
						queue = append(queue, c)
					}
					c[i] -= d
				}
			}
		}
	}

	surfaceArea := 0
	for l := range lava {
		area := 0
		for i := 0; i < 3; i++ {
			for _, d := range delta {
				l[i] += d
				if outside[l] {
					area += 1
				}
				l[i] -= d
			}
		}
		surfaceArea += area
	}
	return surfaceArea
}
