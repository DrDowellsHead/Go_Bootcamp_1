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
		fmt.Println("Invalid input")
		return
	}
	first, err := parseNumbers(scanner.Text())
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	if !scanner.Scan() {
		fmt.Println("Invalid input")
		return
	}
	second, err := parseNumbers(scanner.Text())
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	common := intersection(first, second)
	if len(common) == 0 {
		fmt.Println("Empty intersection")
		return
	}

	result := make([]string, len(common))
	for i, number := range common {
		result[i] = strconv.Itoa(number)
	}
	fmt.Println(strings.Join(result, " "))
}
