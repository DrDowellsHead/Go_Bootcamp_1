package main

import (
	"strconv"
	"strings"
)

func parseNumbers(line string) ([]int, error) {
	var numbers []int
	for _, word := range strings.Fields(line) {
		number, err := strconv.Atoi(word)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

func intersection(first, second []int) []int {
	inSecond := make(map[int]bool)
	for _, number := range second {
		inSecond[number] = true
	}

	seen := make(map[int]bool)
	var result []int
	for _, number := range first {
		if inSecond[number] && !seen[number] {
			result = append(result, number)
			seen[number] = true
		}
	}
	return result
}
