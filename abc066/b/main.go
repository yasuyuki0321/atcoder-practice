package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	max := 0
	for i := 0; i <= len(s)-1; i++ {
		if s[:i/2] == s[i/2:i/2+i/2] {
			if max < len(s[:i/2])+len(s[i/2:i/2+i/2]) {
				max = len(s[:i/2]) + len(s[i/2:i/2+i/2])
			}
		}
	}
	fmt.Println(max)
}
