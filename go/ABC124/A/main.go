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
	var x int
	for i := 0; i < 2; i++ {
		if a > b {
			x += a
			a--
		} else {
			x += b
			b--
		}
	}
	return x
}
