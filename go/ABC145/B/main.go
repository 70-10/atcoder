package main

import (
	"fmt"
)

func main() {
	var n int
	var s string

	fmt.Scanln(&n)
	fmt.Scanln(&s)

	fmt.Println(logic(s))
}

func logic(s string) string {
	if len(s)%2 != 0 {
		return "No"
	}

	if s[0:(len(s))/2] != s[(len(s))/2:] {
		return "No"
	}

	return "Yes"
}
