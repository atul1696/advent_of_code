package main

import (
	"bufio"
	"strconv"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

func parseInput(scanner *bufio.Scanner) *[][]int {
	grid := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0)
		for _, h := range line {
			height, _ := strconv.Atoi(string(h))
			row = append(row, height)
		}
		grid = append(grid, row)
	}
	return &grid
}

func part1(scanner *bufio.Scanner) interface{} {
	grid := parseInput(scanner)
	m, n := len(*grid), len((*grid)[0])
	visible := [][]bool{}
	for i := 0; i < m; i++ {
		row := make([]bool, n)
		if i == 0 || i == m-1 {
			for j := 0; j < n; j++ {
				row[j] = true
			}
		} else {
			row[0] = true
			row[n-1] = true
		}
		visible = append(visible, row)
	}

	height := 0
	for i := 0; i < m; i++ {
		height = (*grid)[i][0]
		for j := 1; j < n-1; j++ {
			h := (*grid)[i][j]
			if h > height {
				height = h
				visible[i][j] = true
			}
		}

		height = (*grid)[i][n-1]
		for j := n - 2; j > 0; j-- {
			h := (*grid)[i][j]
			if h > height {
				height = h
				visible[i][j] = true
			}
		}
	}

	for j := 0; j < n; j++ {
		height = (*grid)[0][j]
		for i := 1; i < m-1; i++ {
			h := (*grid)[i][j]
			if h > height {
				height = h
				visible[i][j] = true
			}
		}

		height = (*grid)[m-1][j]
		for i := m - 2; i > 0; i-- {
			h := (*grid)[i][j]
			if h > height {
				height = h
				visible[i][j] = true
			}
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visible[i][j] {
				count += 1
			}
		}
	}
	return count
}

func part2(scanner *bufio.Scanner) interface{} {
	maxScore := 0
	grid := parseInput(scanner)
	m, n := len(*grid), len((*grid)[0])

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			var count int
			var k int
			score := 1
			h := (*grid)[i][j]

			// left
			count = 0
			for k = j - 1; k >= 0 && (*grid)[i][k] < h; k-- {
				count += 1
			}
			if k >= 0 {
				count += 1
			}
			score *= count

			// right
			count = 0
			for k = j + 1; k < n && (*grid)[i][k] < h; k++ {
				count += 1
			}
			if k < n {
				count += 1
			}
			score *= count

			// up
			count = 0
			for k = i - 1; k >= 0 && (*grid)[k][j] < h; k-- {
				count += 1
			}
			if k >= 0 {
				count += 1
			}
			score *= count

			// down
			count = 0
			for k = i + 1; k < m && (*grid)[k][j] < h; k++ {
				count += 1
			}
			if k < m {
				count += 1
			}
			score *= count

			if score > maxScore {
				maxScore = score
			}
		}

	}
	return maxScore
}
