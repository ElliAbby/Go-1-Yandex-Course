package main

import "fmt"

func main() {
	var ans string
	for {
		fmt.Scanln(&ans)
		if ans == "да" || ans == "нет" || ans == "чёрный" || ans == "белый" {
			fmt.Println("Поражение")
			break
		}
		fmt.Println("Игра продолжается")
	}
}
