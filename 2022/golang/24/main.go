package main

import (
	"bufio"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

type State struct {
	pos  [2]int
	time int
}

type Blizzard struct {
	m, n                  int
	up, down, left, right [][]int
}

func NewBlizzard(grid []string) Blizzard {
	m := len(grid)
	n := len(grid[0])
	blizzard := Blizzard{m, n, make([][]int, 0), make([][]int, 0), make([][]int, 0), make([][]int, 0)}
	for i, row := range grid {
		blizzard.left = append(blizzard.left, []int{})
		blizzard.right = append(blizzard.right, []int{})
		for j, c := range row {
			d := string(c)
			if d == "<" {
				blizzard.left[i] = append(blizzard.left[i], j)
			} else if d == ">" {
				blizzard.right[i] = append(blizzard.right[i], j)
			}
		}
	}
	for j := 0; j < n; j++ {
		blizzard.up = append(blizzard.up, []int{})
		blizzard.down = append(blizzard.down, []int{})
		for i := 0; i < m; i++ {
			d := string(grid[i][j])
			if d == "^" {
				blizzard.up[j] = append(blizzard.up[j], i)
			} else if d == "v" {
				blizzard.down[j] = append(blizzard.down[j], i)
			}
		}
	}
	return blizzard
}

func (b Blizzard) isSafe(x, y, time int) bool {
	for _, j := range b.left[x] {
		p := j - (time % (b.n - 2))
		if p < 1 {
			p = b.n - 2 + p
		}
		if p == y {
			return false
		}
	}
	for _, j := range b.right[x] {
		p := j + (time % (b.n - 2))
		if p >= b.n-1 {
			p = 1 + p - b.n + 1
		}
		if p == y {
			return false
		}
	}
	for _, i := range b.up[y] {
		p := i - (time % (b.m - 2))
		if p < 1 {
			p = b.m - 2 + p
		}
		if p == x {
			return false
		}
	}
	for _, i := range b.down[y] {
		p := i + (time % (b.m - 2))
		if p >= b.m-1 {
			p = 1 + p - b.m + 1
		}
		if p == x {
			return false
		}

	}
	return true
}

func computeShortestTime(start, finish [2]int, initialTime int, grid []string, blizzard Blizzard) int {
	m, n := len(grid), len(grid[0])
	delta := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {0, 0}}
	queue := []State{{pos: start, time: initialTime}}
	for len(queue) > 0 {
		size := len(queue)
		visited := make(map[[2]int]bool)
		for ; size > 0; size-- {
			state := queue[0]
			if state.pos[0] == finish[0] && state.pos[1] == finish[1] {
				return state.time
			}
			queue = queue[1:]
			for _, del := range delta {
				x := state.pos[0] + del[0]
				y := state.pos[1] + del[1]
				if 0 <= x && x < m && 0 <= y && y < n {
					if string(grid[x][y]) == "#" {
						continue
					}
					if blizzard.isSafe(x, y, state.time+1) {
						if !visited[[2]int{x, y}] {
							queue = append(queue, State{[2]int{x, y}, state.time + 1})
							visited[[2]int{x, y}] = true
						}
					}
				}
			}
		}
	}
	return -1
}

func part1(scanner *bufio.Scanner) interface{} {
	grid := []string{}
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	blizzard := NewBlizzard(grid)
	m, n := len(grid), len(grid[0])
	return computeShortestTime([2]int{0, 1}, [2]int{m - 1, n - 2}, 0, grid, blizzard)
}

func part2(scanner *bufio.Scanner) interface{} {
	grid := []string{}
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	blizzard := NewBlizzard(grid)
	m, n := len(grid), len(grid[0])
	start, finish := [2]int{0, 1}, [2]int{m - 1, n - 2}
	time := computeShortestTime(start, finish, 0, grid, blizzard)
	time = computeShortestTime(finish, start, time, grid, blizzard)
	time = computeShortestTime(start, finish, time, grid, blizzard)
	return time
}
