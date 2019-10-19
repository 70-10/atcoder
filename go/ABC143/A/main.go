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
	result := a - b*2
	if result < 0 {
		return 0
	}
	return result
}
