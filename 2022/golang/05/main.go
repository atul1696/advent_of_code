package main

import (
	"bufio"
	"regexp"
	"strconv"
	"unicode"

	"github.com/atul1696/advent_of_code/2022/golang/collections"
	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

func parseInput(scanner *bufio.Scanner) (*[]collections.RuneStack, *[][]int) {
	stackData := []string{}
	moves := [][]int{}
	r_move, _ := regexp.Compile(`move (\d+) from (\d+) to (\d+)`)

	isMove := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isMove = true
			continue
		}

		if isMove {
			match := make([]int, 3)
			for i, m := range r_move.FindStringSubmatch(line)[1:] {
				d, _ := strconv.Atoi(m)
				match[i] = d
			}
			moves = append(moves, match)
		} else {
			stackData = append(stackData, line)
		}
	}

	r_count := regexp.MustCompile(`\d+`)
	line := stackData[len(stackData)-1]
	count := len(r_count.FindAll([]byte(line), -1))
	stacks := make([]collections.RuneStack, count)
	for i := 0; i < count; i++ {
		stacks[i] = collections.NewRuneStack()
	}
	for j := len(stackData) - 2; j >= 0; j-- {
		line := stackData[j]
		for i := 0; i < len(line); i += 4 {
			val := rune(line[i+1])
			if !unicode.IsSpace(val) {
				stacks[i/4].Append(val)
			}
		}
	}
	return &stacks, &moves
}

func part1(scanner *bufio.Scanner) interface{} {
	output := ""
	stacks, moves := parseInput(scanner)
	for _, move := range *moves {
		l, i, j := move[0], move[1], move[2]
		for ; l > 0; l-- {
			val, _ := (*stacks)[i-1].Pop()
			(*stacks)[j-1].Append(val)
		}
	}
	for _, s := range *stacks {
		if val, ok := s.Top(); ok {
			output += string(val)
		}
	}
	return output
}

func part2(scanner *bufio.Scanner) interface{} {
	output := ""
	stacks, moves := parseInput(scanner)
	for _, move := range *moves {
		l, i, j := move[0], move[1], move[2]
		s := collections.NewRuneStack()
		for k := 0; k < l; k++ {
			val, _ := (*stacks)[i-1].Pop()
			s.Append(val)
		}
		for k := 0; k < l; k++ {
			val, _ := s.Pop()
			(*stacks)[j-1].Append(val)
		}
	}
	for _, s := range *stacks {
		if val, ok := s.Top(); ok {
			output += string(val)
		}
	}
	return output
}
