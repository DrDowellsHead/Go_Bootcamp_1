package ex01

import (
	"fmt"
	"strconv"
)

func calc() {
	var left_operand string
	
	for {
		fmt.Println("Введите левый операнд: ")
		fmt.Scan(&left_operand)
		left_operand = strconv.ParseFloat(left_operand string, 64)
			if err != nil {
				fmt.Println("Invalid input", err)
				continue
			} else {
				break
			}
		}

	fmt.Println(&left_operand)
}