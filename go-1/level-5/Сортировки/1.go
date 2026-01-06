package main

import (
	"fmt"
	"strings"
	"slices"
	"sort"
)

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{
			Name: "Ivan",
			Age:  56,
		},
		{
			Name: "Tim",
			Age:  33,
		},
		{
			Name: "Bob",
			Age:  89,
		},
	}

	// пакет sort
	fmt.Println("Пакет sort")
	intSlice := []int{4, 5, 2, 1, 3, 9, 7, 8, 6}
	fmt.Println(sort.IntsAreSorted(intSlice)) // Проверим, отсортирован ли слайс
	sort.Ints(intSlice)	// Сама сортировка
	fmt.Println(intSlice) // [1 2 3 4 5 6 7 8 9]

	// Для сортировки объектов других типов данных, либо по какому-то другому признаку 
	// (например, строки по их длине или пользователей по их возрасту) можно использовать функцию Slice(). 
	// Она принимает функцию less в качестве аргумента:
	// func Slice(x interface{}, less func(i, j int) bool)
	// Функция less определяет, должен ли элемент с индексом i находиться перед элементом с индексом j:
	strs := []string{"a very long string", "a medium string", "a short one"}
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	fmt.Println(strings.Join(strs, ", ")) // a short one, a medium string, a very long string
	// Сортировка структур
	sort.Slice(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})
	fmt.Println(users) // [{Tim 33} {Ivan 56} {Bob 89}]




	// пакет slices
	fmt.Println("Пакет slices")
	smallInts := []int8{0, 42, -10, 8}
	slices.Sort(smallInts)
	fmt.Println(smallInts)

	// Чтобы настроить пользовательскую функцию сравнения, нужна функция SortFunc:
	// func SortFunc[S ~[]E, E any](x S, cmp func(a, b E) int)

	slices.SortFunc(smallInts, func(a, b int8) int {
        // Здесь мы сравниваем сами
        switch {
        case a < b:
            return 1
        case a > b:
            return -1
        default:
            return 0
        }
	})
	fmt.Println(smallInts) // [42 8 0 -10]

	// Сортировка структур
	slices.SortFunc(users, func(a, b User) int {
			// Здесь мы сравниваем сами
			switch {
			case a.Age < b.Age:
				return 1
			case a.Age > b.Age:
				return -1
			default:
				return 0
			}
	})
	fmt.Println(users) // [{Tim 33} {Ivan 56} {Bob 89}]
}