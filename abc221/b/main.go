package main

import (
	"fmt"
	"sort"
	"strings"
)

func funcA(s, t string) bool {
	ans := false
	ss := strings.Split(s, "")

	if s == t {
		ans = true
		return ans
	}

	for i := 0; i < len(s)-1; i++ {
		sort.StringSlice.Swap(ss, i, i+1)
		s := strings.Join(ss, "")

		if s == t {
			ans = true
			break
		}
		sort.StringSlice.Swap(ss, i+1, i)
	}
	return ans
}

func input() (string, string) {
	var s, t string
	fmt.Scan(&s, &t)

	return s, t
}

func main() {
	if funcA(input()) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
