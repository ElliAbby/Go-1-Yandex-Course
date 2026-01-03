package main

import "fmt"

type Student struct {
    name string
    age  int
}

// ресивер по значению -> объект копируется
func (s Student) printData() {
    fmt.Printf("Name: %s, Age: %d\n", s.name, s.age)
}

// ресивер по указателю
func (s *Student) printDataPoint() {
    fmt.Printf("Name: %s, Age: %d\n", s.name, s.age)
}

func main() {
    student1 := Student{name: "vasya", age: 15}
    student2 := Student{name: "petya", age: 20}
    student3 := Student{name: "kolya", age: 25}
    student4 := Student{name: "masha", age: 30}
    student5 := Student{name: "dasha", age: 35}
    student6 := Student{name: "sasha", age: 40}

    fmt.Println(student1)
    fmt.Println(student2)
    fmt.Println(student3)
    fmt.Println(student4)
    fmt.Println(student5)
    fmt.Println(student6)

    student1.printData()
    student2.printDataPoint()
}
