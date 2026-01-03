package main

import (
	"fmt"
	"errors"
)

var Balance float64 = 0.0

// увеличивает баланс на amount
func topUpBalance(amount float64) error {
	if amount <= 0 {
		return errors.New("amount is incorrect")
	}
	Balance += amount
	return nil
}

// уменьшает баланс на amount
func chargeFromBalance(amount float64) error {
	if amount <= 0 || Balance < amount {
		return errors.New("amount is incorrect")
	}
	Balance -= amount
	return nil
}

// вызывают соответствующие функции и возвращают итоговый баланс
func TopUpAndGetBalance(amount float64) (float64, error) {
	err := topUpBalance(amount)
	if err != nil {
		return 0, err
	}

	if Balance < 0 {
		return 0, errors.New("balance is incorrect")
	}
	return Balance, nil
}

func ChargeFromAndGetBalance(amount float64) (float64, error) {
	err := chargeFromBalance(amount)
	if err != nil {
		return 0, err
	}

	if Balance < 0 {
		return 0, errors.New("balance is incorrect")
	}
	return Balance, nil
}

func main() {
	fmt.Println(TopUpAndGetBalance(99999))
	fmt.Println(Balance)

	fmt.Println(topUpBalance(-100))
	fmt.Println(topUpBalance(0))
	fmt.Println(chargeFromBalance(0))
	fmt.Println(chargeFromBalance(-1000))
	fmt.Println(TopUpAndGetBalance(-100))
	fmt.Println(TopUpAndGetBalance(0))
	fmt.Println(ChargeFromAndGetBalance(0))
	fmt.Println(ChargeFromAndGetBalance(-1000))

	Balance -= 100000000
	fmt.Println(TopUpAndGetBalance(1))
}
