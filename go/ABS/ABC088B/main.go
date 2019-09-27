package main

import (
	"fmt"
	"sort"
)

func main() {
	var num int
	fmt.Scanf("%d", &num)

	var n = make([]int, num)
	var s = make([]interface{}, num)
	for i := range n {
		s[i] = &n[i]
	}
	fmt.Scanln(s...)

	fmt.Println(logic(n))
}

func logic(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	var sum int
	for i, c := range nums {
		if i%2 == 0 {
			sum += c
		} else {
			sum -= c
		}
	}
	return sum
}
