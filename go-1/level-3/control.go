package main

import (
	"fmt"
	"sort"
	"strings"
)

func AnalyzeText(text string) {
	separation := []string{",", ".", "!", "?", "  "}
	for _, el := range separation {
		text = strings.Replace(text, el, " ", -1)
	}
	words := strings.Split(text, " ")

	cleanWords := make([]string, 0, len(words))
	for i := 0; i < len(words); i++ {
		if words[i] != "" && words[i] != " " {
			cleanWords = append(cleanWords, words[i])
		}
	}

	wordCount := make(map[string]int)
	for _, word := range cleanWords {
		lower := strings.ToLower(word)
		if lower != "" {
			wordCount[lower]++
		}
	}

	totalWords := len(cleanWords)
	uniqueWords := len(wordCount)

	topWords := getTopWords(wordCount, 5)
	mostFrequent := topWords[0]
	maxCount := wordCount[mostFrequent]

	fmt.Printf("Количество слов: %d\n", totalWords)
	fmt.Printf("Количество уникальных слов: %d\n", uniqueWords)
	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", mostFrequent, maxCount)
	fmt.Println("Топ-5 самых часто встречающихся слов:")
	for _, word := range topWords {
		count := wordCount[word]
		fmt.Printf("\"%s\": %d раз\n", word, count)
	}
}

func getTopWords(wordMap map[string]int, n int) []string {
	swapWordMap := make(map[int]string)
	for key, value := range wordMap {
		swapWordMap[value] = key
	}

	sliceOfKey := make([]int, 0, len(swapWordMap))
	for key := range swapWordMap {
		sliceOfKey = append(sliceOfKey, key)
	}
	sort.Ints(sliceOfKey)

	swapSliceOfKey := make([]int, 0, len(swapWordMap))
	for i := 0; i < len(sliceOfKey); i++ {
		swapSliceOfKey = append(swapSliceOfKey, sliceOfKey[len(sliceOfKey) - i - 1])
	}

	result := make([]string, n)
	for i := 0; i < n; i++ {
		result[i] = swapWordMap[swapSliceOfKey[i]]
	}
	return result
}

func main() {
	AnalyzeText("one two two three three three four four four four five five five five five")
	AnalyzeText("Go очень очень очень ОЧЕНЬ ОчЕнь очень оЧЕНь классный классный! go просто, ну просто классный. GO Классный!")
	AnalyzeText("Я так люблю море. Я на море. Я так люблю. Море! Я море!!! ЛЮБЛЮ МОРЕ")
}
