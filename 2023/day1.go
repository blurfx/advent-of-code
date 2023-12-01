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

func solvePart2(lines []string) int {
	answer := 0
	regex, _ := regexp.Compile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)

	wordToNumber := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	parseNumber := func(value string) int {
		if number, ok := wordToNumber[value]; ok {
			return number
		} else {
			number, _ := strconv.Atoi(value)
			return number
		}
	}

	for _, line := range lines {
		idx := 0
		first, last := -1, -1
		for {
			loc := regex.FindStringIndex(line[idx:])
			if loc == nil {
				break
			}
			if first == -1 {
				first = parseNumber(line[idx+loc[0] : idx+loc[1]])
				last = first
			} else {
				last = parseNumber(line[idx+loc[0] : idx+loc[1]])
			}
			idx += loc[1]
			if loc[1]-loc[0] > 1 {
				idx -= 1
			}
		}

		if first == -1 {
			continue
		}

		answer += first*10 + last
	}

	return answer
}

func main() {
	lines := readInput()
	fmt.Println(solvePart1(lines))
	fmt.Println(solvePart2(lines))
}
