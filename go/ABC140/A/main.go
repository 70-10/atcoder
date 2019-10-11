package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanln(&n)

	fmt.Println(logic(n))
}

func logic(n int) int {
	return int(math.Pow(float64(n), 3.0))
}
