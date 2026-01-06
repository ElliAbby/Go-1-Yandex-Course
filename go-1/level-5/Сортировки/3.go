package main

import (
	"fmt"
	"slices"
)

func SortNames(names []string) {
	slices.Sort(names)
}

func SortNames2(names []string) {
	slices.SortFunc(names, func(a, b string) int {
		switch a < b {
		case true:
			return -1
		case false:
			return 1
		default:
			return 0
		}
	})
}

func main() {
	names := []string{"Арина", "Варвара", "Есения", "Аксинья"}
	SortNames(names)
	fmt.Println(names)

	names2 := []string{"Арина", "Варвара", "Есения", "Аксинья"}
	SortNames2(names2)
	fmt.Println(names2)
}