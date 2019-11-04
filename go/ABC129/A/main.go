package main

import (
	"fmt"
	"sort"
)

func main() {
	var p, q, r int
	fmt.Scanf("%d %d %d", &p, &q, &r)

	fmt.Println(logic(p, q, r))
}

func logic(p, q, r int) int {
	nums := []int{p, q, r}
	sort.Ints(nums)

	return nums[0] + nums[1]
}
