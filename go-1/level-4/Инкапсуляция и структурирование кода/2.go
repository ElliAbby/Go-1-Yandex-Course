package main

import (
	"fmt"
	"errors"
)

type Account struct {
	balance float64 // приватные поля
	owner string // приватные поля
}

// консруктор
func NewAccount(owner string) *Account {
	return &Account{balance: 0.0, owner: owner}
}


// экспортируемые методы для работы с приватными полями
func (a *Account) SetBalance(value float64) error {
	// устанавливает значение баланса
	if value < 0 {
		return errors.New("value must be poitive")
	}
	a.balance = value
	return nil

}

func (a *Account) GetBalance() float64 {
	// получает баланс
	return a.balance

}

func (a *Account) Deposit(value float64) error {
	// увеличивает баланс
	if value > 0 {
		a.balance += value
		return nil
	}
	return errors.New("value must be poitive")

}

func (a *Account) Withdraw(value float64) error {
	// уменьшает баланс
	if value > 0 && value <= a.balance {
		a.balance -= value
		return nil
	}
	return errors.New("value must be poitive")
}

func main() {
	ex := NewAccount("Andrew")
	_ = ex.SetBalance(10)
	fmt.Println(ex.GetBalance())
	_ = ex.Deposit(10)
	fmt.Println(ex.GetBalance())
	_ = ex.Withdraw(5)
	fmt.Println(ex.GetBalance())
}
