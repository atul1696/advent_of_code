package main

import (
	"testing"

	"github.com/atul1696/advent_of_code/utils"
)

func TestPart1(t *testing.T) {
	output := utils.Test(part1)
	utils.Check(t, 15, output)
}

func TestPart2(t *testing.T) {
	output := utils.Test(part2)
	utils.Check(t, 12, output)
}
