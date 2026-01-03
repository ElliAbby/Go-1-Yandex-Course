package main

import "fmt"

type Employee struct {
	name string
	position string
	salary float64
	bonus float64
}

func (e Employee) CalculateTotalSalary() {
	fmt.Printf("Employee: %s\nPosition: %s\nTotal Salary: %.2f\n", e.name, e.position, e.salary + e.bonus)
}


func main() {
	arseniy := Employee{
		name: "Arseniy",
		position: "backend developer",
		salary: 1000.0,
		bonus: 0.02345,
	}

	arseniy.CalculateTotalSalary()
}
