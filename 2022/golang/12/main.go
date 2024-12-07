package main

import (
	"bufio"
	"strings"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

type entry struct {
	x, y, dist int
}

func computeShortestPath(grid *[]string, start *[][]int) int {
	delta := [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	m, n := len(*grid), len((*grid)[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	queue := make([]entry, len(*start))
	for i, s := range *start {
		queue[i] = entry{s[0], s[1], 0}
		visited[s[0]][s[1]] = true
	}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			e := queue[i]
			val := string((*grid)[e.x][e.y])
			if val == "S" {
				val = "a"
			} else if val == "E" {
				return e.dist
			}
			for _, d := range delta {
				x := e.x + d[0]
				y := e.y + d[1]
				if 0 <= x && x < m && 0 <= y && y < n {
					val2 := rune((*grid)[x][y])
					if string((*grid)[x][y]) == "E" {
						val2 = []rune("z")[0]
					}
					if val2-[]rune(val)[0] <= 1 && !visited[x][y] {
						visited[x][y] = true
						queue = append(queue, entry{x, y, e.dist + 1})
					}
				}
			}
		}
		queue = queue[size:]
	}
	return 0
}

func part1(scanner *bufio.Scanner) interface{} {
	grid := make([]string, 0)
	start := [][]int{}
	for scanner.Scan() {
		row := scanner.Text()
		grid = append(grid, row)
		if index := strings.Index(row, "S"); index != -1 {
			start = append(start, []int{len(grid) - 1, index})
		}
	}
	return computeShortestPath(&grid, &start)
}

func part2(scanner *bufio.Scanner) interface{} {
	grid := make([]string, 0)
	start := [][]int{}
	for scanner.Scan() {
		row := scanner.Text()
		grid = append(grid, row)
	}
	for i, row := range grid {
		for j, col := range row {
			val := string(col)
			if val == "a" || val == "S" {
				start = append(start, []int{i, j})
			}
		}
	}
	return computeShortestPath(&grid, &start)
}
