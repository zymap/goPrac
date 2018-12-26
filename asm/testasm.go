package main

import "fmt"

func main() {

	fmt.Println("hello")
	var c chan int
	c <- 12
	fmt.Println("hello")
}
