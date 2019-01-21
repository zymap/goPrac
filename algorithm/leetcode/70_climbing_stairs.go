package main

import "fmt"

func main() {
	r := climbStairs(2)
	fmt.Print(r)
}

var m = make(map[int]int)

func climbStairs(n int) int {
	if v, ok := m[n]; ok {
		return v
	}
	if n == 1 {
		return 1
	}

	if n == 2 {
		m[n] = n
		return 2
	}

	v := climbStairs(n-1) + climbStairs(n-2)
	m[n] = v
	return v
}
