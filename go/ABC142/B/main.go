package main

import "fmt"

func main() {
	var n, k int

	fmt.Scanf("%d %d", &n, &k)

	var nums = make([]int, n)
	var s = make([]interface{}, n)
	for i := range nums {
		s[i] = &nums[i]
	}
	fmt.Scanln(s...)

	fmt.Println(logic(nums, k))
}

func logic(nums []int, k int) int {
	ans := []int{}
	for _, n := range nums {
		if n >= k {
			ans = append(ans, n)
		}
	}

	return len(ans)
}
