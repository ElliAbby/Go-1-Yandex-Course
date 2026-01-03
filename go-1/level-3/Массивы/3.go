package main

import "fmt"

func FindMaxMinInArray(array [10]int) (int, int) {
	min := array[0]
	max := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] < min {
			min = array[i]
		} else if array[i] > max {
			max = array[i]
		}
	}
	return min, max
}

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(FindMaxMinInArray(array))
}
