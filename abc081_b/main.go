package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var s = make([]int, n)

	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		s[i] = a
	}

	c := 0
	b := false
	for {
		for k, v := range s {
			if v%2 != 0 {
				b = true
			}
			s[k] = v / 2
		}
		if b == true {
			break
		}
		c++
	}
	fmt.Println(c)
}
