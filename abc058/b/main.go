package main

import (
	"fmt"
	"strings"
)

func funcA(o, e string) string {
	s := []string{}
	for i := 0; i < len(o); i++ {
		s = append(s, string(o[i]))

		if i == len(e) {
			break
		}
		s = append(s, string(e[i]))
	}

	return strings.Join(s, "")
}

func main() {
	var o, e string
	fmt.Scan(&o, &e)

	fmt.Println(funcA(o, e))
}
