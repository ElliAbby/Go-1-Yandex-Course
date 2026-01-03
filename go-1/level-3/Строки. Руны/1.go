package main

import "fmt"

func main() {
	name := "異體字"
	firstLetter := name[0]
	fmt.Println(firstLetter)         // Вывод: 231
	fmt.Println(string(firstLetter)) // Вывод: ç

	name2 := "異體字"
	firstLetter2 := []rune(name2)[0] // Конвертируем строку в слайс рун
	// И с помощью [0] забираем первый элемент
	fmt.Println(firstLetter2) // Вывод: 30064
	fmt.Println(string(firstLetter2)) // Вывод: 異
}

// utf8.RuneCountInString в отличии от функции len() она выводит не количество байт в строке
// (что, как мы поняли, не всегда совпадает с количеством символов), а именно количество рун
// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	name := "你好 Андрей"
// 	fmt.Println(utf8.RuneCountInString(name)) // Вывод: 9 (9 символов, включая пробел)
// 	fmt.Println(len(name))                    // Вывод: 19 (количество байтов, далеко от нужного нам значения длины строки)
// }
