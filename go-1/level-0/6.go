package main

import (
	"fmt"
	"time"
)

func main() {
	var dateStr string
	fmt.Scanln(&dateStr)

	layout := "2006-01-02/15:04:05"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Ошибка парсинга:", err)
		return
	}

	hour := t.Hour()
	minute := t.Minute()

	fmt.Printf("Текущее время %d часов, %d минут. Ты точно не забыл про важные дела на сегодня?\n", hour, minute)
}
