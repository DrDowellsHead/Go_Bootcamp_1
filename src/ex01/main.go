package main

import "fmt"

func main() {
    left := getFloatInput("Input left operand: ")

    op := getOperationInput("Input operation: ")

    right := getFloatInput("Input right operand: ")

    var result float64

    switch op {
    case "+":
        result = left + right

    case "-":
        result = left - right

    case "*":
        result = left * right

    case "/":
        if right == 0 {
            fmt.Println("Error: Division by zero")
            return
        }
        result = left / right
    }

    fmt.Printf("%.3f\n", result)
}