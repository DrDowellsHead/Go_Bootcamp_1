package main

import (
    "fmt"
    "io"
    "os"
    "strconv"
)

func getFloatInput(prompt string) float64 {
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

        value, err := strconv.ParseFloat(input, 64)

        if err == nil {
            return value
        }

        fmt.Println("Invalid input")
    }
}

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

        if input == "+" || input == "-" ||
            input == "*" || input == "/" {
            return input
        }

        fmt.Println("Invalid input")
    }
}