package main

import "fmt"

func funcA(a, b, c int) bool {
	return b-a == c-b
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	ans := "YES"
	if !funcA(a, b, c) {
		ans = "NO"
	}
	fmt.Println(ans)
}
