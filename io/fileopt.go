package io

import (
	"os"
	"fmt"
)

var FILE = "/home/zy/myuser/go/goproject/goPrac/file/iotest.txt"

func IoOpt()  {
	writeFile()
	readFile()
}

func writeFile()  {
	outputstream,err := os.Open(FILE)
	defer outputstream.Close()
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(err," ","creating")
			outputstream,err = os.Create(FILE)
			if err != nil {
				fmt.Println("create file "+FILE+" failed...")
			}
		}else {
			fmt.Println(FILE," open error: ", err)
			return
		}
	}
	fmt.Println("writing...")
	outputstream.WriteString("hello this is my first write\n")
	outputstream.Write([]byte("the bytes write in\n"))
	outputstream.WriteString("lalalal")
	fmt.Println("writing over....")
}

func readFile()  {
	fmt.Println("reading.....")
	inputstream,err := os.Open(FILE)
	defer inputstream.Close()
	if err != nil {
		fmt.Println("file error: ", err)
		return
	}
	buffer := []byte{0,0,0,0,0,0,0,0}
	for {
		n,_ := inputstream.Read(buffer)
		if n == 0 {
			break
		}
		fmt.Print(string(buffer))
	}
}