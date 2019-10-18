package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Println(logic(a))
}

func logic(nums []int) float64 {
	a := 0.0
	for _, num := range nums {
		a += 1.0 / float64(num)
	}

	return 1.0 / a
}
