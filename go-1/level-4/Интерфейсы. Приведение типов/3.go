package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width float64
	height float64
}

func (c Circle) Area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	var s Shape

	s = Circle{radius: 5}
	fmt.Printf("Площадь круга: %.2f\n", s.Area())

	s = Rectangle{width: 4, height: 3}
	fmt.Printf("Площадь прямоугольника: %.2f\n", s.Area())
}
