package main

import "fmt"

func Mix(nums []int) []int {
	middle := len(nums) / 2
	newSlice := make([]int, len(nums))
	for i := 0; i < middle; i++ {
		newSlice[2*i] = nums[i]
		newSlice[2*i + 1] = nums[middle + i]
	}
	return newSlice
}

func main() {
	fmt.Println(Mix([]int{1, 2, 3, 4, 5, 6}))
}
