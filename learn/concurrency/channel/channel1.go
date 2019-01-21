package main

import (
	"fmt"
	"time"
)

/*
   Writing to a channel.
*/

func writeToChannel(c chan int, x int) {
	fmt.Println(x)
	c <- x // This step will block all function, because the channel need someone to receive.
	close(c)
	fmt.Println(x) // The print will never print.
}

func main() {
	c := make(chan int)

	go writeToChannel(c, 10)
	time.Sleep(1 * time.Second)
}
