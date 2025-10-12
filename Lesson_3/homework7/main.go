package main

import "fmt"

func main() {
	var a int
	fmt.Print("Введите число a: ")
	fmt.Scan(&a)

	fib := fibonacci(a)
	fmt.Println(fib)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
