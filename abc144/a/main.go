package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func input(in io.Reader) (int, int) {
	if in == nil {
		in = os.Stdin
	}
	var a, b int
	fmt.Fscan(in, &a, &b)

	return a, b
}

func funcA(a, b int) int {
	ans := 0
	if a >= 10 || b >= 10 {
		ans = -1
	} else {
		ans = a * b
	}
	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println(funcA(input(in)))
}
