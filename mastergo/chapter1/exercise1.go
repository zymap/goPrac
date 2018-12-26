package main

import (
	"os"
	"log"
	"strconv"
)

func main()  {
	if len(os.Args) == 1 {
		log.Fatal("error args")
	}
	log.Println("args : ", os.Args[1:])
	var result int

	for i, value := range os.Args {
		if i == 0 {
			continue
		}
		num, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(value+" value convert error")
		}
		result += num
	}
	log.Println("resutl is -> ", result)
}