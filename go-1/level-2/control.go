package main

import (
	"fmt"
	"time"
	"errors"
	"strings"
	"unicode/utf8"
)

// TimeNow — внешняя функция, которую не нужно реализовывать
// Предполагается, что она уже определена где-то вне этого файла
func TimeNow() time.Time {
	return time.Now()
}

func currentDayOfTheWeek() string {
	today := TimeNow().Weekday()
	days := map[time.Weekday]string {
		time.Monday: "Понедельник",
		time.Tuesday:   "Вторник",
		time.Wednesday: "Среда",
		time.Thursday:  "Четверг",
		time.Friday:    "Пятница",
		time.Saturday:  "Суббота",
		time.Sunday:    "Воскресенье",
	}
	return days[today]
}

func dayOrNight() string {
	now := TimeNow().Hour()
	if now >= 10 && now <= 22 {
		return "День"
	}
	return "Ночь"
}

func nextFriday() int {
	now := currentDayOfTheWeek()
	days := map[string]int{
			"Понедельник": 3,
			"Вторник":     4,
			"Среда":       5,
			"Четверг":     6,
			"Пятница":     7,
			"Суббота":     1,
			"Воскресенье": 2,
		}
	return 7 - days[now]
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	today := strings.ToLower(currentDayOfTheWeek())
	return today == strings.ToLower(answer)
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) != 4 {
		return false, errors.New("исправь свой ответ, а лучше ложись поспать")
	}
	answerLower := strings.ToLower(answer)
	correctLower := strings.ToLower(dayOrNight())

	return answerLower == correctLower, nil
}

func main() {
	fmt.Println(currentDayOfTheWeek())
	fmt.Println(dayOrNight())
	fmt.Println(nextFriday())
}
