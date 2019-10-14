package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n-1)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}
	for i := 0; i < n-1; i++ {
		fmt.Scan(&c[i])
	}

	fmt.Println(logic(a, b, c))
}

func logic(a, b, c []int) int {
	before := -2
	count := 0
	for _, v := range a {
		i := v - 1
		count += b[i]

		if i-before == 1 {
			count += c[before]
		}

		before = i
	}
	return count
}
