package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scanf("%d %d", &a, &b)

	result, err := logic(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func logic(a, b int) (int, error) {
	if a%2 != b%2 {
		return 0, fmt.Errorf("IMPOSSIBLE")
	}

	return (a + b) / 2, nil
}
