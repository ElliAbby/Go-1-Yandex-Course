package main

import "fmt"

func main() {
	var flag_1, flag_2, flag_3, flag_4, flag_5 int
	var name_1, name_2, name_3, name_4, name_5 string = "-", "-", "-", "-", "-"

	for {
		var name string
		var place int
		fmt.Scanln(&name, &place)


		if name == "" && place == 0 {
			continue
		}

		if (name == "очередь" || name == "конец") && place == 0 {
			fmt.Printf("%d. %s\n", 1, name_1)
			fmt.Printf("%d. %s\n", 2, name_2)
			fmt.Printf("%d. %s\n", 3, name_3)
			fmt.Printf("%d. %s\n", 4, name_4)
			fmt.Printf("%d. %s\n", 5, name_5)
			if name == "конец" {
				break
			}
		}

		if name == "количество" && place == 0 {
			freePlace := 5 - flag_1 - flag_2 - flag_3 - flag_4 - flag_5
			fmt.Printf("Осталось свободных мест: %d\n", freePlace)
			fmt.Printf("Всего человек в очереди: %d\n", 5 - freePlace)
		}

		if place > 5 || place < 0{
			fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", place)
		} else if place != 0 && flag_1 + flag_2 + flag_3 + flag_4 + flag_5 == 5 {
			fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", place)
		} else if name != "количество" && name != "конец" && name != "очередь" {
			switch place {
				case 1:
					if flag_1 == 0 {
						flag_1 = 1
						name_1 = name
					} else {
						fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)
					}
				case 2:
					if flag_2 == 0 {
						flag_2 = 1
						name_2 = name
					} else {
						fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)
					}
				case 3:
					if flag_3 == 0 {
						flag_3 = 1
						name_3 = name
					} else {
						fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)
					}
				case 4:
					if flag_4 == 0 {
						flag_4 = 1
						name_4 = name
					} else {
						fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)
					}
				case 5:
					if flag_5 == 0 {
						flag_5 = 1
						name_5 = name
					} else {
						fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)
					}
				default:
					fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", place)
			}
		}
	}
}
