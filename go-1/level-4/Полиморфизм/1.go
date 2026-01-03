package main

import "fmt"

type Animal interface {
    Speak() string
}

type Dog struct {
    Name string
}

func (d *Dog) Speak() string {
    return "Woof!"
}

type Cat struct {
    Name string
}

func (c *Cat) Speak() string {
    return "Meow!"
}

type Cow struct {
    Name string
}

func (c *Cow) Speak() string {
    return "Moo!"
}

func main() {
    animals := []Animal{
        &Dog{"Fido"},
        &Cat{"Fluffy"},
        &Cow{"Betsy"},
    }

    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}

// Мы определяем интерфейс Animal с методом Speak().
// Затем — три типа, Dog, Cat и Cow, каждый из которых реализует один и тот же метод Speak() интерфейса Animal,
// но поведение в этих методах разное.
