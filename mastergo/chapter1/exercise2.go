package main

import (
	"os"
	"log"
	"strconv"
)

func main()  {
	if len(os.Args) == 1 {
		log.Fatal("input args error, please input more")
	}
	log.Println("args : ", os.Args[1:])
	var result float64

	for _, value := range os.Args[1:] {
		var tmp float64
		tmp, err := strconv.ParseFloat(value, 2)
		if err != nil {
			log.Fatal(value + " convert float error")
		}
		result += tmp
	}

	r := result / float64(len(os.Args) - 1)
	log.Println("the result is -> ", strconv.FormatFloat(r, 'f', 2, 64))
}