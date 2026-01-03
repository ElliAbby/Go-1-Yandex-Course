package main

import (
	"unicode"
	"fmt"
)

func hasMinimumLength(password string, minLength int) bool {
	if len(password) >= minLength {
		return true
	}
	return false
}

func hasUpper(password string) bool {
	for _, letter := range password {
		if unicode.IsUpper(letter) {
			return true
		}
	}
	return false
}

func hasLowerCase(password string) bool {
	for _, letter := range password {
		if unicode.IsLower(letter) {
			return true
		}
	}
	return false
}

func ratePassword(password string) string {
	var c1, c2, c3 int = 0, 0, 0
	if hasUpper(password) {
		c1 = 1
	}
	if hasLowerCase(password) {
		c2 = 1
	}
	if hasMinimumLength(password, 8) {
		c3 = 1
	}
	p := c1 + c2 + c3
	switch p {
		case 1: return "Слабый пароль"
		case 2: return "Средний пароль"
		case 3: return "Сложный пароль"
		default: return "Пароль недостаточно безопасен, придумайте новый"
	}
}

func checkPassword(password string) bool {
	if hasMinimumLength(password, 8) && hasUpper(password) {
		return true
	}
	return false
}

func main() {
	fmt.Println(ratePassword("afaghlonfwdfbomwdfbj"))
}
