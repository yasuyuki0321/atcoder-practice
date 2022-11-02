package main

import (
	"fmt"
	"sort"
)

func funcA(a, b, c int) int {
	s := []int{a, b, c}
	sort.Ints(s)
	return s[0] + s[1]
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(funcA(a, b, c))
}
