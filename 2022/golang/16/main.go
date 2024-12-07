package main

import (
	"bufio"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Run(part1, part2)
}

type Room struct {
	flow      int
	neighbors []string
}

func parseInput(scanner *bufio.Scanner) map[string]Room {
	r := regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnel[s]? lead[s]? to valve[s]? (.+)`)
	rooms := make(map[string]Room)
	for scanner.Scan() {
		match := r.FindStringSubmatch(scanner.Text())
		valve := match[1]
		flow, _ := strconv.Atoi(match[2])
		neighbors := strings.Split(match[3], ", ")
		rooms[valve] = Room{flow, neighbors}
	}
	return rooms
}

type Node struct {
	room  string
	depth int
}

func computeDistance(start string, rooms map[string]Room) map[string]int {
	distance := make(map[string]int)
	queue := []Node{{start, 0}}
	visited := map[string]bool{start: true}

	for len(queue) > 0 {
		for size := len(queue); size > 0; size-- {
			node := queue[0]
			queue = queue[1:]
			distance[node.room] = node.depth

			for _, neighbor := range rooms[node.room].neighbors {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, Node{neighbor, node.depth + 1})
				}
			}
		}
	}
	return distance
}

type State struct {
	time     int
	pressure int
}

func computeDP(rooms map[string]Room, timeLimit int) []map[int]State {
	distance := make(map[string]map[string]int)
	roomIds := []string{}
	for room, data := range rooms {
		if data.flow == 0 && room != "AA" {
			continue
		}
		if room != "AA" {
			roomIds = append(roomIds, room)
		}
		distance[room] = computeDistance(room, rooms)
	}
	sort.Strings(roomIds)

	m := len(roomIds)
	n := 1 << m
	dp := make([]map[int]State, n)
	dp[0] = make(map[int]State)
	for i, r := range roomIds {
		time := distance["AA"][r] + 1
		dp[0][i] = State{time: time, pressure: rooms[r].flow * (timeLimit - time)}
	}
	for mask := 1; mask < n; mask++ {
		dp[mask] = make(map[int]State)
		for i := 0; i < m; i++ {
			d := 1 << i
			if mask&d != d {
				continue
			}
			prev := mask ^ d
			if prev == 0 {
				dp[mask][i] = dp[prev][i]
				continue
			}

			maxPressure := 0
			time := 0
			for j := 0; j < m; j++ {
				d := 1 << j
				if prev&d != d {
					continue
				}
				state := dp[prev][j]
				t := state.time + (distance[roomIds[j]][roomIds[i]] + 1)
				pressure := state.pressure
				if t <= 30 {
					pressure += rooms[roomIds[i]].flow * (timeLimit - t)
				}
				if pressure > maxPressure {
					maxPressure = pressure
					time = t
				}
			}
			dp[mask][i] = State{time: time, pressure: maxPressure}
		}
	}
	return dp
}

func part1(scanner *bufio.Scanner) interface{} {
	rooms := parseInput(scanner)
	dp := computeDP(rooms, 30)
	maxPressure := 0
	for _, state := range dp[len(dp)-1] {
		if state.pressure > maxPressure {
			maxPressure = state.pressure
		}
	}
	return maxPressure
}

func part2(scanner *bufio.Scanner) interface{} {
	rooms := parseInput(scanner)
	dp := computeDP(rooms, 26)
	n := len(dp)
	maxPressure := 0
	for i := 0; i < n; i++ {
		me := i
		elephant := (n - 1) ^ me
		p1, p2 := 0, 0

		if me != 0 {
			for _, state := range dp[me] {
				if state.pressure > p1 {
					p1 = state.pressure
				}
			}
		}
		if elephant != 0 {
			for _, state := range dp[elephant] {
				if state.pressure > p2 {
					p2 = state.pressure
				}
			}
		}
		pressure := p1 + p2
		if pressure > maxPressure {
			maxPressure = pressure
		}
	}
	return maxPressure
}
