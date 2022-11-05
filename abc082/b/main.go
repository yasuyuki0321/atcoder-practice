package main

import (
	"fmt"
	"sort"
	"strings"
)

func input() ([]string, []string) {
	var s, t string
	fmt.Scan(&s, &t)

	ss := strings.Split(s, "")
	st := strings.Split(t, "")

	return ss, st
}

func funcA(s, t []string) string {
	sort.Strings(s)
	sort.Sort(sort.Reverse(sort.StringSlice(t)))

	a := strings.Join(s, "")
	b := strings.Join(t, "")

	ans := "No"
	if a < b {
		ans = "Yes"
	}
	return ans
}

func main() {
	fmt.Println(funcA(input()))
}
