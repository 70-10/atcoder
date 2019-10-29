package main

import (
	"fmt"
)

func main() {
	var n, a, b int

	fmt.Scanf("%d %d %d", &n, &a, &b)
	fmt.Println(logic(n, a, b))
}

func logic(n, a, b int) int {
	if n*a <= b {
		return n * a
	}
	return b
}
