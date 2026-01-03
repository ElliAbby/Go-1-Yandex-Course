package main

import (
	"errors"
	"fmt"
)

var DivisionByZeroError = errors.New("division by zero is not allowed")

// Функция, которая вторым аргументом возвращает ошибку
func Divide(a, b int) (float64, error) {
 	if b == 0 {
  		// Возврат ошибки вторым значением
  		return 0, DivisionByZeroError
  	}
   // Возврат nil вторым значением, если ошибка не произошла
   	return float64(a) / float64(b), nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil { // Проверка, что err не равно nil (в функции divide произошла ошибка)
        fmt.Println("Ошибка:", err)
        return
    }
    fmt.Println("Результат:", result)
}
