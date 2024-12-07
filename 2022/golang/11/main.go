package main

import (
	"bufio"
	"regexp"
	"sort"
	"strconv"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

type Operation int

const (
	ADD Operation = iota
	MUL
	SQR
)

type Monkey struct {
	items                        []int64
	op                           Operation
	operand, divisor, pass, fail int64
}
type MonkeyParser struct {
	number, operation *regexp.Regexp
}

func atoi(s string) int64 {
	num, _ := strconv.Atoi(s)
	return int64(num)
}

func (p MonkeyParser) parse(lines []string) *Monkey {
	var match string
	// items
	items := make([]int64, 0)
	for _, match := range p.number.FindAllString(lines[0], -1) {
		items = append(items, atoi(match))
	}
	match = p.operation.FindString(lines[1])
	var op Operation
	switch match {
	case "+":
		op = ADD
	case "*":
		op = MUL
	}
	match = p.number.FindString(lines[1])
	var operand, div, pass, fail int64
	if len(match) != 0 {
		operand = atoi(match)
	} else {
		op = SQR
	}
	div = atoi(p.number.FindString(lines[2]))
	pass = atoi(p.number.FindString(lines[3]))
	fail = atoi(p.number.FindString(lines[4]))
	return &Monkey{items, op, operand, div, pass, fail}
}

func simulate(scanner *bufio.Scanner, rounds int, relief bool) int64 {
	monkeys := []*Monkey{}
	parser := MonkeyParser{
		number:    regexp.MustCompile(`(\d+)`),
		operation: regexp.MustCompile(`([\+\*])`),
	}
	lines := make([]string, 6)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		lines[i] = line
		i += 1
		if i == 6 {
			i = 0
			monkeys = append(monkeys, parser.parse(lines[1:]))
		}
	}
	var mod int64 = 1
	for _, monkey := range monkeys {
		mod *= monkey.divisor
	}
	inspections := make([]int64, len(monkeys))
	for r := 0; r < rounds; r++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				inspections[i] += 1
				new := item
				switch monkey.op {
				case ADD:
					new += monkey.operand
				case MUL:
					new *= monkey.operand
				case SQR:
					new *= new
				}
				if relief {
					new /= 3
				} else {
					new %= mod
				}
				index := monkey.fail
				if new%monkey.divisor == 0 {
					index = monkey.pass
				}
				(*monkeys[index]).items = append((*monkeys[index]).items, new)
			}
			monkey.items = monkey.items[:0]
		}
	}
	sort.SliceStable(inspections, func(i, j int) bool {
		return inspections[i] >= inspections[j]
	})
	var business int64 = 1
	for _, count := range inspections[:2] {
		business *= count
	}
	return business
}

func part1(scanner *bufio.Scanner) interface{} {
	return simulate(scanner, 20, true)
}

func part2(scanner *bufio.Scanner) interface{} {
	return simulate(scanner, 10000, false)
}
