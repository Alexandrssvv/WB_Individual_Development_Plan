package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Введите число a: ")
	fmt.Scan(&a)
	fmt.Print("Введите число b: ")
	fmt.Scan(&b)

	fmt.Printf("Старт:        a=%d, b=%d\n", a, b)

	x, y := swap(a, b)
	//swapPtr(&a, &b)
	fmt.Printf("x =%d, y = %d, a = %d, b = %d\n", x, y, a, b)
	//fmt.Printf("a = %d, b = %d\n", a, b)
}

func swap(a, b int) (int, int) {
	return b, a
}

//
//func swapPtr(a, b *int) {
//	*a, *b = *b, *a
//}
