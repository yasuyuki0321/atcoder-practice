package main

import "fmt"

func input() (int, int, int) {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	return a, b, c
}

func funcA(a, b, c int) int {
	ans := 0
	if a == b {
		ans = c
	} else if a == c {
		ans = b
	} else {
		ans = a
	}
	return ans
}

func main() {
	fmt.Println(funcA(input()))
}
