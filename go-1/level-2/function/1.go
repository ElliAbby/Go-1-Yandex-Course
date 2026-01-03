package main

import "fmt"

func printPrice(price int, product string) {
	fmt.Printf("%s будет стоить %d рублей\n", product, price)
}

func BuyFries(size string) {
	var price int = 0
	switch size {
		case "S": price = 49
		case "M": price = 89
		case "L": price = 99
		default: fmt.Println("Некорректный размер")
	}
	if price != 0 {
		printPrice(price, "Картошка фри")
	}
}

func BuyCola(size string) {
	var price int = 0
	switch size {
		case "S": price = 99
		case "M": price = 119
		case "L": price = 139
		default: fmt.Println("Некорректный размер")
	}
	if price != 0 {
		printPrice(price, "Кола")
	}
}

func main() {
	BuyCola("L")
	BuyFries("S")
	BuyCola("XXL")
}
