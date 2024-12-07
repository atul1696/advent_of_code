package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/atul1696/advent_of_code/2022/golang/utils"
)

func main() {
	utils.Debug(part1, part2)
}

const (
	ORE int = iota
	CLAY
	OBSIDIAN
	GEODE
)

type Cost [4][4]int
type Robots [4]int
type Resources [4]int

func parseLine(line string) (id int, cost Cost) {
	numRgx := regexp.MustCompile(`\d+`)
	resourceRgx := regexp.MustCompile(`(\d+) (\w+)`)
	p := strings.Split(line, ":")
	m1 := numRgx.FindStringSubmatch(p[0])
	if val, err := strconv.Atoi(m1[0]); err == nil {
		id = val
	}

	for i, l := range strings.Split(p[1], ".") {
		for _, match := range resourceRgx.FindAllStringSubmatch(l, -1) {
			var resourceType int
			switch match[2] {
			case "ore":
				resourceType = ORE
			case "clay":
				resourceType = CLAY
			case "obsidian":
				resourceType = OBSIDIAN
			default:
				resourceType = GEODE
			}
			if val, err := strconv.Atoi(match[1]); err == nil {
				cost[i][resourceType] = val
			}
		}
	}
	return
}

type State struct {
	time      int
	robots    Robots
	resources Resources
}

func dfs(cost Cost, robots Robots, res Resources, time int, visited map[State]int) int {
	state := State{time: time, robots: robots, resources: res}
	fmt.Println(state)
	if val, ok := visited[state]; ok {
		return val
	}

	if time == 24 {
		visited[state] = res[GEODE]
		return res[GEODE]
	}

	var updatedRes Resources
	for i := 0; i < 4; i++ {
		updatedRes[i] = res[i] + robots[i]
	}
	val := dfs(cost, robots, updatedRes, time+1, visited)

	for i := 0; i < 4; i++ {
		canCreate := true
		for j := 0; j < 4; j++ {
			if res[j] < cost[i][j] {
				fmt.Println(i, res[j], cost[i][j])
				canCreate = false
				break
			}
		}
		if !canCreate {
			continue
		}

		var updatedRes Resources
		for j := 0; j < 4; j++ {
			updatedRes[j] = res[j] - cost[i][j] + robots[j]
		}
		robots[i] += 1
		v := dfs(cost, robots, updatedRes, time+1, visited)
		if v > val {
			val = v
		}
		robots[i] -= 1
	}
	visited[state] = val
	return val
}

func part1(scanner *bufio.Scanner) interface{} {
	quantity := 0
	for scanner.Scan() {
		id, cost := parseLine(scanner.Text())
		var robots Robots
		robots[ORE] = 1
		var resources Resources
		visited := make(map[State]int)
		g := dfs(cost, robots, resources, 0, visited)
		quantity += id * g
		fmt.Println(id, g, cost)
	}
	return quantity
}

func part2(scanner *bufio.Scanner) interface{} {
	return 0
}
