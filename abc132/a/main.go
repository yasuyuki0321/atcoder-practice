package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var s string

	fmt.Scan(&s)
	sl := strings.Split(s, "")

	sort.Strings(sl)

	if sl[0] == sl[1] && sl[2] == sl[3] && sl[0] != sl[2] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
