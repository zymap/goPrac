package io

import (
	"os"
	"fmt"
	"bufio"
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

func WriteFile1() {
	file := "/home/zy/myuser/go/goproject/goPrac/file/learn_io_write"
	outputFile, err := os.OpenFile(file, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open error")
		return 
	}
	defer outputFile.Close()
	
	outputWriter := bufio.NewWriter(outputFile)
	outputWriter.WriteString("hello my first write.")
	outputWriter.Flush()
}

func CopyFile()  {

}