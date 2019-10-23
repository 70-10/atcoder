package main

import (
	"fmt"
	"sort"
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
	sort.Ints(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%10 != 0 {
		return sum
	}

	for _, num := range nums {
		if num%10 != 0 {
			sum -= num
			if sum%10 != 0 {
				return sum
			}
		}
	}

	return 0
}
