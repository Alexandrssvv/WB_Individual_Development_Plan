package main

import "fmt"

func main() {
	var a int
	fmt.Print("Введите a: ")
	fmt.Scan(&a)

	length, data := sliceOps(a)
	fmt.Println("len =", length, "data =", data)
}

func sliceOps(n int) (int, []int) {
	if n < 0 {
		n = 0
	}
	nums := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	return len(copyNums), copyNums
}
