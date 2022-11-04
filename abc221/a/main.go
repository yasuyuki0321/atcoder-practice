package main

import (
	"fmt"
)

func funcA(a, b int) int {
	ans := 1
	for i := 0; i < a-b; i++ {
		ans *= 32
	}
	return ans
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(funcA(a, b))
}
