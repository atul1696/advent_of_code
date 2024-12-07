package main

import (
	"bufio"
	"strconv"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

type DLL struct {
	val        int
	index      int
	prev, next *DLL
}

func find(head DLL, x int) (int, *DLL) {
	node := head.next
	index := 0
	for node.index != x {
		index += 1
		node = node.next
	}
	return index, node
}

func compute(nums []int, rounds int) int {
	head := DLL{val: -1, index: -1}
	tail := DLL{val: -1, index: -1}

	curr := &head
	head.next = &tail
	head.prev = &tail
	tail.prev = &head
	tail.next = &head

	zeroIndex := 0
	for i, num := range nums {
		if num == 0 {
			zeroIndex = i
		}
		node := DLL{val: num, index: i, prev: curr, next: curr.next}
		curr.next.prev = &node
		curr.next = &node
		curr = &node
	}

	n := len(nums)
	for r := 0; r < rounds; r++ {
		for i := 0; i < n; i++ {
			num := nums[i]
			oldIdx, node := find(head, i)
			node.prev.next = node.next
			node.next.prev = node.prev
			newIdx := (num + oldIdx) % (n - 1)
			if num < 0 {
				newIdx = oldIdx - ((-num) % (n - 1))
				if newIdx < 0 {
					newIdx += n - 1
				}
			}
			start := &head
			for k := 0; k < newIdx; k++ {
				start = start.next
			}

			node.next = start.next
			node.prev = start
			start.next.prev = node
			start.next = node
		}
	}

	grove := 0
	index, _ := find(head, zeroIndex)
	for i := 1; i <= 3; i++ {
		start := &head
		j := (index + i*1000) % n
		for k := 0; k <= j; k++ {
			start = start.next
		}
		grove += start.val
	}

	return grove
}

func part1(scanner *bufio.Scanner) interface{} {
	nums := make([]int, 0)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return compute(nums, 1)
}

func part2(scanner *bufio.Scanner) interface{} {
	nums := make([]int, 0)
	decryptionKey := 811589153
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num*decryptionKey)
	}
	return compute(nums, 10)
}
