// Программист Арсений мечтает создать машину времени. Конечно, работа не простая, но кто-то же должен это сделать.
// Арсений решил начать с малого и хочет написать код, который будет выводить информацию,
// на сколько лет в будущее он переместился.
// Напишите программу, которая считывает с консоли время будущего и настоящего и возвращает указание,
// сколько лет «назад» был нынешний год. Строка должна выглядеть так: X year ago

package main

import (
	"fmt"
	"time"
)

func main() {
	var lastDate, futureDate string
	fmt.Scanln(&futureDate)
	fmt.Scanln(&lastDate)

	layout := "2006-01-02"
	t_f, err_f := time.Parse(layout, futureDate)
	if err_f != nil {
		fmt.Println("Ошибка парсинга:", err_f)
		return
	}

	t_p, err_p := time.Parse(layout, lastDate)
	if err_p != nil {
		fmt.Println("Ошибка парсинга:", err_p)
		return
	}

	y_f := t_f.Year()
	y_p := t_p.Year()
	diff := y_f - y_p
	fmt.Printf("%d year ago", diff)
}
