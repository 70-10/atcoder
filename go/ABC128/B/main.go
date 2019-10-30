package main

import (
	"fmt"
	"sort"
	"strings"
)

type Restaurant struct {
	ID    int
	Name  string
	Point int
}

type Restaurants []Restaurant

func (r Restaurants) Len() int {
	return len(r)
}

func (r Restaurants) Less(i, j int) bool {
	if r[i].Name == r[j].Name {
		return r[i].Point > r[j].Point
	}
	return r[i].Name < r[j].Name
}

func (r Restaurants) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func main() {
	var n int
	fmt.Scanln(&n)

	r := make([]Restaurant, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%s %d", &r[i].Name, &r[i].Point)
		r[i].ID = i + 1
	}

	fmt.Println(output(logic(r)))
}

func logic(r []Restaurant) []int {
	sort.Sort(Restaurants(r))
	result := []int{}

	for _, i := range r {
		result = append(result, i.ID)
	}
	return result
}

func output(nums []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nums)), "\n"), "[]")
}
