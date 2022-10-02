package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for i := 1; i <= 10000; i++ {
		if i*8/100 == a && i*10/100 == b {
			fmt.Println(i)
			os.Exit(0)
		}
	}
	fmt.Println("-1")
}
