package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	for i := 1; ; i++ {
		if i%2 == 1 {
			c -= b
			if c <= 0 {
				fmt.Println("Yes")
				break
			}
		} else {
			a -= d
			if a <= 0 {
				fmt.Println("No")
				break
			}
		}
	}
}
