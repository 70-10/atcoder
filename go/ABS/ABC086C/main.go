package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	t, x, y := make([]int, n+1), make([]int, n+1), make([]int, n+1)
	t[0], x[0], y[0] = 0, 0, 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&t[i], &x[i], &y[i])
	}

	if logic(t, x, y) {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func logic(t, x, y []int) bool {
	for i := 1; i < len(t); i++ {
		time := t[i] - t[i-1]
		d := abs(x[i]-x[i-1]) + abs(y[i]-y[i-1])

		if time < d || time%2 != d%2 {
			return false
		}

	}
	return true
}

func abs(num int) int {
	return int(math.Abs(float64(num)))
}
