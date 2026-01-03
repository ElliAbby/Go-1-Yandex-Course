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

func checkPassword(password string) bool {
	if hasMinimumLength(password, 8) && hasUpper(password) {
		return true
	}
	return false
}

func main() {
	fmt.Println(hasMinimumLength("a", 10))
	fmt.Println(hasMinimumLength("", 0))
	fmt.Println(hasMinimumLength("aaa", 3))
	fmt.Println(hasMinimumLength("a", 1))
	fmt.Println(hasMinimumLength("afaghlonfwdfbomwdfbj", 10))

	fmt.Println(hasUpper("A"))
	fmt.Println(hasUpper(""))
	fmt.Println(hasUpper("f"))
	fmt.Println(hasUpper("wjefbhqjvoqnihbejbhefhqwlbhdw"))
	fmt.Println(hasUpper("AAAAAAAAAAAAAA"))
	fmt.Println(hasUpper("Aerfbjebrfh"))
	fmt.Println(hasUpper("gerfbjebrfhF"))
	fmt.Println(hasUpper("gerfbJebrfhl"))
}
