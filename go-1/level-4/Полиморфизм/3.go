package main

import (
	"fmt"
	"reflect"
)

type Vehicle interface {
	CalculateTravelTime(distance float64) float64
}

type Car struct {
	Type string
	Speed float64
	FuelType string
}

func (c Car) CalculateTravelTime(distance float64) float64 {
	return distance / c.Speed
}

type Motorcycle struct {
	Type string
	Speed float64
	FuelType string
}

func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
	return distance / m.Speed
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	result := make(map[string]float64)
	for _, v := range vehicles {
		vehicleType := reflect.TypeOf(v).String()
		result[vehicleType] = v.CalculateTravelTime(distance)
	}
	return result
}

func main() {
	car := Car{Type: "Седан", Speed: 60.0, FuelType: "Бензин"}
	motorcycle := Motorcycle{Type: "Мотоцикл", Speed: 80.0}

	car2 := Car{Type: "Седан 2", Speed: 100.0, FuelType: "Бензин"}
	motorcycle2 := Motorcycle{Type: "Мотоцикл 2", Speed: 120.0}

	vehicles := []Vehicle{car, motorcycle, car2, motorcycle2}

	distance := 200.0

	travelTimes := EstimateTravelTime(vehicles, distance)

	fmt.Println(travelTimes)
	fmt.Printf("Ожидается время для автомобиля %.2f часа\n", travelTimes["main.Car"])
	fmt.Printf("Ожидается время для мотоцикла %.2f часа\n", travelTimes["main.Motorcycle"])
}
