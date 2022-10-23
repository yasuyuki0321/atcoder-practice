package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	ps := make([]int, 0)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		ps = append(ps, a)
	}
	sort.Ints(ps)

	sum := 0
	for i := 0; i < n; i++ {
		if i != n-1 {
			sum += ps[i]
		} else {
			sum += ps[i] / 2
		}
	}
	fmt.Println(sum)
}
