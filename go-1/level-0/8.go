package main

import (
	"fmt"
	"time"
	"math"
)

func main() {
	var date, name, lastName, midName string
	var salary1, salary2, salary3 float64

	fmt.Scanln(&date)
	fmt.Scanln(&name)
	fmt.Scanln(&lastName)
	fmt.Scanln(&midName)
	fmt.Scanln(&salary1)
	fmt.Scanln(&salary2)
	fmt.Scanln(&salary3)

	total := salary1 + salary2 + salary3
	totalKop := int64(math.Round(total * 100))
	rub := totalKop / 100
	kop := totalKop % 100

	layout := "02.01.2006"
	startDate, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Ошибка парсинга даты", err)
		return
	}

	signDate := startDate.AddDate(0, 0, 15) // + 15 дней
	signDatestr := signDate.Format("02.01.2006")

	fmt.Printf("Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\n", lastName, name, midName)
	fmt.Printf("Дата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день.\n", signDatestr)
	fmt.Printf("Общая сумма выплат составит %d руб. %d коп.\n", rub, kop)
	fmt.Println()
	fmt.Println("С уважением,\nГл. бух. Иванов А.Е.")
}
