package main

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

type fileNode struct {
	name        string
	isDirectory bool
	size        int
	children    *[](*fileNode)
}

func NewDirectoryNode(name string) *fileNode {
	return &fileNode{name: name, isDirectory: true, children: &([]*fileNode{})}
}

func NewFileNode(name string, size int) *fileNode {
	return &fileNode{name: name, isDirectory: false, size: size}
}

func constructFileSystem(scanner *bufio.Scanner) *fileNode {
	stack := []*fileNode{}
	var root *fileNode
	for scanner.Scan() {
		line := scanner.Text()
		contents := strings.Split(line, " ")
		if contents[0] == "$" {
			if contents[1] == "cd" {
				if contents[2] == ".." {
					stack = stack[:len(stack)-1]
				} else {
					d := NewDirectoryNode(contents[2])
					if contents[2] == "/" {
						root = d
					}
					if len(stack) > 0 {
						parent := *stack[len(stack)-1]
						*parent.children = append(*parent.children, d)
					}
					stack = append(stack, d)
				}
			}
		} else {
			if contents[0] != "dir" {
				size, _ := strconv.Atoi(contents[0])
				f := NewFileNode(contents[1], size)
				parent := *stack[len(stack)-1]
				*parent.children = append(*parent.children, f)
			}
		}
	}
	return root
}

func computeDirSize(root *fileNode, dirSize *map[*fileNode]int) int {
	size := 0
	for _, child := range *root.children {
		if !child.isDirectory {
			size += child.size
		} else {
			size += computeDirSize(child, dirSize)
		}
	}
	(*dirSize)[root] = size
	return size
}

func part1(scanner *bufio.Scanner) interface{} {
	root := constructFileSystem(scanner)
	dirSize := make(map[*fileNode]int)
	computeDirSize(root, &dirSize)
	output := 0
	for _, size := range dirSize {
		if size <= 100000 {
			output += size
		}
	}
	return output
}

func part2(scanner *bufio.Scanner) interface{} {
	root := constructFileSystem(scanner)
	dirSize := make(map[*fileNode]int)

	used := computeDirSize(root, &dirSize)
	available := 70000000 - 30000000

	output := 70000000
	for _, size := range dirSize {
		if used-size <= available {
			if size < output {
				output = size
			}
		}
	}
	return output
}
