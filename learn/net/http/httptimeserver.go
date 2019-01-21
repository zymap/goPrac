package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	httpserver()
}

func httpserver() {
	http.HandleFunc("/time", timeHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, t)
}
