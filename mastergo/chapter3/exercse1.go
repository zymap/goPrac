package main

import "fmt"

func main()  {
	slice := make([]int, 1023)
	fmt.Println(cap(slice))
	for i := 0; i < 1020; i++ {
		slice = append(slice, i)
	}
	fmt.Println(cap(slice))
	fmt.Println(len(slice))
}
