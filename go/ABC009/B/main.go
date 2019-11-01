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
	m := make(map[int]int)
	list := []int{}
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num]++
			list = append(list, num)
		}
	}
	sort.Ints(list)
	return list[len(list)-2]
}
