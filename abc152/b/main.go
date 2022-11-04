package main

import (
	"fmt"
	"strconv"
	"strings"
)

func input() (int, int) {
	var a, b int
	fmt.Scan(&a, &b)

	return a, b
}

func funcA(x, y int) string {
	ax := strconv.Itoa(x)
	ay := strconv.Itoa(y)

	var ans string
	if ax < ay {
		ans = strings.Repeat(ax, y)
	} else {
		ans = strings.Repeat(ay, x)
	}
	return ans
}

func main() {
	fmt.Println(funcA(input()))
}
