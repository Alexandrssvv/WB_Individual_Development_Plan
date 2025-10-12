package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	var size int
	fmt.Print("Введите число элементов массива: ")
	fmt.Scan(&size)

	nums := make([]int, size)
	fmt.Printf("Введите %d элементов массива: ", size)
	for i := 0; i < size; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Printf("Факториал числа %d = %d\n", n, factorial(n))
	fmt.Printf("Максимальное число в срезе %v = %d\n", nums, max(nums))
}

func factorial(x int) int {

	if x == 0 {
		return 1
	}
	if x < 0 {
		return 0
	}

	result := 1
	for i := 1; i <= x; i++ {
		result *= i
	}
	return result
}

func max(arr []int) int {
	m := arr[0]
	for _, v := range arr {
		if v > m {
			m = v
		}
	}
	return m
}
