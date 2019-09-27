package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	if logic(a, b, c) {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}

func logic(a, b, c int) bool {
	nums := []int{a, b, c}
	sort.Ints(nums)
	if equal(nums,[]int{5, 5, 7}) {
		return true
	}
	return false
}

func equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}