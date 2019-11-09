package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, k int
	var s string

	fmt.Scanf("%d %d\n%s", &n, &k, &s)
	fmt.Println(logic(k, s))
}

func logic(k int, s string) string {
	var result string
	for i, c := range s {
		if i == k-1 {
			result += strings.ToLower(string(c))
		} else {
			result += string(c)
		}
	}

	return result
}
