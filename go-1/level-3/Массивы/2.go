package main

import "fmt"

func ThirdElementInArray(array [6]int) int { // 0 1 2 3 4 5
	return array[2]
}

func main() {
	array := [6]int{23, 11, 45, 67, 7, 8}
	fmt.Println(ThirdElementInArray(array))
}
