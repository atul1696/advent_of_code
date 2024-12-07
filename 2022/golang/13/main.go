package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

type nestedIntArray struct {
	val   int
	array *[]*nestedIntArray
	isInt bool
}

func parseLine(line string) *nestedIntArray {
	var arr *nestedIntArray

	stack := []*nestedIntArray{}
	start := -1
	for i, c := range line {
		chr := string(c)
		if chr == "[" {
			a := nestedIntArray{}
			a.array = &([]*nestedIntArray{})
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				*parent.array = append(*parent.array, &a)
			} else {
				arr = &a
			}
			stack = append(stack, &a)
		} else if chr == "]" || chr == "," {
			if start != -1 {
				val, _ := strconv.Atoi(line[start:i])
				start = -1
				child := nestedIntArray{val: val, isInt: true}
				a := *stack[len(stack)-1]
				*a.array = append(*a.array, &child)
			}
			if chr == "]" {
				stack = stack[:len(stack)-1]
			}
		} else if chr != "," {
			if start == -1 {
				start = i
			}
		}
	}
	return arr
}

func printArr(a *nestedIntArray) {
	if !a.isInt {
		fmt.Print("[")
		for _, a := range *a.array {
			printArr(a)
		}
		fmt.Print("],")
	} else {
		fmt.Printf("%d,", a.val)
	}
}

func compareInt(i1, i2 int) int {
	if i1 < i2 {
		return -1
	} else if i1 == i2 {
		return 0
	} else {
		return 1
	}
}

func compare(arr1 *nestedIntArray, arr2 *nestedIntArray) int {
	if arr1.isInt && arr2.isInt {
		return compareInt(arr1.val, arr2.val)
	} else if !arr1.isInt && !arr2.isInt {
		l1 := len(*arr1.array)
		l2 := len(*arr2.array)

		var length int
		if l1 >= l2 {
			length = l2
		} else {
			length = l1
		}

		res := 0
		for i := 0; i < length; i++ {
			c := compare((*arr1.array)[i], (*arr2.array)[i])
			if c != 0 {
				res = c
				break
			}
		}
		if res == 0 {
			res = compareInt(l1, l2)
		}
		return res
	} else {
		if arr1.isInt {
			a := []*nestedIntArray{arr1}
			arr1 = &nestedIntArray{isInt: false, array: &a}
		} else {
			a := []*nestedIntArray{arr2}
			arr2 = &nestedIntArray{isInt: false, array: &a}
		}
		return compare(arr1, arr2)
	}
}

func part1(scanner *bufio.Scanner) interface{} {
	pairId := 0
	sum := 0
	for scanner.Scan() {
		pairId += 1
		arr1 := parseLine(scanner.Text())

		scanner.Scan()
		arr2 := parseLine(scanner.Text())

		if compare(arr1, arr2) < 0 {
			sum += pairId
		}
		scanner.Scan()
	}
	return sum
}

func part2(scanner *bufio.Scanner) interface{} {
	nestedArrays := []*nestedIntArray{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		arr1 := parseLine(line)
		nestedArrays = append(nestedArrays, arr1)
	}
	div2 := parseLine("[[2]]")
	div6 := parseLine("[[6]]")
	nestedArrays = append(nestedArrays, div2, div6)

	sort.SliceStable(nestedArrays, func(i, j int) bool {
		return compare(nestedArrays[i], nestedArrays[j]) <= 0
	})
	key := 1
	for i, arr := range nestedArrays {
		if arr == div2 || arr == div6 {
			key *= (i + 1)
		}
	}
	return key
}
