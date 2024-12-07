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

func getRanges(str string) ([]int, []int) {
	sections := strings.Split(str, ",")
	r1 := make([]int, 2)
	for i, v := range strings.Split(sections[0], "-") {
		if d, err := strconv.Atoi(v); err == nil {
			r1[i] = d
		}
	}
	r2 := make([]int, 2)
	for i, v := range strings.Split(sections[1], "-") {
		if d, err := strconv.Atoi(v); err == nil {
			r2[i] = d
		}
	}
	return r1, r2
}

func part1(scanner *bufio.Scanner) interface{} {
	overlaps := 0
	for scanner.Scan() {
		r1, r2 := getRanges(scanner.Text())
		if (r1[0] >= r2[0] && r1[1] <= r2[1]) || (r2[0] >= r1[0] && r2[1] <= r1[1]) {
			overlaps += 1
		}
	}

	return overlaps
}

func part2(scanner *bufio.Scanner) interface{} {
	overlaps := 0
	for scanner.Scan() {
		r1, r2 := getRanges(scanner.Text())
		if !(r1[0] > r2[1] || r2[0] > r1[1]) {
			overlaps += 1
		}
	}

	return overlaps
}
