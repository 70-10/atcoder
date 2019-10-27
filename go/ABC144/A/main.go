package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(logic(a, b))
}

func logic(a, b int) int {
	if a > 9 || b > 9 {
		return -1
	}
	return a * b
}
