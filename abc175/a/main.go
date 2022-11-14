package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			cnt += 1
		}
	}
	if cnt == 2 && s[1] == 'S' {
		cnt -= 1
	}

	fmt.Println(cnt)
}
