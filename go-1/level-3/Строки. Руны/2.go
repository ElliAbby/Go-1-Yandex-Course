package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func CountLengthAndBytes(first, second string) string {
	words := []string{first, second}
	joinStr := strings.Join(words, "")
	countLength := utf8.RuneCountInString(joinStr)
	countBytes := len(joinStr)
	res := fmt.Sprintf("Объединённая строка: %s. Количество байт: %d. Количество символов: %d.", joinStr, countBytes, countLength)
	return res
}

func main() {
	fmt.Println(CountLengthAndBytes("Hello,", "World"))
}
