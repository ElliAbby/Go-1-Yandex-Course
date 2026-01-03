package main

import "fmt"

func PrintComplexNumber(z complex64) {
	r := real(z)
	i := imag(z)
	fmt.Println(fmt.Sprintf("Действительная часть: %.2f. Мнимая часть: %.2f", r, i))
}

func main() {
	var z complex64 = 3+4i

	PrintComplexNumber(z)
}
