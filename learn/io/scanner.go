package io

import (
	"fmt"
	"bufio"
	"os"
)

func Scanner1()  {
	fmt.Println("please input a value : ")
	a := 1
	fmt.Scanln(&a)
	fmt.Println("your input value : ", a)
}

var inputReader *bufio.Reader
var input string
var err error

func Scanner2()  {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("please input something : ")
	input, err = inputReader.ReadString('e')
	if err == nil {
		fmt.Println("your input : ", input)
	}
}

