package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stoi(s []string) []int {
	i := make([]int, 0)
	for _, v := range s {
		iv, _ := strconv.Atoi(v)
		i = append(i, iv)
	}
	return i
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	sc.Scan()
	p := strings.Split(sc.Text(), " ")
	ip := stoi(p)

	c := 0
	for i := 1; i < len(ip)-1; i++ {
		if (ip[i-1] < ip[i] && ip[i] < ip[i+1]) ||
			(ip[i+1] < ip[i] && ip[i] < ip[i-1]) {
			c += 1
		}
	}
	fmt.Println(c)
}
