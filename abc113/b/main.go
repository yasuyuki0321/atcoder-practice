package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func conv2Float(s string) float64 {
	i, _ := strconv.Atoi(s)
	return float64(i)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	// n := conv2Int(sc.Text())

	sc.Scan()
	is := strings.Split(sc.Text(), " ")
	t := conv2Float(is[0])
	a := conv2Float(is[1])

	sc.Scan()
	hs := strings.Split(sc.Text(), " ")

	min := 1.7976931348623157e+308
	index := 0
	for k, v := range hs {
		var d float64
		fv := conv2Float(v)
		d = math.Abs(a - (t - fv*0.006))
		if d < min {
			min = d
			index = k
		}
	}
	fmt.Println(index + 1)
}
