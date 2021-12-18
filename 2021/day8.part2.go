package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	answer := 0
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " | ")
		signals := strings.Fields(split[0])
		for i := range signals {
			signals[i] = SortString(signals[i])
		}
		inputs := strings.Fields(split[1])
		for i := range inputs {
			inputs[i] = SortString(inputs[i])
		}
		segments := make([]string, 10)
		knownPair := [4][2]int{{1, 2}, {7, 3}, {4, 4}, {8, 7}}
		top := ""
		bottom := ""

		for _, pair := range knownPair {
			for i, signal := range signals {
				if len(signal) != pair[1] {
					continue
				}
				segments[pair[0]] = signal
				signals = append(signals[:i], signals[i+1:]...)
				break
			}
		}

		for _, char := range segments[7] {
			if !strings.ContainsRune(segments[1], char) {
				top = string(char)
			}
		}

		diffSignal := segments[4] + top
		for i, signal := range signals {
			if len(signal) != 6 {
				continue
			}

			uncontained, err := "", false
			for _, char := range signal {
				if !strings.ContainsRune(diffSignal, char) {
					if uncontained != "" {
						err = true
						break
					}
					uncontained = string(char)
				}
			}
			if !err {
				bottom = uncontained
				segments[9] = signal
				signals = append(signals[:i], signals[i+1:]...)
				break
			}
		}

		diffSignal = segments[1] + top + bottom
		for loop := 0; loop < 2; loop++ {
			for i, signal := range signals {
				containsAll := true
				for _, char := range diffSignal {
					containsAll = containsAll && strings.ContainsRune(signal, char)
				}

				if containsAll {
					switch len(signal) - len(diffSignal) {
					case 1:
						segments[3] = signal
					case 2:
						segments[0] = signal
					}
					signals = append(signals[:i], signals[i+1:]...)
					break
				}
			}
		}
		for i, signal := range signals {
			if len(signal) == 6 {
				segments[6] = signal
				signals = append(signals[:i], signals[i+1:]...)
				break
			}
		}
		containsAll := true
		for _, char := range signals[0] {
			containsAll = containsAll && strings.ContainsRune(segments[6], char)
		}
		if containsAll {
			segments[5] = signals[0]
			segments[2] = signals[1]
		} else {
			segments[5] = signals[1]
			segments[2] = signals[0]
		}
		input := 0
		for _, value := range inputs {
			for i, segment := range segments {
				if value == segment {
					input = input*10 + i
				}
			}
		}
		answer += input
	}
	fmt.Println(answer)
}
