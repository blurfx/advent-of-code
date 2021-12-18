package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	positionsStr := strings.Split(line, ",")
	positions := make([]int, len(positionsStr))
	for i, value := range positionsStr {
		num, err := strconv.ParseInt(value, 10, 0)
		if err != nil {
			panic(any(err))
		}
		positions[i] = int(num)
	}
	sort.Ints(positions)
	target := positions[len(positions)/2]
	cost := 0
	for _, value := range positions {
		cost += int(math.Abs(float64(target - value)))
	}
	fmt.Println(cost)
}
