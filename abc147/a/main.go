package main

import "fmt"

func funcA(a1, a2, a3 int) string {
	var ans string

	if a1+a2+a3 >= 22 {
		ans = "bust"
	} else {
		ans = "win"
	}
	return ans
}

func main() {
	var a1, a2, a3 int
	fmt.Scan(&a1, &a2, &a3)

	fmt.Println(funcA(a1, a2, a3))
}
