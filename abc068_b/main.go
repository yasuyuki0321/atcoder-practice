package main

import (
	"fmt"
)

func divCount(n int) int {
	c := 0
	for {
		if n%2 == 0 {
			c++
			n = n / 2
		} else {
			break
		}
	}
	return c
}

func main() {
	var n int
	fmt.Scan(&n)

	max := 0
	ans := 0
	for i := 1; i <= n; i++ {
		c := divCount(i)

		if c >= max {
			max = c
			ans = i
		}
	}
	fmt.Println(ans)
}
