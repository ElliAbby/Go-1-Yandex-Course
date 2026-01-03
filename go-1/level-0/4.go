package main

import (
	"fmt"
	"math"
)

func main() {
	// const
	fmt.Println(math.Pi)
	fmt.Println(math.E)
	fmt.Println(math.Sqrt2) // корень из 2

	// популярные функции
	fmt.Println(math.Abs(-10))
	fmt.Println(math.Min(-1.0, 10))
	fmt.Println(math.Max(-10.5, 1.5))
	fmt.Println(math.Sqrt(25))
	fmt.Println(math.Sin(45))
	fmt.Println(math.Pow(2, 3))

	// функции округления
	fmt.Println(math.Ceil(5.3)) // округление вверх до целого
	fmt.Println(math.Round(5.3)) // округление по матем правилам
	fmt.Println(math.Floor(5.3)) // округление вниз до целого
}
