package main

import (
	"fmt"
	"math"
)

func main() {
	r := divide(-2147483648, -1)
	fmt.Println(r)
}

func divide(dividend int, divisor int) int {
	r := dividend / divisor
	if r > math.MaxInt32 {
		return math.MaxInt32
	}
	return r
}
