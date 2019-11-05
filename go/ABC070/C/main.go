package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&a[i])
	}

	fmt.Println(logic(a))
}

func logic(nums []int) int {
	result := 1

	for _, n := range nums {
		result = lcm(result, n)
	}

	return result
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
