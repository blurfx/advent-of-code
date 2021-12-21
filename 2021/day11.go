package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type point struct {
	y, x int
}

type cell struct {
	energy  int
	flashed bool
}

type grid [][]*cell

var delta = [3]int{-1, 0, 1}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	arr := make(grid, 0)

	for scanner.Scan() {
		line := scanner.Text()
		cells := make([]*cell, len(line))

		for i, char := range line {
			num, _ := strconv.Atoi(string(char))
			cells[i] = &cell{num, false}
		}
		arr = append(arr, cells)
	}

	answer := 0
	for i := 0; i < 100; i++ {
		answer += arr.countFlash()
	}
	fmt.Println(answer)
}

func (arr grid) countFlash() int {
	var flashedPoints []point
	for i, row := range arr {
		for j, column := range row {
			column.energy++
			if column.energy > 9 {
				column.flashed = true
				flashedPoints = append(flashedPoints, point{i, j})
			}
		}
	}

	for len(flashedPoints) != 0 {
		flashedPoints = arr.spread(flashedPoints)
	}
	count := 0
	for _, row := range arr {
		for _, column := range row {
			if column.energy > 9 {
				column.energy = 0
				column.flashed = false
				count++
			}
		}
	}
	return count
}

func (arr grid) isSafePosition(y, x int) bool {
	return y >= 0 && y < len(arr) && x >= 0 && x < len(arr[0])
}

func (arr grid) spread(points []point) []point {
	var flashedPoints []point
	for _, p := range points {
		for _, dy := range delta {
			for _, dx := range delta {
				if dy == 0 && dx == 0 {
					continue
				}
				if !arr.isSafePosition(p.y+dy, p.x+dx) {
					continue
				}
				cell := arr[p.y+dy][p.x+dx]
				cell.energy++
				if cell.energy > 9 && !cell.flashed {
					cell.flashed = true
					flashedPoints = append(flashedPoints, point{p.y + dy, p.x + dx})
				}
			}
		}
	}
	return flashedPoints
}
