package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	var n, a, b int
	fmt.Scanf("%d %d %d", &n, &a, &b)

	fmt.Println(logic(n, a, b))
}

func logic(n, a, b int) int {
	var count int
	for i := 1; i <= n; i++ {
		val := sum(i)
		if val >= a && val <= b {
			count += i
		}
	}
	return count
}

func sum(i int) int {
	var sum int
	s := strconv.Itoa(i)

	for _, c := range s[:len(s)] {
		buf := make([]byte, 1)
		_ = utf8.EncodeRune(buf, c)
		val, _ := strconv.Atoi(string(buf))
		sum += val
	}
	return sum
}
