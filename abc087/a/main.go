package main

import "fmt"

func input() (int, int, int) {
	var x, a, b int
	fmt.Scan(&x, &a, &b)

	return x, a, b
}

func funcA(x, a, b int) int {
	return (x - a) % b
}

func main() {
	fmt.Println(funcA(input()))
}
