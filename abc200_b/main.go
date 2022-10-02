package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	for i := 0; i < k; i++ {
		if n%200 == 0 {
			n = n / 200
		} else {
			s := strconv.Itoa(n) + "200"
			n, _ = strconv.Atoi(s)
		}
	}
	fmt.Println(n)
}
