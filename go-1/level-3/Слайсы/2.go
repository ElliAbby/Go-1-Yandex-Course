package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int{4, 5, 2, 1, 3, 9, 7, 8, 6}
	fmt.Println(sort.IntsAreSorted(intSlice)) // Проверим, отсортирован ли слайс
	sort.Ints(intSlice)	// Сама сортировка
	fmt.Println(intSlice) // [1 2 3 4 5 6 7 8 9]

	stringSlice := []string{"Go", "Bravo", "Gopher", "YaLyceum", "Alpha", "Grin", "Delta"}
	sort.Strings(stringSlice) // Сортировка уже для строк
	fmt.Println(stringSlice) // [Alpha Bravo Delta Go Gopher Grin YaLyceum]
}
