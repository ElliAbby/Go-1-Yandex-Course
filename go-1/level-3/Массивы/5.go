package main

import "fmt"

func PrettyArrayOutput(array [9]string) {
	for i := 0; i < 7; i++ {
		fmt.Printf("%d я уже сделал: %s\n", i + 1, array[i])
	}
	for i := 7; i < 9; i++ {
		fmt.Printf("%d не успел сделать: %s\n", i + 1, array[i])
	}
}

func main() {
	array := [9]string{"ДЗ", "убраться", "фильм", "прога", "PS5", "гулять", "gym", "прочитать", "матан"}
	PrettyArrayOutput(array)
}
