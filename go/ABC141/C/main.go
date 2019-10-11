package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, k, q int
	var answers []int
	fmt.Scanf("%d %d %d", &n, &k, &q)

	for i := 0; i < q; i++ {
		var a int
		fmt.Scanln(&a)
		answers = append(answers, a)
	}

	fmt.Println(strings.Join(logic(n, k, answers), "\n"))

}

func logic(n, k int, answers []int) []string {

	limit := len(answers) - k
	var result []string

	for i := 1; i <= n; i++ {
		num := 0
		for _, a := range answers {
			if i == a {
				num++
			}
		}

		if num > limit {
			result = append(result, "Yes")
		} else {
			result = append(result, "No")
		}
	}

	return result
}
