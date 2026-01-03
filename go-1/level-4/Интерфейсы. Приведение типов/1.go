package main

import "fmt"

type Rectangle struct {
	Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Shape interface {
	Area() float64
}

func CalculateArea(s Shape) float64 {
	return s.Area()
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	area := CalculateArea(rect)
	fmt.Println(area)
}
