package main

import (
	"fmt"
	"math"
)

func  findDiscriminant(a, b, c float64) float64 {
	discriminant := math.Pow(b, 2) - 4.0 * a * c
	return discriminant
}

func SquareRoots(a, b, c float64) (float64, float64) {
	discriminant := findDiscriminant(a, b, c)
	if discriminant < 0 {
		return 0, 0
	} else {
		x1 := (-1 * b - math.Sqrt(discriminant)) / float64(2 * a)
		x2 := (-1 * b + math.Sqrt(discriminant)) / float64(2 * a)
		if x1 > x2 {
			return x2, x1
		}
		return x1, x2
	}
}

func main() {
	fmt.Println(SquareRoots(1, -3, 2))
	fmt.Println(SquareRoots(1, 4, 4.0))
	fmt.Println(SquareRoots(1, 1, 1))
	fmt.Println(SquareRoots(4, 4, 1))
	fmt.Println(SquareRoots(4, 4, -1))
	fmt.Println(SquareRoots(2, 2, 8))
	fmt.Println(SquareRoots(1, 4, -5))
	fmt.Println(SquareRoots(1, 0, -9))
	fmt.Println(SquareRoots(1, 5, 0))
}
