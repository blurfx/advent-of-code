package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const length = 12
	var one [length]int
	var zero [length]int
	scanner := bufio.NewScanner(os.Stdin)
	gamma, epsilon := 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		for i, bit := range line {
			if bit == 48 {
				zero[i] += 1
			} else {
				one[i] += 1
			}
		}
	}
	for i := 0; i < length; i++ {
		if one[length-i-1] > zero[length-i-1] {
			gamma += 1 << i
		} else {
			epsilon += 1 << i
		}
	}

	fmt.Println(gamma * epsilon)
}
