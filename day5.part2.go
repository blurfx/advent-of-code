package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var arr [1000][1000]int
	scanner := bufio.NewScanner(os.Stdin)
	answer := 0

	for scanner.Scan() {
		var x1, y1, x2, y2 int
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			panic(any(err))
		}

		xDelta, yDelta := 1, 1
		if x1 > x2 {
			xDelta = -1
		} else if x1 == x2 {
			xDelta = 0
		}

		if y1 > y2 {
			yDelta = -1
		} else if y1 == y2 {
			yDelta = 0
		}

		x, y := x1, y1
		for {
			arr[y][x]++
			if arr[y][x] == 2 {
				answer++
			}
			if x == x2 && y == y2 {
				break
			}
			x += xDelta
			y += yDelta
		}
	}
	fmt.Println(answer)
}
