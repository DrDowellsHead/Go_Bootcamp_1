package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    if !scanner.Scan() {
        return
    }
    words := scanner.Text()

    if !scanner.Scan() {
        return
    }

    var k int
    _, err := fmt.Sscan(scanner.Text(), &k)

    if err != nil || k < 0 {
        fmt.Println("Invalid input")
        return
    }

    fmt.Println(TopKFrequent(words, k))
}