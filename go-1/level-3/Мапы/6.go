package main

import "fmt"

func SumOfValuesInMap(m map[int]int) int {
	total := 0
	for _, value := range m {
		total += value
	}
	return total
}

func main() {
	m := map[int]int{
		10: 10,
		20: 20,
	}
	fmt.Println(SumOfValuesInMap(m))
}
