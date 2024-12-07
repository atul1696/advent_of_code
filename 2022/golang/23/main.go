package main

import (
	"bufio"
	"math"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

type Position struct {
	x, y int
}

func simulate(elves map[Position]bool, directions []string) (map[Position]bool, []string, bool) {
	proposed := make(map[Position]Position)
	proposals := make(map[Position]int)
	for currPos := range elves {
		skip := true
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				} else {
					currPos.x += i
					currPos.y += j

					if elves[currPos] {
						skip = false
					}

					currPos.x -= i
					currPos.y -= j
				}
			}
		}
		if skip {
			continue
		}
		for _, dir := range directions {
			possible := true
			var newPos Position
			switch dir {
			case "N", "S":
				{
					if dir == "N" {
						newPos = Position{currPos.x - 1, currPos.y}
					} else {
						newPos = Position{currPos.x + 1, currPos.y}
					}
					for j := -1; j <= 1; j++ {
						newPos.y += j
						if elves[newPos] {
							possible = false
						}
						newPos.y -= j
					}
				}
			case "W", "E":
				{
					if dir == "W" {
						newPos = Position{currPos.x, currPos.y - 1}
					} else {
						newPos = Position{currPos.x, currPos.y + 1}
					}

					for i := -1; i <= 1; i++ {
						newPos.x += i
						if elves[newPos] {
							possible = false
						}
						newPos.x -= i
					}
				}
			}
			if possible {
				proposed[currPos] = newPos
				proposals[newPos] += 1
				break
			}
		}
	}
	updated := make(map[Position]bool)
	for currPos := range elves {
		if newPos, ok := proposed[currPos]; ok {
			if proposals[newPos] == 1 {
				updated[newPos] = true
				continue
			}
		}
		updated[currPos] = true
	}
	elves = updated
	directions = append(directions, directions[0])[1:]
	return elves, directions, len(proposals) == 0
}

func part1(scanner *bufio.Scanner) interface{} {
	elves := make(map[Position]bool)
	i := 0
	for scanner.Scan() {
		for j, c := range scanner.Text() {
			if string(c) == "#" {
				elves[Position{i, j}] = true
			}
		}
		i += 1
	}

	directions := []string{"N", "S", "W", "E"}
	for r := 0; r < 10; r++ {
		elves, directions, _ = simulate(elves, directions)
	}
	minX, maxX, minY, maxY := math.MaxInt16, math.MinInt16, math.MaxInt16, math.MinInt16
	for pos := range elves {
		if pos.x < minX {
			minX = pos.x
		} else if pos.x > maxX {
			maxX = pos.x
		}

		if pos.y < minY {
			minY = pos.y
		} else if pos.y > maxY {
			maxY = pos.y
		}
	}
	return (maxX-minX+1)*(maxY-minY+1) - len(elves)
}

func part2(scanner *bufio.Scanner) interface{} {
	elves := make(map[Position]bool)
	i := 0
	for scanner.Scan() {
		for j, c := range scanner.Text() {
			if string(c) == "#" {
				elves[Position{i, j}] = true
			}
		}
		i += 1
	}

	directions := []string{"N", "S", "W", "E"}
	stop := false
	r := 0
	for r = 0; !stop; r++ {
		elves, directions, stop = simulate(elves, directions)
	}
	return r
}
