package main

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

func turn(dx, dy int, dir string) (int, int) {
	if dir == "L" {
		if dx == 0 {
			dx = -dy
			dy = 0
		} else {
			dy = dx
			dx = 0
		}
	} else if dir == "R" {
		if dx == 0 {
			dx = dy
			dy = 0
		} else {
			dy = -dx
			dx = 0
		}
	}
	return dx, dy
}

func facingScore(dx, dy int) int {
	if dx == 0 && dy == 1 {
		return 0
	} else if dx == 1 && dy == 0 {
		return 1
	} else if dx == 0 && dy == -1 {
		return 2
	} else {
		return 3
	}
}

func part1(scanner *bufio.Scanner) interface{} {
	grid := [][]string{}
	rowBoundary := make([][]int, 0)
	var path string
	cols := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			scanner.Scan()
			path = scanner.Text()
		}
		grid = append(grid, strings.Split(line, ""))
		start, end := -1, -1
		for j, c := range line {
			if len(line) > cols {
				cols = len(line)
			}
			if string(c) != " " {
				if start == -1 {
					start = j
				}
				end = j
			}
		}
		rowBoundary = append(rowBoundary, []int{start, end})
	}

	colBoundary := make([][]int, 0)
	for j := 0; j < cols; j++ {
		start, end := -1, -1
		for i := 0; i < len(grid); i++ {
			if len(grid[i]) <= j {
				continue
			}
			if string(grid[i][j]) != " " {
				if start == -1 {
					start = i
				}
				end = i
			}
		}
		colBoundary = append(colBoundary, []int{start, end})
	}

	dx, dy := 0, 1
	x, y := 0, 0
	for j, c := range grid[0] {
		if string(c) == "." {
			y = j
			break
		}
	}
	r := regexp.MustCompile(`(\d+)(\w)?`)
	for _, match := range r.FindAllStringSubmatch(path, -1) {
		dist, _ := strconv.Atoi(match[1])
		for ; dist > 0; dist-- {
			i, j := x+dx, y+dy
			if dx == 0 {
				if rowBoundary[x][0] <= j && j <= rowBoundary[x][1] {
					if grid[i][j] == "#" {
						break
					}
				} else {
					c := y
					if dy == 1 {
						c = rowBoundary[x][0]
					} else {
						c = rowBoundary[x][1]
					}
					if grid[x][c] == "#" {
						break
					}
					j = c
				}
			} else {
				if colBoundary[y][0] <= i && i <= colBoundary[y][1] {
					if grid[i][j] == "#" {
						break
					}
				} else {
					r := x
					if dx == 1 {
						r = colBoundary[y][0]
					} else {
						r = colBoundary[y][1]
					}
					if grid[r][y] == "#" {
						break
					} else {
						i = r
					}
				}
			}
			x, y = i, j
		}
		dx, dy = turn(dx, dy, match[2])
	}
	return 1000*(x+1) + 4*(y+1) + facingScore(dx, dy)
}

func part2(scanner *bufio.Scanner) interface{} {
	return 0
}
