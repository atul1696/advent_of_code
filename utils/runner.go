package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run(part1, part2 func(*bufio.Scanner) interface{}) {
	execute(part1, part2, "input.txt")
}

func Debug(part1, part2 func(*bufio.Scanner) interface{}) {
	execute(part1, part2, "sample.txt")
}

func Test(fn func(*bufio.Scanner) interface{}) interface{} {
	return run(fn, "sample.txt")
}

func run(fn func(*bufio.Scanner) interface{}, filename string) interface{} {
	scanner := FileInputScanner(filename)
	return fn(scanner)
}

func execute(part1, part2 func(*bufio.Scanner) interface{}, filename string) {
	part := os.Args[1]
	var output interface{}
	var fn func(*bufio.Scanner) interface{}
	if strings.Contains(part, "1") {
		fn = part1
	} else {
		fn = part2
	}
	output = run(fn, filename)
	fmt.Println(output)
}
