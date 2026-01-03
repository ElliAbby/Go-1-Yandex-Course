package main

import "fmt"

func SumOfArray(array [6]int) int {
	total := 0
	for _, element := range array {
		total += element
	}
	return total
}

func main() {
	fmt.Println(SumOfArray([6]int{1, 2, 3, 4, 5, 6}))
}
