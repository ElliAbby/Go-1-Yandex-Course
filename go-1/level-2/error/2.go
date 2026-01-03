package main

import "fmt"

func main() {
    fmt.Println("Start")
    defer func() { // Отложенный вызов recover()
        if r := recover(); r != nil { // Проверка, что recover() != nil (паника отловлена)
            fmt.Println("Recovered from panic:", r)
        }
    }()

    panic("Something went wrong") // Намеренный вызов паники
    fmt.Println("End")
}
