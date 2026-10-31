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
	scanner.Buffer(make([]byte, 64*1024), 10*1024*1024)

	if !scanner.Scan() {
		if scanner.Err() != nil {
			fmt.Println("Invalid input")
		}
		return
	}
	words := scanner.Text()
	if strings.TrimSpace(words) == "" {
		fmt.Println()
		return
	}

	if !scanner.Scan() {
		fmt.Println("Invalid input")
		return
	}
	k, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil || k < 0 {
		fmt.Println("Invalid input")
		return
	}

	fmt.Println(TopKFrequent(words, k))
}
