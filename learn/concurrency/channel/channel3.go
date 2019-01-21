package main

import (
	"fmt"
	"time"
)

/*
   Channel as function parameters
*/

func f1(c chan int, x int) {
	fmt.Println(x)
	c <- x
}

func f2(c chan<- int, x int) {
	fmt.Println(x)
	c <- x
}

func main() {
	c1 := make(chan int)
	go f1(c1, 12)

	c2 := make(chan<- int)
	go f2(c2, 33)

	c3 := make(chan int)
	go f2(c3, 33)

	time.Sleep(time.Second)
}
