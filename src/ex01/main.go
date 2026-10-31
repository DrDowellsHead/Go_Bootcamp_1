package main

import "fmt"

func main() {
	left := getFloatInput("Input left operand: ")
	op := getOperationInput("Input operation: ")

	var right float64
	for {
		right = getFloatInput("Input right operand: ")
		if op == "/" && right == 0 {
			fmt.Println("Invalid input")
			continue
		}
		break
	}

	var result float64
	switch op {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		result = left / right
	}

	if op == "/" {
		fmt.Printf("%.3f\n", result)
	} else {
		fmt.Println(result)
	}
}
