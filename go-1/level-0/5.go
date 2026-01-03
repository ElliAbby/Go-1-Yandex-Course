package main

import (
    "fmt"
    "math"
)

func main() {
    var x, y, z, m, n float64
    fmt.Scanln(&x)
    fmt.Scanln(&y)
    fmt.Scanln(&z)
    fmt.Scanln(&m)
    fmt.Scanln(&n)
    // (5x * sin(2y)) / (z + n ^ ln(m))
    res := (5*x * math.Sin(2*y)) / (z + math.Pow(n, math.Log(m)))
    fmt.Println(res)
}
