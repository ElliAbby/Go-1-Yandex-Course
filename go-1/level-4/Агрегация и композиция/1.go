// При агрегации один объект содержит в себе другой как свою часть.
// Однако внутренний объект не абсолютно зависим от внешнего.
// Например, автомобильный двигатель может существовать и сам по себе, без автомобиля.

// Агрегация в Go происходит, когда структура включает в себя другую структуру как своё поле
package main

import "fmt"

type Engine struct {
    Model string
}

type Car struct {
    Make   string
    Engine Engine
}

func main() {
    engine := Engine{Model: "V6"}
    car := Car{Make: "Toyota", Engine: engine}

    fmt.Printf("Car Make: %s\n", car.Make)
    fmt.Printf("Engine Model: %s\n", car.Engine.Model)
}


// При композиции один объект — это часть другого и не может существовать сам по себе.
// Например, если развязать узел на надутом воздушном шарике, воздух из него выйдет, сольётся с окружающей атмосферой
// и не сможет сам по себе снова надуть шарик.

// Композиция в Go происходит, когда одна структура встраивается в другую.
// При этом имя структуры внутри внешней не указывается явно.
package main

import "fmt"

type Engine struct {
    Model string
}

type Car struct {
    Make string
    Engine // Встраивание структуры
}

func main() {
    car := Car{Make: "Ford", Engine: Engine{Model: "V8"}}

    fmt.Printf("Car Make: %s\n", car.Make)
    fmt.Printf("Engine Model: %s\n", car.Model) // Мы можем обращаться к Engine как к полю Car
}
