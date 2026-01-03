package main

import "fmt"

func main() {
	// в golang только цикл for
	for i, letter := range "Hello, world!" {
		fmt.Println(i, string(letter))
	}

	s := "Hello, world!"
	for i := 0; i < len(s); i++ {
		fmt.Println(i, string(s[i]))
	}
}
