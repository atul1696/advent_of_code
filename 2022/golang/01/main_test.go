package main

import (
	"testing"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func TestPart1(t *testing.T) {
	output := utils.Test(part1)
	utils.Check(t, 24000, output)
}

func TestPart2(t *testing.T) {
	output := utils.Test(part2)
	utils.Check(t, 45000, output)
}
