package main

import "fmt"

var arr [5]int = [5]int{1, 2, 3, 4, 5}

var fruits = map[string]int{
	"apple":     1,
	"pear":      2,
	"peach":     3,
	"orange":    4,
	"pineapple": 5,
}

func main() {
	sum := 0

	fmt.Println("Элементы массива arr:")
	for i, v := range arr {
		fmt.Printf("Элемент %d = %d\n", i, v)
		sum += v
	}

	fmt.Println("Сумма элементов массива arr =", sum)

	for key, value := range fruits {
		fmt.Printf("%s : %d\n", key, value)
	}

	if _, ok := fruits["apple"]; ok {
		fmt.Println("Количество яблок =", fruits["apple"])
	} else {
		fmt.Println("Нет такого фрукта")
	}

}
