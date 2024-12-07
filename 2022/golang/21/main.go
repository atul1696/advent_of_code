package main

import (
	"bufio"
	"regexp"
	"strconv"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

type Node struct {
	val         int
	op          string
	left, right string
	leaf        bool
}

func parseInput(scanner *bufio.Scanner) map[string]Node {
	nodes := make(map[string]Node)
	leafRgx := regexp.MustCompile(`(\w+): (\d+)`)
	nonLeafRgx := regexp.MustCompile(`(\w+): (\w+) (.) (\w+)`)
	for scanner.Scan() {
		line := scanner.Text()
		if match := leafRgx.FindStringSubmatch(line); match != nil {
			val, _ := strconv.Atoi(match[2])
			nodes[match[1]] = Node{val: val, leaf: true}
		} else {
			match = nonLeafRgx.FindStringSubmatch(line)
			nodes[match[1]] = Node{left: match[2], right: match[4], op: match[3]}
		}
	}
	return nodes
}

func compute(nodes map[string]Node, id string) int {
	node := nodes[id]
	if node.leaf {
		return node.val
	} else {
		leftOperand := compute(nodes, node.left)
		rightOperand := compute(nodes, node.right)
		result := 0
		switch node.op {
		case "+":
			result = leftOperand + rightOperand
		case "-":
			result = leftOperand - rightOperand
		case "*":
			result = leftOperand * rightOperand
		case "/":
			result = leftOperand / rightOperand
		}
		return result
	}
}

func part1(scanner *bufio.Scanner) interface{} {
	nodes := parseInput(scanner)
	return compute(nodes, "root")
}

func calculate(nodes map[string]Node, id string, results map[string]int) bool {
	node := nodes[id]
	if node.leaf {
		if id == "humn" {
			return false
		}
		results[id] = node.val
		return true
	} else if id == "root" {
		calculate(nodes, node.left, results)
		calculate(nodes, node.right, results)
		return false
	} else {
		leftCalculated := calculate(nodes, node.left, results)
		rightCalculated := calculate(nodes, node.right, results)
		if leftCalculated && rightCalculated {
			result := 0
			leftVal := results[node.left]
			rightVal := results[node.right]
			switch node.op {
			case "+":
				result = leftVal + rightVal
			case "-":
				result = leftVal - rightVal
			case "*":
				result = leftVal * rightVal
			case "/":
				result = leftVal / rightVal
			}
			results[id] = result
			return true
		}
		return false
	}
}

func part2(scanner *bufio.Scanner) interface{} {
	nodes := parseInput(scanner)
	root := nodes["root"]
	nodes["root"] = Node{op: "=", left: root.left, right: root.right}
	results := make(map[string]int)
	calculate(nodes, "root", results)
	id := "root"
	for id != "humn" {
		node := nodes[id]
		if _, ok := results[node.left]; !ok {
			val := 0
			switch node.op {
			case "+":
				val = results[id] - results[node.right]
			case "-":
				val = results[id] + results[node.right]
			case "*":
				val = results[id] / results[node.right]
			case "/":
				val = results[id] * results[node.right]
			case "=":
				val = results[node.right]
			}
			results[node.left] = val
			id = node.left
		} else {
			val := 0
			switch node.op {
			case "+":
				val = results[id] - results[node.left]
			case "-":
				val = results[node.left] - results[id]
			case "*":
				val = results[id] / results[node.left]
			case "/":
				val = results[node.left] / results[id]
			case "=":
				val = results[node.left]
			}
			results[node.right] = val
			id = node.right
		}
	}
	return results["humn"]
}
