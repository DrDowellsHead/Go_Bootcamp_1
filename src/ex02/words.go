package main

import (
	"sort"
	"strings"
)

// WordFreq держит слово и число повторений рядом, чтобы сортировать их вместе.
type WordFreq struct {
	Word  string
	Count int
}

// TopKFrequent выбирает K самых частых слов.
// При равенстве частот раньше идёт лексикографически меньшее слово.
func TopKFrequent(words string, k int) string {
	if words == "" || k <= 0 {
		return ""
	}
	// Счётчики складываем в map: поиск слова в среднем занимает O(1).
	freqMap := make(map[string]int)
	// Разделитель в условии — пробел. Пустые части от нескольких пробелов пропускаем.
	for _, word := range strings.Split(words, " ") {
		if word != "" {
			freqMap[word]++
		}
	}
	// Map не гарантирует порядок обхода. Для сортировки переносим пары в срез.
	// Длина пока ноль, зато вместимость сразу под все уникальные слова.
	uniqueWords := make([]WordFreq, 0, len(freqMap))
	for word, count := range freqMap {
		uniqueWords = append(uniqueWords, WordFreq{Word: word, Count: count})
	}
	// Сортируем по убыванию частоты, при равенстве — по слову.
	// Для u уникальных слов сложность сортировки O(u log u).
	sort.Slice(uniqueWords, func(i, j int) bool {
		if uniqueWords[i].Count != uniqueWords[j].Count {
			return uniqueWords[i].Count > uniqueWords[j].Count
		}
		return uniqueWords[i].Word < uniqueWords[j].Word
	})
	// Ограничиваем K числом уникальных слов, иначе выйдем за границы среза.
	if k > len(uniqueWords) {
		k = len(uniqueWords)
	}
	result := make([]string, k)
	for i := 0; i < k; i++ {
		result[i] = uniqueWords[i].Word
	}
	return strings.Join(result, " ")
}
