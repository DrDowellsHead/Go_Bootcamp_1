package main

import "fmt"

func main() {
	left := getFloatInput("Input left operand: ")
	op := getOperationInput("Input operation: ")

	// Для деления ноль не подходит, поэтому правый операнд читаем в цикле.
	// Для остальных операций ноль допустим.
	var right float64
	for {
		right = getFloatInput("Input right operand: ")
		if op == "/" && right == 0 {
			fmt.Println("Invalid input")
			continue
		}
		break
	}

	// Операцию проверили при вводе, так что здесь всего четыре допустимых случая.
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

	// Три знака после точки по условию нужны именно для деления.
	if op == "/" {
		fmt.Printf("%.3f\n", result)
	} else {
		fmt.Println(result)
	}
}
