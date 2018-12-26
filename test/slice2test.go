package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 3, 4, 5}
	fmt.Println(a)
	replace(a)
	fmt.Println(a)

	a = []int{1, 3, 4, 5}
	var tmp []int
	fmt.Println(a)
	replace(append(tmp, a...))
	fmt.Println(a)

	a = []int{1, 3, 4, 5}
	var temp []int
	copy(temp, a)
	fmt.Println(temp)
	fmt.Println(a)
	replace(temp)
	fmt.Println(a)
}

func replace(s []int) {
	s[0] = 100
}
