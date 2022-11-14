package main

import "fmt"

func input() int {
	var n int
	fmt.Scan(&n)

	return n
}

func check(n int) int {

	ans := 0
	if 1 <= n && n <= 125 {
		ans = 4
	} else if 126 <= n && n <= 211 {
		ans = 6
	} else {
		ans = 8
	}
	return ans
}

func main() {
	fmt.Println(check(input()))
}
