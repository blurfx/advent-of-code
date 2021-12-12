package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	prev := int64(0)
	initialized := false
	answer := 0

	for scanner.Scan() {
		value, _ := strconv.ParseInt(scanner.Text(), 10, 0)

		if initialized && value > prev {
			answer += 1
		}

		prev = value
		initialized = true
	}

	fmt.Println(answer)
}
