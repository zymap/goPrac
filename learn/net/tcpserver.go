package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	server("0.0.0.0:8080")
}

func server(address string) {
	fmt.Println("listening : ", address)
	conn, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("open server error.")
		panic(err)
	}
	defer conn.Close()

	c, err := conn.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("server reciver: ", string(netData))

		localtime := time.Now().Format(time.Kitchen) + "\n"
		c.Write([]byte(localtime))
	}
}
