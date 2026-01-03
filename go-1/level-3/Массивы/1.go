package main

import "fmt"

func FiveSteps(array [5]int) [5]int {
	var resultArray [5]int
	for i := 0; i < 5; i++ {
		resultArray[i] = array[4 - i]
	}
	return resultArray
}

func main() {
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(FiveSteps(array))
	fmt.Println(len(array))
}
