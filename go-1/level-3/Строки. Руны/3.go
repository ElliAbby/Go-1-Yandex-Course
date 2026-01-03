package main

import (
	"fmt"
	"unicode"
)

func CheckOnlyASCII(s string) bool {
	for _, letter := range s {
		if letter > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(CheckOnlyASCII("прhello"))
}
