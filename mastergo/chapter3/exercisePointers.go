package main

import "fmt"

func main() {
	i := -10
	j := 25
	pI := &i
	pJ := &j

	fmt.Println(pI)
	fmt.Println(pJ)
	fmt.Println(*pI)
	fmt.Println(*pJ)

	*pI = 123456
	*pI--
	fmt.Println(i)
}

func getPointer(n *int) {
	*n = *n * *n
}

func returnPointer(n int) *int {
	v := n * n
	return &v
}
