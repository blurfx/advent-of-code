package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	answer := 0
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " | ")
		inputs := strings.Fields(split[1])

		for _, input := range inputs {
			if len(input) == 2 || len(input) == 3 || len(input) == 4 || len(input) == 7 {
				answer++
			}
		}
	}
	fmt.Println(answer)
}
