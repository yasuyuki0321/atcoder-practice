package main

import "fmt"

func a(n int) int {
	s := 1
	for i := 1; i <= n; i++ {
		s *= i
	}
	return s
}

func main() {
	var p int
	fmt.Scan(&p)

	ans := 0
	for i := 10; i != 0; i-- {

		s := p / a(i)
		p %= a(i)

		ans += s
	}
	fmt.Println(ans)
}
