package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var action string
    var count int

    depth, position, aim := 0, 0, 0

    for scanner.Scan() {
        _, err := fmt.Sscanf(scanner.Text(), "%s %d", &action, &count)

        if err != nil {
            panic(err)
        }

        switch action {
        case "forward":
            position += count
            depth += count * aim
        case "up":
            aim -= count
        case "down":
            aim += count
        }
    }

    fmt.Println(position * depth)
}
