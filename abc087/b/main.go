package main

import "fmt"

func input() (int, int, int, int) {
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)

	return a, b, c, x
}

func funcA(a, b, c, x int) int {
	count := 0

	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if (500*i)+(100*j)+(50*k) == x {
					count += 1
				}
			}
		}
	}
	return count
}

func main() {
	fmt.Println(funcA(input()))
}
