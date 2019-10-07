package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	fmt.Println(logic(n))
}

func logic(n int) float64 {
	numerator := math.Ceil(float64(n) / 2)

	return numerator / float64(n)
}
