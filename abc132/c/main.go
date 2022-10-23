package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	d := make([]int, 0)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		d = append(d, a)
	}

	sort.Ints(d)
	fmt.Println((d[n/2]) - (d[n/2-1]))
}
