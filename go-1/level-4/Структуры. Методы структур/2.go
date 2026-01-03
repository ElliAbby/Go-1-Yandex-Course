package main

import "fmt"

type Person struct {
	name string
	age int
	address string
}

func (p Person) PrettyPrint() {
	fmt.Printf("Name: %s\nAge: %d\nAddress: %s\n", p.name, p.age, p.address)
}

func main() {
	person := Person{name: "Misha", age: 18, address: "Moscow"}
	person.PrettyPrint()
}
