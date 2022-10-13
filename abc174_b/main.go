package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := strings.Split(sc.Text(), " ")

	n, _ := strconv.Atoi(s[0])
	d, _ := strconv.Atoi(s[1])

	c := 0
	for i := 0; i < n; i++ {

		sc.Scan()
		s := strings.Split(sc.Text(), " ")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])

		if x*x+y*y <= d*d {
			c += 1
		}
	}
	fmt.Println(c)
}
