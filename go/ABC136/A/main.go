package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	fmt.Scanf("%d %d %d", &a, &b, &c)

	fmt.Println(logic(a, b, c))
}

func logic(a, b, c int) int {
	result := b + c - a
	if result < 0 {
		return 0
	}
	return result
}
