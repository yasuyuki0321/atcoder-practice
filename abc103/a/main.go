package main

import (
	"fmt"
	"sort"
)

func funcA(s []int) int {
	sort.Ints(s)

	sum := 0
	sum += s[2] - s[1]
	sum += s[1] - s[0]

	return sum
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	s := []int{a, b, c}

	fmt.Println(funcA(s))
}
