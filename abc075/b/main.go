package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	var s = make([][]string, h)
	for i := 0; i < h; i++ {
		var a string
		fmt.Scan(&a)

		for j := 0; j < w; j++ {
			s[i] = append(s[i], string(a[j]))
		}
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			funcA(s, i, j)
		}
	}
}

func funcA(s [][]string, i, j int) {
	x := []int{0, 0, -1, -1, -1, 1, 1, 1}
	y := []int{1, -1, 0, 1, -1, 0, 1, -1}

	count := 0
	for i := 0; i < len(x); i++ {
		a := i - x[i]
		b := j - y[i]
		if a < 0 || b < 0 || a >= len(s) || b >= len(s[i]) {
			continue
		} else {
			// fmt.Print(s[a][b])
			if s[a][b] == "#" {
				count += 1
			}
		}
	}
	fmt.Println(count)
}
