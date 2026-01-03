package main

import "fmt"

type Animal interface {
	MakeSound() string
	GetName() string
	GetInfo() string
}

type animal struct {
	name string
	species string
	age int
	sound string
}

func (a *animal) MakeSound() string {
	return a.sound
}

func (a *animal) GetName() string {
	return a.name
}

func (a *animal) GetInfo() string {
	resultMessage := fmt.Sprintf("Имя: %s, Вид: %s, Возраст: %d", a.GetName(), a.species, a.age)
	return resultMessage
}

type ZooKeeper struct {}

func (z ZooKeeper) Feed(animal Animal) {
	resultMessage := fmt.Sprintf("Смотритель зоопарка кормит %s. %s!", animal.GetName(), animal.MakeSound())
	fmt.Println(resultMessage)
}

func NewAnimal(name, species string, age int, sound string) Animal {
	return &animal{name, species, age, sound}
}

func ZooShow(animals []Animal) {
	for _, a := range animals {
		fmt.Println(a.GetInfo())
		fmt.Println(a.MakeSound())
	}
}

func main() {
	tiger := NewAnimal("Барсик", "Тигр", 5, "Ррр")
	penguin := NewAnimal("Пиня", "Пингвин", 2, "Кря")
	ZooShow([]Animal{tiger, penguin})

	keeper := ZooKeeper{}
	keeper.Feed(tiger)
}
