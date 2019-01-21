package main

import (
	"fmt"
	"sync"
)

/*
   Sync.Add using less than Sync.Done
*/

func main() {
	count := 20

	var waitGroup sync.WaitGroup

	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}
	waitGroup.Done()

	waitGroup.Wait()
	fmt.Println("\nExiting...R")
}
