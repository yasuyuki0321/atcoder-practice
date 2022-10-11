package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]int, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		s[i] = a
	}

	sort.Sort(sort.Reverse(sort.IntSlice(s)))

	var a, b int
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			a += s[i]
		} else {
			b += s[i]
		}
	}
	fmt.Println(a - b)
}
