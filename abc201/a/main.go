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

	a := strings.Split(sc.Text(), " ")

	s := make([]int, 0)
	for _, v := range a {
		n, _ := strconv.Atoi(v)
		s = append(s, n)
	}

	sort.Ints(s)
	sa := s[1] - s[0]
	f := true

	for i := 1; i < len(s); i++ {
		if s[i]-s[i-1] != sa {
			f = false
		}
	}

	if f {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
