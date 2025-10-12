package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	s := isEven(n)
	fmt.Println(s)
}

func isEven(n int) bool {
	return n%2 == 0
}
