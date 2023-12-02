package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInput() []string {
	bytes, _ := os.ReadFile("day2/input")
	return strings.Split(string(bytes), "\n")
}

func solvePart1(lines []string) int {
	answer := 0
	const R = 12
	const G = 13
	const B = 14

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		r, g, b := 0, 0, 0
		regexR, _ := regexp.Compile(`(\d+) red`)
		regexG, _ := regexp.Compile(`(\d+) green`)
		regexB, _ := regexp.Compile(`(\d+) blue`)
		matchesR := regexR.FindAllStringSubmatch(line, -1)
		matchesG := regexG.FindAllStringSubmatch(line, -1)
		matchesB := regexB.FindAllStringSubmatch(line, -1)

		for _, match := range matchesR {
			v, _ := strconv.Atoi(match[1])
			if v > r {
				r = v
			}
		}

		for _, match := range matchesG {
			v, _ := strconv.Atoi(match[1])
			if v > g {
				g = v
			}
		}

		for _, match := range matchesB {
			v, _ := strconv.Atoi(match[1])
			if v > b {
				b = v
			}
		}

		if R >= r && G >= g && B >= b {
			answer += i + 1
		}
	}
	return answer
}

func solvePart2(lines []string) int {
	answer := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		r, g, b := 0, 0, 0
		regexR, _ := regexp.Compile(`(\d+) red`)
		regexG, _ := regexp.Compile(`(\d+) green`)
		regexB, _ := regexp.Compile(`(\d+) blue`)
		matchesR := regexR.FindAllStringSubmatch(line, -1)
		matchesG := regexG.FindAllStringSubmatch(line, -1)
		matchesB := regexB.FindAllStringSubmatch(line, -1)

		for _, match := range matchesR {
			v, _ := strconv.Atoi(match[1])
			if v > r {
				r = v
			}
		}

		for _, match := range matchesG {
			v, _ := strconv.Atoi(match[1])
			if v > g {
				g = v
			}
		}

		for _, match := range matchesB {
			v, _ := strconv.Atoi(match[1])
			if v > b {
				b = v
			}
		}

		answer += r * g * b
	}
	return answer
}

func main() {
	lines := readInput()
	fmt.Println(solvePart1(lines))
	fmt.Println(solvePart2(lines))
}
