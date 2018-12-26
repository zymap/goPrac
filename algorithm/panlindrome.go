package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(1111111111111111221111111111111111))
}

func isPalindrome(x int) bool {
	a := x
	var r int

	for x > r {
		r = 10*r + x%10
		x = x / 10
	}
	fmt.Println(x)
	return a == r
}
