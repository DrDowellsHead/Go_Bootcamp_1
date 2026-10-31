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
	// Scanner вернул строку, а parseNumbers проверяет и превращает её в []int.
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

	// Пересечение уже сохраняет порядок первого списка — сортировать ничего не нужно.
	common := intersection(first, second)
	if len(common) == 0 {
		fmt.Println("Empty intersection")
		return
	}

	// fmt.Println(common) добавил бы квадратные скобки.
	// Поэтому переводим числа в строки через Itoa и собираем их через Join.
	result := make([]string, len(common))
	for i, number := range common {
		result[i] = strconv.Itoa(number)
	}
	fmt.Println(strings.Join(result, " "))
}
