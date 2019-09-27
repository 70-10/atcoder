package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	fmt.Println(logic(a, b, c))
}

func logic(a, b, c int) int {
	if a == b {
		return c
	}

	if a == c {
		return b
	}

	return a
}
