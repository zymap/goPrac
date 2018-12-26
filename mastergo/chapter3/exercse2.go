package main

import "fmt"

func main() {
	s1 := make([]int , 1024)
	reslice := s1[1:3]
	fmt.Println(s1)
	fmt.Println(reslice)

	reslice[0] = 100
	reslice[1] = 123123
	fmt.Println(s1)
	fmt.Println(reslice)
}