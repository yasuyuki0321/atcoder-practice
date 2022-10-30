package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func funcA(in io.Reader) string {
	if in == nil {
		in = os.Stdin
	}

	s := []string{}
	for i := 0; i < 26; i++ {
		var a rune
		fmt.Fscan(in, &a)
		s = append(s, string(a+96))
	}

	ss := strings.Join(s, "")
	return ss
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println(funcA(in))
}
