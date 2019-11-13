package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)
	cities := make([]City, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%s %d", &cities[i].Name, &cities[i].Population)
	}

	fmt.Println(logic(cities))
}

func logic(cities []City) string {
	var sum int
	for _, city := range cities {
		sum += city.Population
	}

	for _, city := range cities {
		if sum/2 < city.Population {
			return city.Name
		}
	}

	return "atcoder"
}

type City struct {
	Name       string
	Population int
}
