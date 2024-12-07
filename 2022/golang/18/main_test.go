package main

import (
	"testing"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func TestPart1(t *testing.T) {
	utils.Check(t, 64, utils.Test(part1))
}

func TestPart2(t *testing.T) {
	utils.Check(t, 58, utils.Test(part2))
}
