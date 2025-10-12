package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Print("Введите a : ")
	fmt.Scan(&a)
	fmt.Print("Введите b : ")
	fmt.Scan(&b)
	double := makeMultiplier(a)
	fmt.Println(double(b))

}

func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
