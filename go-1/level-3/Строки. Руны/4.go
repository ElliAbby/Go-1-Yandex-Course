package main

import "fmt"

func NumbersToLetters(input string) string {
	dictNumToLet := map[string]string {
		"0": "ноль",
		"1": "один",
		"2": "два",
		"3": "три",
		"4": "четыре",
		"5": "пять",
		"6": "шесть",
		"7": "семь",
		"8": "восемь",
		"9": "девять",
		"+": "плюс",
		"-": "минус",
		"*": "умножить на",
		"/": "разделить на",
	}
	var result string
	for _, letter := range input {
		stChar := string(letter)
		if dictNumToLet[stChar] != "" {
			result += dictNumToLet[stChar]
		} else {
			result += stChar
		}
	}
	return result
}

func main() {
	fmt.Println(NumbersToLetters("(1 + 2) * 3 / 8"))
	fmt.Println(NumbersToLetters("2 + 2 * 2"))
	fmt.Println(NumbersToLetters("один - 2 / 2"))
	fmt.Println(NumbersToLetters("Хочу попасть на стажировку в Яндекс!"))
	fmt.Println(NumbersToLetters("1 2 3 4 5 6 7 8 9 + + - * /"))
}
