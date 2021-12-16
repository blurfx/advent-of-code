package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type filterFunc = func(string) bool

func filter(array []string, filterCallback filterFunc) []string {
	filteredSlice := make([]string, 0)

	for _, value := range array {
		if filterCallback(value) {
			filteredSlice = append(filteredSlice, value)
		}
	}

	return filteredSlice
}

func boolToString(flag bool) string {
	if flag {
		return "1"
	}
	return "0"
}

func filterRating(slice []string, prefixBuilder *strings.Builder, index int, flip bool) []string {
	if len(slice) > 1 {
		var count [2]int
		for _, item := range slice {
			if item[index] == '1' {
				count[1] += 1
			} else {
				count[0] += 1
			}
		}

		if count[1] >= count[0] {
			prefixBuilder.WriteString(boolToString(flip))
		} else {
			prefixBuilder.WriteString(boolToString(!flip))
		}

		slice = filter(slice, func(value string) bool {
			return strings.HasPrefix(value, prefixBuilder.String())
		})
	}
	return slice
}

func main() {
	const length = 12
	var oxygen []string
	var co2 []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		oxygen = append(oxygen, line)
		co2 = append(co2, line)
	}

	var oxygenPrefix strings.Builder
	var co2Prefix strings.Builder
	for i := 0; i < length; i++ {
		oxygen = filterRating(oxygen, &oxygenPrefix, i, false)
		co2 = filterRating(co2, &co2Prefix, i, true)
	}

	oxygenRating, _ := strconv.ParseInt(oxygen[0], 2, 0)
	co2Rating, _ := strconv.ParseInt(co2[0], 2, 0)

	fmt.Println(oxygenRating * co2Rating)
}
