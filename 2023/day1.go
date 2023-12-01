package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInput() []string {
	bytes, _ := os.ReadFile("day1.input")
	return strings.Split(string(bytes), "\n")
}

func solvePart1(lines []string) int {
	answer := 0
	regex, _ := regexp.Compile(`\d`)
	for _, line := range lines {
		matches := regex.FindAllString(line, -1)
		if len(matches) == 0 {
			continue
		}
		first, _ := strconv.Atoi(matches[0])
		last, _ := strconv.Atoi(matches[len(matches)-1])

		answer += first*10 + last
	}

	return answer
}

func main() {
	lines := readInput()
	fmt.Println(solvePart1(lines))
}
