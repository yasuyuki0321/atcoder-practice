package main

import (
	"fmt"
)

func main() {
	var x, s int
	fmt.Scan(&x)

	s = 100
	for i := 1; ; i++ {
		s += (s / 100)

		if s >= x {
			fmt.Println(i)
			break
		}
	}
}
