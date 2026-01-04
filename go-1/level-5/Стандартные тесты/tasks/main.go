package tasks

import (
	"unicode"
	"errors"
	"unicode/utf8"
)

// task 1
func Sum(a, b int) int {
	return a + b
}

// task 2
func Length(a int) string {
    switch {
    case a < 0:
        return "negative"
    case a == 0:
        return "zero"
    case a < 10:
        return "short"
    case a < 100:
        return "long"
    }
    return "very long"
}

// task 3
func Multiply(a, b int) int {
	return a * b
}

// task 4
func DeleteVowels(s string) string {
	newString := ""
	for _, char := range s {
		newChar := unicode.ToLower(char)
		switch newChar {
		case 'a', 'e', 'i', 'o', 'u':
			continue
		default:
			newString += string(char)
		}
	}
	return newString
}

// task 5
var ErrInvalidUTF8 = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
    if !utf8.Valid(input) {
        return 0, ErrInvalidUTF8
    }

    return utf8.RuneCount(input), nil
}