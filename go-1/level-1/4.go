// Считается, что много повторений одних и тех же слов в одном тексте — плохо.
// К сожалению, программист Арсений нередко этим грешит.
// Помогите ему и напишите программу, которая будет принимать
// на вход число слов в тексте (без знаков препинания и специальных символов),
// слово, повторения которого нужно подсчитать, и сам текст,
// а выводить количество повторений.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)

	var target string
	fmt.Scan(&target)
	target = strings.ToLower(target)

	var s string
	count := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		s = strings.ToLower(s)
		if len(s) == len(target) {
			j := 0
			for ; j < len(s); j++ {
				if target[j] != s[j] {
					break
				}
			}
			if j == len(s) {
				count++
			}
		}
	}
	fmt.Println(count)
}
