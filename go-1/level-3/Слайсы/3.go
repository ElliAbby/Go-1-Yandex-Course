package main

import "fmt"

func SliceCopy(nums []int) []int {
	lenght := len(nums)
	newSlice := make([]int, lenght, lenght)
	copy(newSlice, nums)
	return newSlice
}

func main() {
	fmt.Println(SliceCopy([]int{1, 2, 3}))
}
