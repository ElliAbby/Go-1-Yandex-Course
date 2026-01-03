package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("Введите свое имя:")
	fmt.Scanln(&name)

	var age int
	fmt.Println("Введите свой возраст:")
	fmt.Scanln(&age)

	var mark float32
	fmt.Println("Введите свой средний балл по любимому предмету:")
	fmt.Scanln(&mark)

	const secretNumber = 54423617

	fmt.Println(
			fmt.Sprintf(
				"Привет, я робот-предсказатель! Вижу, что тебе сейчас %d, а зовут тебя %s. Уникальное число твоего предсказания %d. А средняя оценка по твоему любимому предмету %f. Как я догадался? Пусть это останется секретом моего мастерства!",
				age,
				name,
				secretNumber,
				mark,
        ),
    )
}
