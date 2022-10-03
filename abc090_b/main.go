package main

import (
	"fmt"
	"strconv"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	c := 0
	for i := a; i <= b; i++ {
		s := strconv.Itoa(i)

		if s == reverse(s) {
			c++
		}
	}
	fmt.Println(c)
}
