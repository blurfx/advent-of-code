package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	initialValues := strings.Split(scanner.Text(), ",")
	fishes := make([]int, len(initialValues))
	for i, value := range initialValues {
		num, err := strconv.ParseInt(value, 10, 0)
		if err != nil {
			panic(any(err))
		}
		fishes[i] = int(num)
	}

	for day := 0; day < 80; day++ {
		newFishes := 0
		for i := range fishes {
			fishes[i]--
			if fishes[i] < 0 {
				fishes[i] = 6
				newFishes++
			}
		}
		for ; newFishes > 0; newFishes-- {
			fishes = append(fishes, 8)
		}
	}
	fmt.Println(len(fishes))
}
