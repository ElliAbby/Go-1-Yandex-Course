package main

import "fmt"

func CountingSort(contacts []string) map[string]int {
	resultMap := make(map[string]int)

	for _, el := range contacts {
		resultMap[el]++
	}
	return resultMap
}

func main() {
	fmt.Println(CountingSort([]string{"apple", "tomato"})) // map[apple:1 tomato:1]
	fmt.Println(CountingSort([]string{"ef", "ef", "ef", "ef", "ef22"})) // map[ef:4 ef22:1]
}
