package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	fmt.Println(logic(a, b))
}

func logic(a, b int) int {
	s := 1
	num := 0

	for s < b {
		s--
		s += a
		num++
	}
	return num
}
