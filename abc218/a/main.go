package main

import "fmt"

func funcA(s string, n int) string {
	a := "Yes"
	if s[n-1] != 'o' {
		a = "No"
	}
	return a
}

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	fmt.Println(funcA(s, n))
}
