package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	as := strings.Split(sc.Text(), " ")

	ns := make([]int, 0)
	for i := 1; i <= n; i++ {
		ns = append(ns, i)
	}

	is := make([]int, 0)
	for i := 0; i < len(ns); i++ {
		a, _ := strconv.Atoi(as[i])
		is = append(is, a)
	}
	sort.Ints(is)

	if reflect.DeepEqual(ns, is) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
