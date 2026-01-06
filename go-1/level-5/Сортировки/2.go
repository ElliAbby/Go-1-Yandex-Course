package main

import (
	"fmt"
	"slices"
	"sort"
)

func SortNums(nums []uint) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
}

func SortNums2(nums []uint) {
	slices.Sort(nums)
}

func main() {
	s := []uint{1, 2, 3, 4, 124, 45, 0}
	SortNums(s)
	fmt.Println(s)
	s2 := []uint{1, 2, 3, 4, 124, 45, 0}
	SortNums2(s2)
	fmt.Println(s2)
}