package main

import (
	"strings"
	"fmt"
)

func CheckLetters(text string) string {
	count := strings.Count(text, "е")
	if count == 0 {
		return "Текст готов к публикации!"
	}
	return fmt.Sprintf("Количество возможных ошибок: %d, перепроверьте текст", count)
}

func main() {
	fmt.Println(CheckLetters(""))
	fmt.Println(CheckLetters("Тут точно нет сложных букв, хотя ж выглядит подозрительно"))
	fmt.Println(CheckLetters("Ох уж эти е и ё"))
	fmt.Println(CheckLetters("Е ее ее тум турун тум турун тум"))
}
