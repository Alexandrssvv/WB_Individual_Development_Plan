package main

import "fmt"

var day int

func main() {
	fmt.Print("Введите число от 1 до 7: ")
	fmt.Scan(&day)

	switch day {
	case 1:
		fmt.Printf("%d - Понедельник\n", day)
	case 2:
		fmt.Printf("%d - Вторник\n", day)
	case 3:
		fmt.Printf("%d - Среда\n", day)
	case 4:
		fmt.Printf("%d - Четверг\n", day)
	case 5:
		fmt.Printf("%d - Пятница\n", day)
	case 6:
		fmt.Printf("%d - Суббота\n", day)
		fmt.Println("Это выходной!")
	case 7:
		fmt.Printf("%d - Воскресенье\n", day)
		fmt.Println("Это выходной!")
	default:
		fmt.Println("Ошибка: нет такого дня недели")

	}
}
