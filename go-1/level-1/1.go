package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)

	if a == 0 {
		fmt.Println("Число 0")
	} else if (a > 0 && a < 10) || (a < 0 && a > -10) {
		fmt.Println("Число однозначное")
	} else if a % 2 == 0 {
		fmt.Println("Число чётное")
	} else if a > 0 {
		fmt.Println("Число положительное")
	} else if a < -9 && a % 2 != 0 {
		fmt.Println("Число красивое")
	}
}
