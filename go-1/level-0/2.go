package main

import (
    "fmt"
)

func main() {
    var name string
    fmt.Scanln(&name)

    var numFlat int
    fmt.Scanln(&numFlat)

    var password int
    fmt.Scanln(&password)

    var time float64
    fmt.Scanln(&time)

    resultMessage := fmt.Sprintf("Привет, %s! Приглашаю тебя на соревнование по программированию, которое пройдёт, как всегда, в квартире %d. Оно будет длиться примерно %.1f часа. Не забудь секретный пароль для входа: %d.", name, numFlat, time, password)
	fmt.Println(resultMessage)
}
