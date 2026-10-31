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
	// Обычного буфера Scanner может не хватить для длинного списка слов.
	// Начинаем с 64 КиБ, разрешаем рост до 10 МиБ.
	scanner.Buffer(make([]byte, 64*1024), 10*1024*1024)

	// Сначала читаем строку со словами. Ошибку Scanner отличаем от конца ввода.
	if !scanner.Scan() {
		if scanner.Err() != nil {
			fmt.Println("Invalid input")
		}
		return
	}
	words := scanner.Text()
	// Пустой список должен дать пустой результат. K тогда даже не читаем.
	if strings.TrimSpace(words) == "" {
		fmt.Println()
		return
	}

	if !scanner.Scan() {
		fmt.Println("Invalid input")
		return
	}
	// Atoi разбирает строку целиком: "2 abc" уже не пройдёт,
	// в отличие от прежнего варианта с Sscan.
	k, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil || k < 0 {
		fmt.Println("Invalid input")
		return
	}

	fmt.Println(TopKFrequent(words, k))
}
