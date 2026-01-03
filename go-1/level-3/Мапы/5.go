package main

import (
	"fmt"
	"math"
)

func FindMaxKey(m map[int]int) int {
	maxKey := int(math.Inf(-1))
	for key, _ := range m {
		if key > maxKey {
			maxKey = key
		}
	}
	return maxKey
}

func main() {
	m := map[int]int{
		10: 37,
		3: 19,
	}

	fmt.Println(FindMaxKey(m))
}
