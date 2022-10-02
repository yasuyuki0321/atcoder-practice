package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := 0
	for i := 0; ; i++ {
		s += i

		if s >= n {
			fmt.Println(i)
			break
		}
	}
}
