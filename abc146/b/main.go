package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func funcA(s string, n int) string {
	ans := []string{}
	for i := 0; i < len(s); i++ {
		if s[i]+byte(n) > 90 {
			ans = append(ans, string(s[i]+byte(n)-26))
		} else {
			ans = append(ans, string(s[i]+byte(n)))
		}
	}
	return strings.Join(ans, "")
}

func main() {
	var n int
	var s string

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &s)

	fmt.Println(funcA(s, n))
}
