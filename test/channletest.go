package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 1)

	go func() {
		i := 1
		for {
			fmt.Println("sleep", i)
			time.Sleep(time.Second)
			i++

			if i == 20 {
				break
			}
		}
		c <- 1
	}()

	m := <-c

	fmt.Println(m)
}
