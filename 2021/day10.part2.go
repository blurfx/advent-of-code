package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	pair := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	point := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	scores := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		err := false
		score := 0
		stack := make([]rune, 0)
		for _, char := range line {
			switch char {
			case '(', '{', '<', '[':
				stack = append(stack, char)
			default:
				last := stack[len(stack)-1]
				if pair[last] == char {
					stack = stack[:len(stack)-1]
				} else {
					err = true
				}
			}
		}
		if err {
			continue
		}
		for i := len(stack) - 1; i >= 0; i-- {
			score = score*5 + point[stack[i]]
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
