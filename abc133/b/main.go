package main

import (
	"fmt"
	"math"
)

func main() {
	var n, d int
	fmt.Scan(&n, &d)

	m := make([][]float64, n)
	for i := 0; i < n; i++ {
		m[i] = []float64{}
		for j := 0; j < d; j++ {
			var x float64
			fmt.Scan(&x)

			m[i] = append(m[i], x)
		}
	}

	ans := 0
	for i := 0; i < len(m); i++ {
		var a float64
		for j := 0; j < len(m[i]); j++ {

			if i == len(m)-1 {
				a += math.Pow(m[len(m)-1][j]-m[0][j], 2)
			} else {
				a += math.Pow(m[i+1][j]-m[i][j], 2)
			}
		}
		if math.Floor(math.Sqrt(a)) == math.Sqrt(a) {
			ans += 1
		}
	}
	fmt.Println(ans)
}
