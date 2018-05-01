package io

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"io/ioutil"
)

var filename = "/home/zy/myuser/go/goproject/goPrac/file/learn_io.txt"
var readfile = "/home/zy/myuser/go/goproject/goPrac/file/learn_io_read"
var writefile = "/home/zy/myuser/go/goproject/goPrac/file/learn_io_write"

func InputFile()  {
	inputfile, inputError := os.Open(filename)
	if inputError != nil {
		fmt.Println("open file error.")
		return
	}
	defer inputfile.Close()

	inputReader := bufio.NewReader(inputfile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("the input was: %s\n", inputString)
		if readerError == io.EOF {
			return
		}
	}
}

func InputFile2()  {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("file error")
		return
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(writefile,buf,0644)
	if err != nil {
		panic(err.Error())
	}
}

func InputFile3()  {
	buffer := make([]byte, 128)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("open error")
		return
	}
	defer file.Close()

	inputReader = bufio.NewReader(file)
	for {
		n, err := inputReader.Read(buffer)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Println("read error")
			return
		}
		fmt.Println("read bytes : ", n, " content : ", string(buffer))
	}
}

var io4 = "/home/zy/myuser/go/goproject/goPrac/file/learn_io_4.txt"

func InputFile4()  {
	file, err := os.Open(io4)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}

