package main

import "fmt"

// slice is a 前闭后开区间
func main() {
	s := make([]int, 10)

	// init
	i := 0
	for i < 10 {
		s[i] = i
		i++
	}

	// 1 -- 5
	fmt.Println(s[:5])

	fmt.Println(s[5])

	// 5 -- 10
	fmt.Println(s[5:6])
}
