package main

import "fmt"

func main() {
	var num int
	var nums []int

	fmt.Scanln(&num)
	for i := 0; i < num; i++ {
		var n int
		fmt.Scanln(&n)
		nums = append(nums, n)
	}

	fmt.Println(logic(nums))
}

func logic(nums []int) int {
	return len(unique(nums))
}

func unique(nums []int) []int {
	uniq := []int{}
	m := make(map[int]bool)

	for _, n := range nums {
		if !m[n] {
			m[n] = true
			uniq = append(uniq, n)
		}
	}
	return uniq
}
