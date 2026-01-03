package main

import (
	"fmt"
	"errors"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if nums == nil {
		return nil, errors.New("nums не должен быть nil")
	}
	if n < 0 {
		return nil, errors.New("n не может быть отрицательным")
	}

	resultSlice := make([]int, 0, n)
	for i := 0; i < len(nums); i++ {
		if nums[i] < limit {
			resultSlice = append(resultSlice, nums[i])
			if len(resultSlice) == n {
				break
			}
		}
	}
	return resultSlice, nil
}

func main() {
	nums := []int{4, 7, 89, 3, 21, 2, 5, 7, 32, 4, 6, 8, 0, 3, 4, 6, 2, 115, 12}
	limit := 3
	n := 5
	fmt.Println(UnderLimit(nums, limit, n))
}
