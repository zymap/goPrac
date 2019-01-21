package main

import "fmt"

func main() {
	r := getSum(-11, 11)
	fmt.Println(r)
}

func getSum(a int, b int) int {
	if b == 0 {
		return a
	}

	sum := a ^ b
	carry := (a & b) << 1

	return getSum(sum, carry)
}
