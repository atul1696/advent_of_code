package main

import (
	"bufio"
	"log"
	"unicode"

	"github.com/atul1696/advent_of_code/2022/golang/collections"
	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

func getPriority(r rune) int {
	if unicode.IsUpper(r) {
		return int(r-'A') + 27
	} else {
		return int(r-'a') + 1
	}
}

func part1(scanner *bufio.Scanner) interface{} {
	priority := 0
	for scanner.Scan() {
		rugsack := scanner.Text()
		items := make(map[rune]bool)
		for i, item := range rugsack {
			if i < len(rugsack)/2 {
				items[item] = true
			} else if items[item] {
				priority += getPriority(item)
				break
			}
		}
	}
	return priority
}

func part2(scanner *bufio.Scanner) interface{} {
	sum := 0
	scan := scanner.Scan()
	for scan {
		rugsacks := make([]string, 3)
		for i := 0; i < 3; i++ {
			rugsacks[i] = scanner.Text()
			scan = scanner.Scan()
		}
		set := collections.StringIntersection(rugsacks...)
		if set.Size() != 1 {
			log.Fatalf("Invalid number of badges: %d, %+v\n", set.Size(), set.Array())
		} else {
			for _, val := range set.Array() {
				sum += getPriority(val)
			}
		}
	}
	return sum
}
