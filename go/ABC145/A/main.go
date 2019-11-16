package main

import (
	"fmt"
)

func main() {
	var r int
	fmt.Scanln(&r)

	fmt.Println(logic(r))
}

func logic(r int) int {
	return r * r
}
