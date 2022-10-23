package main

import (
	"fmt"
	"strings"
)

func main() {
	var d int
	fmt.Scan(&d)

	fmt.Println("Christmas" + strings.Repeat(" Eve", 25-d))
}
