package main

import "fmt"

func input() (int, int) {
	var s, t int
	fmt.Scan(&s, &t)

	return s, t
}

func check(s, t int) int {

	cnt := 0
	for a := 0; a <= s; a++ {
		for b := 0; b <= s; b++ {
			for c := 0; c <= s; c++ {
				if a+b+c <= s && a*b*c <= t {
					cnt += 1
				}
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(check(input()))
}
