package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var in = bufio.NewReader(os.Stdin)

func input(in io.Reader) int {
	var n int
	fmt.Fscan(in, &n)

	return n
}

func funcA(n int) string {
	ans := "No"
	for i := 0; i < 100/4; i++ {
		for j := 0; j < 100/7; j++ {
			if 4*i+7*j == n {
				ans = "Yes"
				break
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(funcA(input(in)))
}
