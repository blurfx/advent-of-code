package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	queue := make([]int64, 0)
	prefixSum := int64(0)
	answer := 0

	for i := 0; scanner.Scan(); i++ {
		value, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		queue = append(queue, value)

		if i < 3 {
			prefixSum += value
			continue
		}

		next := prefixSum - queue[0] + value
		if next > prefixSum {
			answer += 1
		}

		prefixSum = next
		queue = queue[1:]
	}

	fmt.Println(answer)
}
