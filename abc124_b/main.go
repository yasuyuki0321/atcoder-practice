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
	// n, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	s := strings.Split(sc.Text(), " ")

	is := make([]int, 0)
	for _, v := range s {
		a, _ := strconv.Atoi(v)
		is = append(is, a)
	}

	saw := is[0]
	c := 0
	for _, v := range is {
		if saw <= v {
			c += 1
			if saw < v {
				saw = v
			}
		}
	}
	fmt.Println(c)
}
