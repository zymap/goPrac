package main

import (
	"fmt"
	"time"
)

/*
   This function is used to create multiple
   goroutines.
*/

func main() {
	n := 100

	fmt.Printf("Going to create %d goroutines.\n", n)

	for i := 0; i < n; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("\nExiting...")
}
