package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)

	fmt.Println(logic(s))
}

func logic(s string) string {
	count := make(map[string]int)

	for _, word := range s {
		count[string(word)]++
	}

	if (count["N"] > 0 && count["S"] <= 0) || (count["N"] <= 0 && count["S"] > 0) {
		return "No"
	}
	if (count["E"] > 0 && count["W"] <= 0) || (count["E"] <= 0 && count["W"] > 0) {
		return "No"
	}

	return "Yes"
}
