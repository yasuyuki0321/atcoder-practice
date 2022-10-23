package main

import "fmt"

func main() {
	var H, W, h, w int
	fmt.Scan(&H, &W)
	fmt.Scan(&h, &w)

	s := make([][]int, H)
	for i := 0; i < H; i++ {
		s[i] = make([]int, W)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < W; j++ {
			s[i][j] = 1
		}
	}

	for i := 0; i < w; i++ {
		for j := 0; j < H; j++ {
			s[j][i] = 1
		}
	}

	c := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if s[i][j] == 0 {
				c += 1
			}
		}
	}
	fmt.Println(c)
}
