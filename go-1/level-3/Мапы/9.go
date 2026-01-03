package main

import "fmt"

func DeleteLongKeys(m map[string]int) map[string]int {
	for key := range m {
		if len([]rune(key)) < 6 {
			delete(m, key)
		}
	}
	return m
}

func main() {
	m := map[string]int{
		"Longest": 10,
		"Middle": 12,
		"Small": 9,
	}
	fmt.Println(DeleteLongKeys(m))
}
