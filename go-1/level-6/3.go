package main

import (
	"fmt"
)

const (
	minPrice = 99.0
	maxPrice = 20000
)

func ApplyPriceLimits(price float64) float64 {
	if price < minPrice {
		return minPrice
	} 
	if price > maxPrice {
		return maxPrice
	}
	return price
}

func main() {
	for _, price := range []float64{100, 1000, 10000, 100000} {
		fmt.Printf("Цена %d рублей: %d рублей\n", int(price), int(ApplyPriceLimits(price)))
	}
}