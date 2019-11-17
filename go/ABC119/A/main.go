package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Scanln(&s)

	fmt.Println(logic(s))
}

func logic(s string) string {
	a := strings.Replace(s, "/", "", -1)
	b, _ := strconv.Atoi(a)

	if b < 20190501 {
		return "Heisei"
	}
	return "TBD"
}
