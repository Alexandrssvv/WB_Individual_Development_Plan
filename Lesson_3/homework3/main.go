package main

import "fmt"

func main() {
	var size int

	fmt.Print("Введите число элементов массива: ")
	fmt.Scan(&size)

	nums := make([]int, size)
	fmt.Printf("Введите %d элементов массива: ", size)
	for i := 0; i < size; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println("В массиве: ", nums)

	min, max := minMax(nums)
	fmt.Printf("The minimum number is %d\nThe maximum number is %d\n", min, max)
}

func minMax(arr []int) (min, max int) {
	if len(arr) == 0 {
		return 0, 0
	}
	min, max = arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}
