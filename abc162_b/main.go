package main

import (
	"fmt"
)

func checkFizzBuzz(i int) bool {
	if !(i%3 == 0 || i%5 == 0 || (i%3 == 0 && i%5 == 0)) {
		return true
	} else {
		return false
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	s := 0
	for i := 1; i <= n; i++ {
		if checkFizzBuzz(i) {
			s += i
		}
	}
	fmt.Println(s)

}
