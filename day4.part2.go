package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MARK = "X"

func findAndMark(board [][]string, value string) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == value {
				board[i][j] = MARK
			}
		}
	}
}

func check(board [][]string) bool {
	for i := 0; i < len(board); i++ {
		vertical, horizontal := true, true
		for j := 0; j < len(board); j++ {
			vertical = vertical && board[j][i] == MARK
			horizontal = horizontal && board[i][j] == MARK
		}

		if vertical || horizontal {
			return true
		}
	}
	return false
}

func getUnmarkedSum(board [][]string) int {
	sum := 0
	for _, row := range board {
		for _, column := range row {
			if column != MARK {
				num, _ := strconv.ParseInt(column, 10, 0)
				sum += int(num)
			}
		}
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.Split(scanner.Text(), ",")

	bingoBoards := make([][][]string, 0)
	isDone := make([]bool, 0)
	for i := 0; scanner.Scan(); i++ {
		scanner.Text()
		bingoBoards = append(bingoBoards, make([][]string, 5))
		isDone = append(isDone, false)
		for j := 0; j < 5; j++ {
			scanner.Scan()
			bingoBoards[i][j] = strings.Fields(scanner.Text())
		}
	}

	rest := len(bingoBoards)
	for _, value := range line {
		for i, board := range bingoBoards {
			findAndMark(board, value)
			if !isDone[i] && check(board) {
				rest -= 1
				isDone[i] = true
				if rest == 0 {
					sum := getUnmarkedSum(board)
					val, _ := strconv.ParseInt(value, 10, 0)
					fmt.Println(sum * int(val))
					return
				}
			}
		}
	}
}
