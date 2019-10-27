package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Scanln(&a)
	fmt.Println(logic(a))
}

func logic(a int) string {
	for i := 1; i <= 9; i++ {
		if a/i < 10 && a%i == 0 {
			return "Yes"
		}
	}

	return "No"
}
