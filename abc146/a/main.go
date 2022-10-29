package main

import "fmt"

func funcA(s string) int {
	m := map[string]int{"SUN": 7, "MON": 6, "TUE": 5, "WED": 4, "THU": 3, "FRI": 2, "SAT": 1}
	ans := (m[s])
	return ans
}

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(funcA(s))
}
