package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 10進数に変換
func convert(k, n int) int {
	s := 0
	for i, v := range reverse(strconv.Itoa(n)) {
		iv, _ := strconv.Atoi(string(v))

		s += int(math.Pow(float64(k), float64(i))) * iv
	}
	return s
}

func main() {
	var a, b, k int
	fmt.Scan(&k)
	fmt.Scan(&a, &b)

	fmt.Println(convert(k, a) * convert(k, b))
}
