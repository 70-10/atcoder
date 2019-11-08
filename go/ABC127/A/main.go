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
	if a <= 5 {
		return 0
	}

	if a <= 12 {
		return b / 2
	}

	return b
}
