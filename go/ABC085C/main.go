package main

import "fmt"

func main() {
	var n, y int
	fmt.Scanf("%d %d", &n, &y)

	fmt.Println(logic(n, y))
}

func logic(n, y int) (int, int, int) {
	var a, b, c int
	var x int
	for i := (y / 10000); i >= 0; i-- {
		a = i
		x = 10000 * i
		if x == y && i == n {
			return a, b, c
		}

		if x >= y {
			continue
		}

		for j := (n - i); j >= 0; j-- {
			b = j
			x = 10000*i + 5000*j

			if x == y && i+j == n {
				return a, b, c
			}

			if x >= y {
				continue
			}

			for k := (n - i - j); k >= 0; k-- {
				c = k
				x = 10000*i + 5000*j + 1000*k

				if x == y && i+j+k == n {
					return a, b, c
				}

				if x >= y {
					continue
				}
			}
		}

	}
	return -1, -1, -1
}
