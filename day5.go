package main

import (
	"bufio"
	"fmt"
	"os"
)

func swap(x, y int) (int, int) {
	return y, x
}

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
		if x1 != x2 && y1 != y2 {
			continue
		}

		if x1 > x2 {
			x1, x2 = swap(x1, x2)
		}

		if y1 > y2 {
			y1, y2 = swap(y1, y2)
		}

		for y := y1; y <= y2; y++ {
			for x := x1; x <= x2; x++ {
				arr[y][x]++
				if arr[y][x] == 2 {
					answer++
				}
			}
		}
	}
	fmt.Println(answer)
}
