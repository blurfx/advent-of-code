package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInput() []string {
	bytes, _ := os.ReadFile("day3/input")
	return strings.Split(string(bytes), "\n")
}

func hasNeighborSymbol(lines []string, y, x int) bool {
	adjPoints := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for _, adj := range adjPoints {
		dx, dy := x+adj[0], y+adj[1]
		if dy < 0 || dy >= len(lines) || dx < 0 || dx >= len(lines[dy]) {
			continue
		}
		if lines[dy][dx] != '.' && (lines[dy][dx] < '0' || lines[dy][dx] > '9') {
			return true
		}
	}
	return false
}

func solvePart1(lines []string) int {
	answer := 0
	regex, _ := regexp.Compile(`\d+`)

	for i := range lines {
		matches := regex.FindAllStringIndex(lines[i], -1)
		for _, match := range matches {
			for j := match[0]; j < match[1]; j++ {
				if hasNeighborSymbol(lines, i, j) {
					n, _ := strconv.Atoi(lines[i][match[0]:match[1]])
					answer += n
					break
				}
			}
		}
	}
	return answer
}

func findNumber(line string, x int) (int, bool) {
	regex, _ := regexp.Compile(`\d+`)
	matches := regex.FindAllStringIndex(line, -1)
	for _, match := range matches {
		if x >= match[0] && x < match[1] {
			n, _ := strconv.Atoi(line[match[0]:match[1]])
			return n, true
		}
	}
	return 0, false
}

func solvePart2(lines []string) int {
	answer := 0
	adjPoints := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] != '*' {
				continue
			}
			nums := make([]int, 0)
			for _, adj := range adjPoints {
				dx, dy := j+adj[0], i+adj[1]
				if dy < 0 || dy >= len(lines) || dx < 0 || dx >= len(lines[dy]) {
					continue
				}
				if lines[dy][dx] >= '0' && lines[dy][dx] <= '9' {
					num, ok := findNumber(lines[dy], dx)
					if ok {
						duplicated := false
						for _, n := range nums {
							if n == num {
								duplicated = true
								break
							}
						}
						if !duplicated {
							nums = append(nums, num)
						}
					}
				}
			}
			if len(nums) == 2 {
				answer += nums[0] * nums[1]
			}
		}
	}
	return answer
}

func main() {
	lines := readInput()
	fmt.Println(solvePart1(lines))
	fmt.Println(solvePart2(lines))
}
