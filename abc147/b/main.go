package main

import "fmt"

func funcA(s string) int {
	c := 0
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			c += 1
		}
	}
	return c
}

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(funcA(s))
}
