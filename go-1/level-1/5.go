package main

import "fmt"

func main() {
	var s, res string
	fmt.Scanln(&s)

	for _, letter := range s {
		if letter != 'а' && letter != 'о' {
			res += string(letter)
		}
	}
	fmt.Println(res)
}
