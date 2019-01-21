package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	server1()
}

func server1() {
	server := "localhost:8080"

	s, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	l, err := net.ListenTCP("tcp", s)
	if err != nil {
		fmt.Println(err)
		return

	}

	buffer := make([]byte, 1024)
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("EXITING")
			conn.Close()
			return
		}

		fmt.Println("> ", string(buffer[0:n-1]))

		_, err = conn.Write(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
