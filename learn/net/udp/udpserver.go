package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {
	udpserver()
}

func random(min, max int) int {
	return rand.Intn(max-min) + max
}

func udpserver() {
	s, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp", s)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close()

	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Print("-> ", string(buffer[0:n-1]))
		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("EXITING SERVER")
			return
		}

		data := []byte(strconv.Itoa(random(1, 1001)))
		fmt.Printf("data: %s\n", string(data))
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
