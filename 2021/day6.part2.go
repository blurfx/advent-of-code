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
	timerPool := make([]int, 9)
	for _, value := range initialValues {
		num, err := strconv.ParseInt(value, 10, 0)
		if err != nil {
			panic(any(err))
		}
		timerPool[int(num)]++
	}

	for day := 0; day < 256; day++ {
		tick := timerPool[0]
		timerPool = append(timerPool[1:], tick)
		timerPool[6] += tick
	}
	sum := 0
	for _, fishes := range timerPool {
		sum += fishes
	}
	fmt.Println(sum)
}
