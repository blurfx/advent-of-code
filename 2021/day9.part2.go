package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var arr []string
var visited [][]bool
var dy = [4]int{-1, 0, 1, 0}
var dx = [4]int{0, -1, 0, 1}
var height, width int

func isValidPosition(y, x int) bool {
	return y >= 0 && y < height && x >= 0 && x < width
}

func dfs(y, x int) int {
	if visited[y][x] {
		return 0
	}
	visited[y][x] = true

	size := 1
	for i := 0; i < 4; i++ {
		ny, nx := y+dy[i], x+dx[i]
		if isValidPosition(ny, nx) && arr[ny][nx] != '9' {
			size += dfs(ny, nx)
		}
	}
	return size
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	height, width = len(arr), len(arr[0])
	visited = make([][]bool, height)
	for i := range arr {
		visited[i] = make([]bool, width)
	}

	sizes := make([]int, 0)
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] == '9' || visited[i][j] {
				continue
			}
			sizes = append(sizes, dfs(i, j))
		}
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})
	fmt.Println(sizes)
	fmt.Println(sizes[0] * sizes[1] * sizes[2])
}
