package main

import (
	"runtime"
	"log"
	"time"
)

func printStates(mem runtime.MemStats)  {
	runtime.ReadMemStats(&mem)
	log.Println("mem.Alloc: ", mem.Alloc)
	log.Println("mem.TotalAlloc: ", mem.TotalAlloc)
	log.Println("mem.HeapAlloc: ", mem.HeapAlloc)
	log.Println("mem.NumGC: ", mem.NumGC)
	log.Println("===============================")
}

func main()  {
	var mem runtime.MemStats
	printStates(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 10240000)
		if s == nil {
			log.Println("Operation failed !")
		}
	}

	printStates(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 20480000)
		if s == nil {
			log.Println("Operation faild")
		}
		printStates(mem)
		time.Sleep(5 * time.Second)
	}
	printStates(mem)
}
