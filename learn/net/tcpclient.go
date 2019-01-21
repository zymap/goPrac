package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	client("localhost:8080")
}

func client(out string) {
	fmt.Println("connect : ", out)
	c, err := net.Dial("tcp", out)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		if text == "EXIT" {
			return
		}
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Println("server said : ", message)
	}
}
