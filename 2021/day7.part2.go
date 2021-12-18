package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	positionsStr := strings.Split(line, ",")
	positions := make([]int, len(positionsStr))
	min := math.MaxInt
	average := 0
	for i, value := range positionsStr {
		num, err := strconv.ParseInt(value, 10, 0)
		if err != nil {
			panic(any(err))
		}
		positions[i] = int(num)
		if positions[i] < min {
			min = positions[i]
		}
		average += positions[i]
	}
	average = (average / len(positions)) + 1

	cost := math.MaxInt
	for target := min; target <= average; target++ {
		localCost := 0
		for _, value := range positions {
			diff := int(math.Abs(float64(target - value)))
			localCost += (diff*diff + diff) / 2
		}
		if cost > localCost {
			cost = localCost
		}
	}
	fmt.Println(cost)
}
