package main

import (
	"fmt"
	"sort"
)

func main() {
	var k int
	n := make([]int, 5)

	for i := 0; i < 5; i++ {
		fmt.Scanln(&n[i])
	}

	fmt.Scanln(&k)

	fmt.Println(logic(n, k))
}

func logic(n []int, k int) string {
	sort.Ints(n)

	if n[len(n)-1]-n[0] <= k {
		return "Yay!"
	}
	return ":("
}
