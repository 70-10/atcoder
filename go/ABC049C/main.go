package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanln(&s)

	if logic(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func logic(s string) bool {
	var result string
	result = s
	result = strings.Replace(result, "eraser", "", -1)
	result = strings.Replace(result, "erase", "", -1)
	result = strings.Replace(result, "dreamer", "", -1)
	result = strings.Replace(result, "dream", "", -1)

	if result == "" {
		return true
	}
	return false
}
