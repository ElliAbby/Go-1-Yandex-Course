package main

import (
	"fmt"
	"time"
)

func main() {
	// получаем текущее время
	now := time.Now()
	fmt.Println("Current time:", now)

	// разница в часах с ДР
	timeOfBirth := time.Date(2006, 8, 6, 9, 45, 30, 0, time.UTC) // год, месяц, день, часы, минуты, секунды и наносекунды
	diff := now.Sub(timeOfBirth)
	fmt.Println(diff.Hours())

	// форматирование времени
	// В качестве аргумента метода передаётся строка с указанием нужного формата.
	// при этом выводиться в соотвествующем формате текущая дата
	fmt.Println("1. Время в формате RFC3339:", now.Format(time.RFC3339))
	fmt.Println("2. Полная дата и время:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("3. Краткая дата:", now.Format("2006-01-02"))
	fmt.Println("4. Время в 24-часовом формате:", now.Format("15:04:05"))
	fmt.Println("5. Время в 12-часовом формате с AM/PM:", now.Format("03:04 PM"))
	fmt.Println("6. День недели:", now.Format("Monday"))
	fmt.Println("7. Сокращённый месяц:", now.Format("Jan"))
}
