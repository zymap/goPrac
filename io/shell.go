package io

import (
	"flag"
	"fmt"
	"os"
)

//go run main.go -id xxx -name xxx

func Shell()  {
	id := flag.Int("id",0, "input id:")
	name := flag.String("name","default", "input name")
	flag.Parse()
	fmt.Println(*id,*name)

	args := os.Args
	fmt.Println(args)

	if args == nil {
		fmt.Println("name ?")
	}
}