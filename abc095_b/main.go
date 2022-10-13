package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := strings.Split(sc.Text(), " ")

	n, _ := strconv.Atoi(s[0])
	x, _ := strconv.Atoi(s[1])

	ms := make([]int, 0)
	sum := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		m, _ := strconv.Atoi(sc.Text())
		ms = append(ms, m)
		sum += m
	}

	sort.Ints(ms)
	fmt.Println(n + (x-sum)/ms[0])
}
