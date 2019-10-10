package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)

	if logic(s) {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}

func logic(s string) bool {
	for i, c := range s {
		str := string(c)
		if i%2 != 0 && !(str == "L" || str == "U" || str == "D") {
			return false
		}
		if i%2 == 0 && !(str == "R" || str == "U" || str == "D") {
			return false
		}
	}
	return true
}
