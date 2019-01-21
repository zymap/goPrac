package main

import (
	"fmt"
	"sync"
)

/*
   Sync.Add using more than Sync.Done
*/

func main() {
	count := 20

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	waitGroup.Wait()
	fmt.Println("\nExiting...R")
}
