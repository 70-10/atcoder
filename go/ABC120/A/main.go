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
	if b/a < c {
		return b / a
	}
	return c
}
