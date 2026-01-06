package main

import (
	"fmt"
	"slices"
)

func SortAndMerge(left, right []int) []int {
	slices.Sort(left)
	slices.Sort(right)

	result := make([]int, 0, len(left)+len(right))
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		switch {
		case left[l] <= right[r]:
			result = append(result, left[l])
			l++
		default:
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

func main() {
	fmt.Println(
		SortAndMerge([]int{1, 2, 3, 4, 5, 9}, []int{1, 2, 3, 4, 5, 8}),
	)
}