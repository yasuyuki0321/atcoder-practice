package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var in = bufio.NewReader(os.Stdin)

func input(in io.Reader) (int, int, int) {
	if in == nil {
		in = os.Stdin
	}

	var n, a, b int
	fmt.Fscan(in, &n, &a, &b)

	return n, a, b
}

func funcA(n, a, b int) int {
	d := n * a
	t := b

	ans := 0
	if d < t {
		ans = d
	} else {
		ans = t
	}

	return ans
}

func main() {
	fmt.Println(funcA(input(in)))
}
