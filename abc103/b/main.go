package main

import (
	"fmt"
)

func funcA(s, t string) string {
	ans := "No"
	for i := 0; i < len(s); i++ {

		a := string(s[len(s)-i-1:]) + string(s[0:len(s)-i-1])
		if a == t {
			ans = "Yes"
			break
		}
	}
	return ans
}

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	fmt.Println(funcA(s, t))
}
