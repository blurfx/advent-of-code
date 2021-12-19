package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	arr := make([]string, 0)
	dy := [4]int{-1, 0, 1, 0}
	dx := [4]int{0, -1, 0, 1}
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	height, width := len(arr), len(arr[0])
	answer := 0
	for y, line := range arr {
		for x, char := range line {
			isLowPoint := true
			for i := 0; i < 4; i++ {
				ny, nx := y+dy[i], x+dx[i]
				if ny >= 0 && ny < height && nx >= 0 && nx < width {
					isLowPoint = isLowPoint && char < rune(arr[ny][nx])
				}
			}
			if isLowPoint {
				answer += int(char) - '0' + 1
			}
		}
	}
	fmt.Println(answer)
}
