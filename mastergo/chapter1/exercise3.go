package main

import (
	"os"
	"bufio"
	"log"
	"strconv"
)

func main()  {
	f := os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	log.Println("please input int number: ")
	for scanner.Scan() {
		read := scanner.Text()
		if read == "STOP" {
			log.Fatal("EXIT....")
		}
		if value, err := strconv.Atoi(read); err == nil {
			log.Println("you input : ", value)
			continue
		}
		log.Println("input error, please input int")
	}
}