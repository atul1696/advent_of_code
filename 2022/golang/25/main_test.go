package main

import (
	"testing"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func TestPart1(t *testing.T) {
	utils.Check(t, "2=-1=0", utils.Test(part1))
}

func TestDecimalToSnafu(t *testing.T) {
	tests := []struct {
		num      int
		expected string
	}{
		{2022, "1=11-2"},
		{12345, "1-0---0"},
		{314159265, "1121-1110-1=0"},
		{1747, "1=-0-2"},
		{906, "12111"},
		{198, "2=0="},
		{11, "21"},
		{201, "2=01"},
		{31, "111"},
		{1257, "20012"},
		{32, "112"},
		{353, "1=-1="},
		{107, "1-12"},
		{7, "12"},
		{3, "1="},
		{37, "122"},
	}
	for _, test := range tests {
		if actual := snafu(test.num); actual != test.expected {
			t.Errorf("Expected %+v, got %+v", test.expected, actual)
		}
	}
}
