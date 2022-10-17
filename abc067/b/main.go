package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func stoi(s []string) []int {
	is := make([]int, 0)
	for _, v := range s {
		i, _ := strconv.Atoi(v)
		is = append(is, i)
	}
	return is
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(strings.Split(sc.Text(), " ")[0])
	k, _ := strconv.Atoi(strings.Split(sc.Text(), " ")[1])

	sc.Scan()
	l := strings.Split(sc.Text(), " ")
	is := stoi(l)

	sort.Ints(is)

	sum := 0
	for i := n - k; i < len(is); i++ {
		sum += is[i]
	}

	fmt.Println(sum)
}
