package main

import "fmt"

func input() int {
	var n int
	fmt.Scan(&n)

	return n
}

func funcA(n int) string {
	ans := "No"
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i*j == n {
				ans = "Yes"
				break
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(funcA(input()))
}
