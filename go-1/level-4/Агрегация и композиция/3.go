package main

import (
	"fmt"
	"errors"
)

type User struct {
	Name string
	Age int
	IsActive bool
}

func NewUser(name string, age int) (*User, error) {
	if name == "" {
		return nil, errors.New("name is empty for user")
	}

	if age == 0 {
		age = 18
	}
	return &User{Name: name, Age: age, IsActive: true}, nil
}

func main() {
	user, err := NewUser("Elli", 0)
	if err != nil {
		fmt.Println("Ошибка при создании пользователя:", err)
		return
	}
	fmt.Println(*user)

	userEmpty, errEmpty := NewUser("", 27)
	if errEmpty != nil {
		panic(errEmpty)
		return
	}
	fmt.Println(*userEmpty)
}
