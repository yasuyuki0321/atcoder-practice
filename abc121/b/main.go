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
	var n, m, c int
	fmt.Scan(&n, &m, &c)

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	b := strings.Split(sc.Text(), " ")
	bi := stoi(b)

	sum := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		a := strings.Split(sc.Text(), " ")
		ai := stoi(a)

		d := 0
		for j := 0; j < m; j++ {
			d += ai[j] * bi[j]
		}

		if d+c > 0 {
			sum += 1
		}
	}
	fmt.Println(sum)
}
