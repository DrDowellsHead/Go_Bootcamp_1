package main

import (
	"strconv"
	"strings"
)

// parseNumbers читает целые числа из строки; при первой ошибке
// возвращает err вместо частично собранного списка.
func parseNumbers(line string) ([]int, error) {
	var numbers []int
	// Fields разбивает строку по пробелам и табуляциям, не создавая пустых слов.
	for _, word := range strings.Fields(line) {
		// Atoi превращает строку в int, отбрасывает дроби и слишком большие числа.
		number, err := strconv.Atoi(word)
		if err != nil {
			return nil, err
		}
		// Изначально срез nil, но append умеет добавлять в него элементы.
		numbers = append(numbers, number)
	}
	return numbers, nil
}

// intersection возвращает уникальные общие числа в порядке первого среза.
func intersection(first, second []int) []int {
	// Сначала собираем числа второго списка в map.
	// Поиск по ключу в среднем O(1), поэтому вложенный цикл не нужен.
	inSecond := make(map[int]bool)
	for _, number := range second {
		inSecond[number] = true
	}

	// seen отмечает уже добавленные числа, чтобы не выводить дубликаты.
	seen := make(map[int]bool)
	var result []int
	// Идём по первому срезу, поэтому порядок результата сохраняется сам.
	// В среднем сложность всего поиска O(len(first) + len(second)).
	for _, number := range first {
		// Отсутствующий ключ в map[int]bool даёт false.
		// Добавляем только то, что есть во втором списке и чего ещё нет в seen.
		if inSecond[number] && !seen[number] {
			result = append(result, number)
			seen[number] = true
		}
	}
	return result
}
