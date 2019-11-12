package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scanln(&s)

	fmt.Println(logic(s))
}

func logic(s string) string {
	if s == "A" {
		return "T"
	}

	if s == "T" {
		return "A"
	}

	if s == "G" {
		return "C"
	}

	return "G"
}
