package main

import (
	"sort"
	"strings"
)

type WordFreq struct {
	Word  string
	Count int
}

func TopKFrequent(words string, k int) string {
	if words == "" || k <= 0 {
		return ""
	}
	freqMap := make(map[string]int)
	for _, word := range strings.Split(words, " ") {
		if word != "" {
			freqMap[word]++
		}
	}
	uniqueWords := make([]WordFreq, 0, len(freqMap))
	for word, count := range freqMap {
		uniqueWords = append(uniqueWords, WordFreq{Word: word, Count: count})
	}
	sort.Slice(uniqueWords, func(i, j int) bool {
		if uniqueWords[i].Count != uniqueWords[j].Count {
			return uniqueWords[i].Count > uniqueWords[j].Count
		}
		return uniqueWords[i].Word < uniqueWords[j].Word
	})
	if k > len(uniqueWords) {
		k = len(uniqueWords)
	}
	result := make([]string, k)
	for i := 0; i < k; i++ {
		result[i] = uniqueWords[i].Word
	}
	return strings.Join(result, " ")
}
