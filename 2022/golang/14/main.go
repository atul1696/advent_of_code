package main

import (
	"bufio"
	"math"
	"strconv"
	"strings"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

type Coordinate struct {
	X, Y int
}

func main() {
	utils.Run(part1, part2)
}

func parseGrid(scanner *bufio.Scanner) (map[Coordinate]bool, int, int, int) {
	data := []*[]Coordinate{}
	minX, maxX := math.MaxInt32, 0
	maxY := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := []Coordinate{}
		for _, pos := range strings.Split(line, " -> ") {
			i := strings.Split(pos, ",")
			x, _ := strconv.Atoi(i[0])
			y, _ := strconv.Atoi(i[1])
			if x < minX {
				minX = x
			} else if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
			row = append(row, Coordinate{X: x, Y: y})
		}
		data = append(data, &row)
	}
	grid := make(map[Coordinate]bool)
	for _, line := range data {
		row := *line
		for k := 1; k < len(row); k++ {
			if row[k].X == row[k-1].X {
				start := row[k-1].Y
				end := row[k].Y
				if row[k].Y < start {
					start = row[k].Y
					end = row[k-1].Y
				}
				for i := start; i <= end; i++ {
					grid[Coordinate{X: row[k].X, Y: i}] = true
				}
			} else {
				start := row[k-1].X
				end := row[k].X
				if row[k].X < start {
					start = row[k].X
					end = row[k-1].X
				}
				for j := start; j <= end; j++ {
					grid[Coordinate{X: j, Y: row[k].Y}] = true
				}
			}
		}
	}
	return grid, minX, maxX, maxY
}

func part1(scanner *bufio.Scanner) interface{} {
	grid, minX, maxX, maxY := parseGrid(scanner)
	isFull := false
	sand := 0
	for !isFull {
		x := 500
		y := 0
		stop := false
		for !stop {
			if x-1 < minX || x+1 > maxX || y+1 > maxY {
				isFull = true
				break
			}
			if !grid[Coordinate{X: x, Y: y + 1}] {
				y += 1
			} else if !grid[Coordinate{X: x - 1, Y: y + 1}] {
				x -= 1
				y += 1
			} else if !grid[Coordinate{X: x + 1, Y: y + 1}] {
				x += 1
				y += 1
			} else {
				stop = true
			}
		}
		grid[Coordinate{X: x, Y: y}] = true
		if !isFull {
			sand += 1
		}
	}
	return sand
}

func part2(scanner *bufio.Scanner) interface{} {
	grid, _, _, maxY := parseGrid(scanner)
	floor := maxY + 2
	sand := 0
	for !grid[Coordinate{X: 500, Y: 0}] {
		x := 500
		y := 0
		stop := false
		for !stop && y+1 < floor {
			if !grid[Coordinate{X: x, Y: y + 1}] {
				y += 1
			} else if !grid[Coordinate{X: x - 1, Y: y + 1}] {
				x -= 1
				y += 1
			} else if !grid[Coordinate{X: x + 1, Y: y + 1}] {
				x += 1
				y += 1
			} else {
				stop = true
			}
		}
		grid[Coordinate{X: x, Y: y}] = true
		sand += 1
	}
	return sand
}
