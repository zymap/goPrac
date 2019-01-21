package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	udpclient()
}

func udpclient() {
	server := ":8080"
	s, err := net.ResolveUDPAddr("udp", server)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The udp server is %s", c.RemoteAddr().String())
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("<< ")
		text, _ := reader.ReadString('\n')
		data := []byte(text + "\n")
		_, err = c.Write(data)
		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("exit udp client")
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Reply: %s\n", string(buffer[0:n]))
	}
}
