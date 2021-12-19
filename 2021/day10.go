package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	answer := 0
	pair := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	point := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	for scanner.Scan() {
		line := scanner.Text()
		err := false
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
					answer += point[char]
					err = true
				}
			}
			if err {
				break
			}
		}
	}
	fmt.Println(answer)
}
