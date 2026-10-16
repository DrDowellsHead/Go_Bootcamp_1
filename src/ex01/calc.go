package ex01

import (
	"fmt"
	"strconv"
)

func calc() error {
	var left_operand float64
	
	for {
		fmt.Println("Введите левый операнд: ")
		left_operand = strconv.ParseFloat(left_operand, 64)
			if err != nil {
				fmt.Println("Invalid input", err)
				continue
			} else {
				break
			}

	fmt.Printf(&left_operand)
}