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
	m := make(map[string]int)
	for _, c := range s {
		m[string(c)]++
	}

	if len(m) != 2 {
		return "No"
	}

	for _, i := range m {
		if i != 2 {
			return "No"
		}
	}
	return "Yes"
}
