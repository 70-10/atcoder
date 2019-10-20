package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b int

	fmt.Scanf("%d %d", &a, &b)

	fmt.Println(logic(a, b))
}

func logic(a int, b int) int {
	result := []int{a + b, a - b, a * b}
	sort.Ints(result)
	return result[len(result)-1]
}
