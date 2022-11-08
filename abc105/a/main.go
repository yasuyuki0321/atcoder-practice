package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var in = bufio.NewReader(os.Stdin)

func input(in io.Reader) (int, int) {
	if in == nil {
		in = os.Stdin
	}
	var n, k int
	fmt.Fscan(in, &n, &k)

	return n, k
}

func funcA(n, k int) int {
	ans := 0
	if n%k != 0 {
		ans = 1
	}
	return ans
}

func main() {
	fmt.Println(funcA(input(in)))
}
