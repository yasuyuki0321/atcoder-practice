package main

import "fmt"

func input() (int, int) {
	var n, m int
	fmt.Scan(&n, &m)

	return n, m
}

func funcA(n, m int) string {
	ans := "No"
	if n == m {
		ans = "Yes"
	}
	return ans
}

func main() {
	fmt.Println(funcA(input()))
}
