package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b int
	var s string

	fmt.Scan(&a, &b, &s)
	ss := strings.Split(s, "-")

	if len(ss[0]) == a && len(ss[1]) == b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
