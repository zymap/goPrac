package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	src := readFile("/Users/zhangyong/test/utxo.log.a")
	des := readFile("/Users/zhangyong/test/copernicus/utxo.log.back")

	cache := make(map[string]interface{})

	for {
		srcline, err := src.ReadBytes('\n')
		if err != nil {
			fmt.Println("src readline error")
		}

		desline, err := des.ReadBytes('\n')
		if err != nil {
			fmt.Println("des readline error")
		}

		if string(desline) == string(srcline) {
			continue
		}

		for k, _ := range cache {
		    if k == string()
        }
    }
}

func checkQueue(q []string, word string) bool {
	for i, v := range q {
		if v == word {
			q[i] = ""
			return true
		}
	}
	return false
}

func readFile(filename string) *bufio.Reader {

	file, err := os.Open(filename)
	if err != nil {
		panic("open " + filename + " error !")
	}
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			fmt.Println("hello")
			os.Exit(0)
		}
		fmt.Println(string(line))
		//fmt.Println(string(line))
	}

	return reader
}
