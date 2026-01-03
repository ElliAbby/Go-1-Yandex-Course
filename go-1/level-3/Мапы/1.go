package main

import "fmt"

func containsDuplicate(nums []int) bool {
    set := make(map[int]bool)
    for _, num := range nums {
        if _, ok := set[num]; ok { // Если такой ключ уже есть, переходим к return
            return true
        }
        set[num] = true
    }
    return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 1, 2, 3, 4})) // содержит дубликаты -> true
	fmt.Println(containsDuplicate([]int{1, 0, 2, 3, 4})) // не содержит дубликаты -> false

	// удаление из мапы с помощью delete()
	var marks = map[string]int{"Arseniy": 4, "Polina": 5}

	fmt.Println(marks) // map[Arseniy:4 Polina:5]
	delete(marks, "Arseniy")
	fmt.Println(marks) // map[Polina:5]
}
