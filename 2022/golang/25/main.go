package main

import (
	"bufio"
	"strings"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

func snafu(num int) string {
	s := make([]string, 0)
	max := 0
	pow := 1
	digits := 0

	for {
		digits += 1
		max += 2 * pow
		if num <= max {
			break
		}
		pow *= 5
	}

	curr := 0
	for ; digits > 0; digits-- {
		max -= 2 * pow
		if curr <= num {
			if num-curr <= max {
				s = append(s, "0")
			} else if num-curr-pow <= max {
				s = append(s, "1")
				curr += pow
			} else {
				s = append(s, "2")
				curr += 2 * pow
			}
		} else {
			if -max <= curr-pow-num && curr-pow-num <= max {
				s = append(s, "-")
				curr -= pow
			} else if -max <= curr-2*pow-num && curr-2*pow-num <= max {
				s = append(s, "=")
				curr -= 2 * pow
			} else {
				s = append(s, "0")
			}
		}
		pow /= 5
	}

	return strings.Join(s, "")
}

func decimal(s string) int {
	num := 0
	n := len(s)
	pow := 1
	mapping := map[string]int{"0": 0, "1": 1, "2": 2, "=": -2, "-": -1}
	for i := 0; i < n; i++ {
		num += pow * mapping[string(s[n-1-i])]
		pow *= 5
	}
	return num
}

func part1(scanner *bufio.Scanner) interface{} {
	fuel := 0
	for scanner.Scan() {
		fuel += decimal(scanner.Text())
	}
	return snafu(fuel)
}

func part2(scanner *bufio.Scanner) interface{} {
	return 0
}
