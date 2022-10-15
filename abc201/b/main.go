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
	n, _ := strconv.Atoi(sc.Text())

	m := make(map[int]string)

	for i := 0; i < n; i++ {
		sc.Scan()
		sl := strings.Split(sc.Text(), " ")

		s := sl[0]
		t, _ := strconv.Atoi(sl[1])

		m[t] = s
	}

	h := make([]int, 0)
	for k, _ := range m {
		h = append(h, k)
	}
	sort.Ints(h)

	fmt.Println(m[h[len(h)-2]])
}
