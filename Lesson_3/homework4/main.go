package main

import "fmt"

func main() {
	var b, h int

	fmt.Print("Введите ширину прямоугольника b = ")
	fmt.Scan(&b)

	fmt.Print("Введите высоту прямоугольника h = ")
	fmt.Scan(&h)

	area, perimeter := rectangle(b, h)
	fmt.Printf("Площадь прямоугольника равна %d\nПериметр прямоугольника равен %d\n", area, perimeter)

}

func rectangle(width, height int) (area, perimeter int) {
	if width <= 0 || height <= 0 {
		panic("Стороны прямоугольника не могут быть отрицательными")
	}
	area = width * height
	perimeter = 2 * (width + height)
	return
}
