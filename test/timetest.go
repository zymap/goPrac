package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	go test()

	time.Sleep(time.Minute * 2)
}

func test() {
	for {
		fmt.Printf("it reduce 1 second.")
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("hello time is up...===================")
			//default:
			//	fmt.Println("hello it reduce 1 second.")
		}

		go func() {
			i := 0
			for {
				time.Sleep(time.Second)
				fmt.Printf("it's %d time.\n", i)
				i++

				if i == 20 {
					os.Exit(0)
				}
			}
		}()
		time.Sleep(time.Second)
	}
}
