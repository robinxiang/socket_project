package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("HTTP server is start……")

	//create url route
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Host, "get login")
		writer.Write([]byte("do you want login"))
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Host, "get /")
		writer.Write([]byte("hello there"))
	})

	//Start a http server
	if err := http.ListenAndServe("0.0.0.0:8888", nil); err != nil {
		fmt.Println("server start error:", err)
		return
	}

}
