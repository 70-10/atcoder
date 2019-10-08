package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int

	fmt.Scanf("%d", &n)

	var nums = make([]int, n)
	var s = make([]interface{}, n)
	for i := range nums {
		s[i] = &nums[i]
	}
	fmt.Scanln(s...)

	fmt.Println(logic(nums))
}

func logic(nums []int) string {
	result := make([]string, len(nums))

	for i, n := range nums {
		result[n-1] = fmt.Sprintf("%d", i+1)
	}

	return strings.Join(result, " ")
}
