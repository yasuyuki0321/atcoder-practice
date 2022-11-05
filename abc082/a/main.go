package main

import (
	"fmt"
	"math"
)

func input() (float64, float64) {
	var a, b float64
	fmt.Scan(&a, &b)

	return a, b
}

func funcA(a, b float64) float64 {
	return math.Ceil((a + b) / 2)

}

func main() {
	fmt.Println(funcA(input()))
}
