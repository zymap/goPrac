package main

import (
	"sync"
	"fmt"
	"strconv"
	"time"
	"runtime"
)

var syn = sync.Mutex{}

func main() {
	go test2()
	go test1()

	time.Sleep(4 * time.Second)
}

func test1()  {
	for i := 0; i < 10000; i++ {
		syn.Lock()
		runtime.Gosched()
		fmt.Println("test1 print : " + strconv.Itoa(i))
		syn.Unlock()
	}
}

func test2()  {
	time.Sleep(500 * time.Millisecond)
	for i := 0; i < 10000; i++ {
		syn.Lock()
		fmt.Println("test2 print: " + strconv.Itoa(i))
		syn.Unlock()
	}
}