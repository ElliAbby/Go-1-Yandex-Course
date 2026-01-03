package main

import "fmt"

// 1 способ
func Join1(nums1, nums2 []int) []int {
	lenghtOne := len(nums1)
	lenghtTwo := len(nums2)
	newSlice := make([]int, lenghtOne + lenghtTwo, lenghtOne + lenghtTwo)
	for i := 0; i < lenghtOne; i++ {
		newSlice[i] = nums1[i]
	}
	for i := 0; i < lenghtTwo; i++ {
		newSlice[lenghtOne + i] = nums2[i]
	}
	return newSlice
}

// 2 способ
func Join2(nums1, nums2 []int) []int {
	lenghtOne := len(nums1)
	lenghtTwo := len(nums2)
	newSlice := make([]int, lenghtOne + lenghtTwo, lenghtOne + lenghtTwo)
	copy(newSlice[:lenghtOne], nums1)
	copy(newSlice[lenghtOne:], nums2)
	return newSlice
}

func main() {
	fmt.Println(Join1([]int{1, 2, 3}, []int{2, 3})) // [1 2 3 2 3]
	fmt.Println(Join2([]int{1, 2, 3}, []int{2, 3})) // [1 2 3 2 3]
}
