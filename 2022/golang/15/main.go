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

type Coordinate struct {
	X, Y int
}

type SensorInfo struct {
	sensor, beacon Coordinate
}

func atoi(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

func parse(scanner *bufio.Scanner) []SensorInfo {
	data := []SensorInfo{}
	r := regexp.MustCompile(`-?(\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := r.FindAllString(line, -1)
		s := SensorInfo{
			sensor: Coordinate{X: atoi(matches[0]), Y: atoi(matches[1])},
			beacon: Coordinate{X: atoi(matches[2]), Y: atoi(matches[3])},
		}
		data = append(data, s)
	}
	return data
}

func part1(scanner *bufio.Scanner) interface{} {
	Y := 2000000
	beacons := make(map[int]bool)
	sensors := make(map[int]bool)
	row := make([][]int, 0)

	for _, info := range parse(scanner) {
		dist := abs(info.sensor.X-info.beacon.X) + abs(info.sensor.Y-info.beacon.Y)
		if info.beacon.Y == Y {
			beacons[info.beacon.X] = true
		}

		if info.sensor.Y == Y {
			sensors[info.sensor.X] = true
		}

		if info.sensor.Y-dist <= Y && Y <= info.sensor.Y+dist {
			j := abs(info.sensor.Y - Y)
			row = append(row, []int{info.sensor.X - dist + j, info.sensor.X + dist - j})
		}
	}

	if len(row) == 0 {
		return 0
	}

	count := 0
	sort.SliceStable(row, func(i, j int) bool {
		return row[i][0] < row[j][0] || row[i][0] == row[j][0] && row[i][1] > row[j][1]
	})
	start := row[0][0]
	end := row[0][1]
	for _, r := range row[1:] {
		if r[0] > end {
			count += end - start + 1
			start = r[0]
			end = r[1]
		} else if r[1] > end {
			end = r[1]
		}
	}
	count += end - start + 1
	count -= len(beacons) + len(sensors)
	return count
}

func part2(scanner *bufio.Scanner) interface{} {
	beacons := make(map[int]bool)
	sensors := make(map[int]bool)
	data := parse(scanner)
	bound := 4000000
	minY, maxY := 4000000, 0
	for _, info := range data {
		dist := abs(info.sensor.X-info.beacon.X) + abs(info.sensor.Y-info.beacon.Y)
		y := info.sensor.Y - dist
		if y < minY {
			minY = y
		} else if y > maxY {
			maxY = y
		}
	}
	if minY < 0 {
		minY = 0
	}
	if maxY > bound {
		maxY = bound
	}
	for Y := minY; Y <= maxY; Y++ {
		row := make([][]int, 0)
		for _, info := range data {
			dist := abs(info.sensor.X-info.beacon.X) + abs(info.sensor.Y-info.beacon.Y)
			if info.beacon.Y == Y {
				beacons[info.beacon.X] = true
			}

			if info.sensor.Y == Y {
				sensors[info.sensor.X] = true
			}

			if info.sensor.Y-dist <= Y && Y <= info.sensor.Y+dist {
				j := abs(info.sensor.Y - Y)
				row = append(row, []int{info.sensor.X - dist + j, info.sensor.X + dist - j})
			}
		}

		if len(row) == 0 {
			continue
		}

		sort.SliceStable(row, func(i, j int) bool {
			return row[i][0] < row[j][0] || row[i][0] == row[j][0] && row[i][1] < row[j][1]
		})
		end := row[0][1]
		for _, r := range row[1:] {
			if r[0] > end {
				if 0 <= end && r[0] <= bound && r[0]-end > 1 {
					return (end+1)*bound + Y
				}
				end = r[1]
			} else if r[1] > end {
				end = r[1]
			}
		}

	}
	return -1
}
