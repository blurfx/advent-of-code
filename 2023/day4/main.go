package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func readInput() []string {
	bytes, _ := os.ReadFile("day4/input")
	return strings.Split(string(bytes), "\n")
}

func solvePart1(lines []string) int {
	answer := 0
	regex, _ := regexp.Compile(`\d+`)
	for _, line := range lines {
		split := strings.Split(line, ": ")
		split = strings.Split(split[1], "|")
		winning := regex.FindAllString(split[0], -1)
		my := regex.FindAllString(split[1], -1)
		count := 0
		for _, x := range winning {
			for _, y := range my {
				if x == y {
					count += 1
				}
			}
		}
		if count > 0 {
			answer += int(math.Pow(float64(2), float64(count-1)))
		}
	}
	return answer
}

func solvePart2(lines []string) int {
	answer := 0
	regex, _ := regexp.Compile(`\d+`)
	duplications := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		duplications[i] = 1
	}
	for i, line := range lines {
		split := strings.Split(line, ": ")
		split = strings.Split(split[1], "|")
		winning := regex.FindAllString(split[0], -1)
		my := regex.FindAllString(split[1], -1)
		count := 0
		for _, x := range winning {
			for _, y := range my {
				if x == y {
					count += 1
				}
			}
		}
		for j := i + 1; j < len(lines) && j <= i+count; j++ {
			duplications[j] += duplications[i]
		}
		answer += duplications[i]
	}
	return answer
}

func main() {
	lines := readInput()
	fmt.Println(solvePart1(lines))
	fmt.Println(solvePart2(lines))
}
