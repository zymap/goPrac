package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
	"os"
)

func main() {
	starttraceserver()
}

func starttraceserver() {
	client := http.Client{}
	req, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			fmt.Println("First response byte!")
		},
		GotConn: func(info httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", info)
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", info)
		},
		ConnectStart: func(network, addr string) {
			fmt.Println("Dial start")
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Println("Dial Done!")
		},
		WroteHeaders: func() {
			fmt.Println("Wrote headers")
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	fmt.Println("Requesting data from server !")
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, response.Body)
}
