package main

import (
	"testing"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func TestPart1(t *testing.T) {
	utils.Check(t, 24, utils.Test(part1))
}

func TestPart2(t *testing.T) {
	utils.Check(t, 93, utils.Test(part2))
}
