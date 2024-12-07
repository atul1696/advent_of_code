package main

import (
	"bufio"
	"sort"
	"strconv"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

func part1(scanner *bufio.Scanner) interface{} {
	maxCalories := 0
	calories := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if maxCalories < calories {
				maxCalories = calories
			}
			calories = 0
		} else {
			if cal, err := strconv.Atoi(line); err == nil {
				calories += cal
			}
		}
	}
	if maxCalories < calories {
		maxCalories = calories
	}
	return maxCalories
}

func part2(scanner *bufio.Scanner) interface{} {
	caloryArray := []int{}
	calories := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			caloryArray = append(caloryArray, calories)
			calories = 0
		} else if cal, err := strconv.Atoi(line); err == nil {
			calories += cal
		}
	}
	caloryArray = append(caloryArray, calories)

	sort.Ints(caloryArray)
	top3 := 0
	for i, j := len(caloryArray)-1, 0; j < 3; i, j = i-1, j+1 {
		top3 += caloryArray[i]
	}
	return top3
}
