package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

// getFloatInput просит число, пока не получит конечный float64.
func getFloatInput(prompt string) float64 {
	for {
		fmt.Print(prompt)
		var input string
		_, err := fmt.Scanln(&input)
		// EOF означает конец ввода. Просить ввод снова в этой ситуации бессмысленно.
		if err == io.EOF {
			fmt.Println()
			os.Exit(0)
		}
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		// ParseFloat принимает даже NaN и Inf. Через math дополнительно отсеиваем эти значения.
		value, err := strconv.ParseFloat(input, 64)
		if err == nil && !math.IsNaN(value) && !math.IsInf(value, 0) {
			return value
		}
		fmt.Println("Invalid input")
	}
}

// getOperationInput принимает +, -, * или /. Неверный ввод повторяем.
func getOperationInput(prompt string) string {
	for {
		fmt.Print(prompt)
		var input string
		_, err := fmt.Scanln(&input)
		if err == io.EOF {
			fmt.Println()
			os.Exit(0)
		}
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		if input == "+" || input == "-" || input == "*" || input == "/" {
			return input
		}
		fmt.Println("Invalid input")
	}
}
