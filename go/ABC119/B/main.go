package main

import (
	"fmt"
)

type Otoshidama struct {
	Value float64
	Unit  string
}

func main() {
	var n int
	fmt.Scanln(&n)

	o := make([]Otoshidama, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%f %s", &o[i].Value, &o[i].Unit)
	}

	fmt.Println(logic(o))
}

func logic(o []Otoshidama) float64 {
	var sum float64

	for _, otoshidama := range o {
		if otoshidama.Unit == "BTC" {
			sum += 380000.0 * otoshidama.Value
		} else {
			sum += otoshidama.Value
		}
	}

	return sum
}
